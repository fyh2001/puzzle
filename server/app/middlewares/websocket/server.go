package websocket

import (
	"encoding/json"
	"log"
	jwt "puzzle/utils/jwt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// InitWsServer 初始化WebSocket服务端
func InitWsServer() {
	go ClientManagerInstance.Start()
}

// InitClient 初始化客户端
func InitClient(conn *websocket.Conn) *Client {

	// clientID也可以用其他方式生成，只要能保证在所有服务端中都能保证唯一即可
	clientID := uuid.New().String()

	client := &Client{
		ID:            clientID,
		Conn:          conn,
		Send:          make(chan []byte, 1024),
		Message:       new(Message),
		BindUserId:    "",
		JoinGroup:     make([]string, 0),
		LastHeartbeat: time.Now(),
	}

	// 注册客户端
	ClientManagerInstance.Register <- client

	// if err := conn.WriteMessage(websocket.TextMessage, []byte(clientID)); err != nil {
	// 	ClientManagerInstance.Unregister <- client
	// 	return nil
	// }

	return client
}

// Start 启动客户端管理
func (manager *ClientManager) Start() {
	defer func() {
		close(manager.Register)
		close(manager.Unregister)
		close(manager.Broadcast)
	}()

	//协程 另一条线程
	for {
		//阻塞监听channel消息

		// 通道
		select {
		case conn := <-manager.Register:
			manager.Clients.Store(conn.ID, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients.Load(conn); ok {
				conn.CloseClient()
			}

		case message := <-manager.Broadcast:
			log.Println("广播消息")
			manager.Clients.Range(func(key, value interface{}) bool {
				client := value.(*Client)
				select {
				case client.Send <- message:
				default:
					client.CloseClient()
				}
				return true
			})
		}
	}
}

// closeClient 关闭客户端
func (c *Client) CloseClient() {
	if value, ok := ClientManagerInstance.Clients.Load(c.ID); ok {
		client := value.(*Client)
		if client.BindUserId != "" {
			clientUnBindUid(c.ID, client.BindUserId)
		}

		if len(client.JoinGroup) > 0 {
			clientLeaveGroup(c.ID)
		}

		ClientManagerInstance.Clients.Delete(c.ID)
	}

	close(c.Send)  // 关闭发送消息通道
	c.Conn.Close() // 关闭连接
}

// writer 发送消息
func (c *Client) Writer() {
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

// reader 读取消息
func (c *Client) Reader(manager *ClientManager) {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("读取消息失败")
			return
		}

		json.Unmarshal(message, c.Message)

		switch c.Message.Type {
		case "ping": // 心跳消息
			message, _ := json.Marshal(&Message{
				Type:    "pong",
				Content: "pong",
			})
			c.Send <- message            // 回复心跳消息
			c.LastHeartbeat = time.Now() // 更新心跳时间

		case "auth": // 认证消息
			claims, _ := jwt.ParseToken(c.Message.Content)
			if claims != nil && claims.Id != 0 { // 认证成功
				userId := strconv.FormatInt(claims.Id, 10)
				c.bindUserId(userId) // 绑定用户ID
			} else { // 认证失败
				message, _ := json.Marshal(&Message{
					Type:    "auth",
					Content: "登录信息有误，请重新登录",
				})
				c.Send <- message       // 回复认证消息
				manager.Unregister <- c // 注销客户端
			}
		}
	}
}

// clientUnBindUid 客户端解绑用户ID
func (c *Client) bindUserId(userId string) {
	c.BindUserId = userId

	// 绑定用户ID
	userBase, ok := GatewayUser.Load(userId)
	if !ok {
		userBase = &WebSocketUser{
			Uid:      userId,
			ClientID: make([]string, 0),
		}
	}

	userBase.(*WebSocketUser).ClientID = append(userBase.(*WebSocketUser).ClientID, c.ID)

	GatewayUser.Store(userId, userBase)
}
