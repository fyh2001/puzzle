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

// WsServer WebSocket 服务端
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

	// 获取用户ID
	// userId, _ := c.Get("userId")
	// userIdStr := strconv.FormatInt(userId.(int64), 10)

	// 初始化客户端
	client := ws.InitClient(conn)

	// 发送客户端唯一标识 ID
	if client.ID == "" {
		return
	}

	go ws.ClientHeartbeatCheck(client.ID)       // 心跳检测
	go client.Reader(&ws.ClientManagerInstance) // 读取消息
	go client.Writer()                          // 发送消息
}
