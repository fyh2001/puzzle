package controllers

import (
	"fmt"
	"net/http"
	"puzzle/app/common/result"
	"time"

	ws "puzzle/app/middlewares/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
}

var (
	socketSet = websocket.Upgrader{
		// 设置消息发送缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在发送大型消息时遇到问题
		WriteBufferSize: 1024,
		// 设置消息接收缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在接收大型消息时遇到问题
		ReadBufferSize: 1024,
		// 消息包启用压缩
		EnableCompression: false,
		// ws握手超时时间
		HandshakeTimeout: 5 * time.Second,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Connect 连接服务端
func (WebSocketController) Connect(c *gin.Context) {
	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("WsServer panic: %v\n", err)
		}
	}()

	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := socketSet.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(200, result.Fail("升级协议失败"))
		return
	}

	// 初始化客户端
	client := ws.ClientManagerInstance.InitClient(conn)

	// 发送客户端唯一标识 ID
	if client.ID == "" {
		return
	}

	go ws.ClientManagerInstance.ClientHeartbeatCheck(client.ID) // 心跳检测
	go client.Reader(&ws.ClientManagerInstance)                 // 读取消息
	go client.Writer()                                          // 发送消息
}

// WsServer WebSocket 根据ID连接服务端
func (WebSocketController) ConnectById(c *gin.Context) {
	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("WsServer panic: %v\n", err)
		}
	}()

	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := socketSet.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(200, result.Fail("升级协议失败"))
		return
	}

	id := c.Query("id")

	clientManager, ok := ws.Managers.Load(id)
	if !ok {
		clientManager = ws.InitClientManager(id)
		go clientManager.(*ws.ClientManager).Start() //开启监听
		go ws.ClientManagerHearBeatCheck(id)         // 心跳检测
	}

	client := clientManager.(*ws.ClientManager).InitClient(conn)

	// 发送客户端唯一标识 ID
	if client.ID == "" {
		return
	}

	go clientManager.(*ws.ClientManager).ClientHeartbeatCheck(client.ID) // 心跳检测
	go client.Reader(clientManager.(*ws.ClientManager))                  // 读取消息
	go client.Writer()                                                   // 发送消息
}
