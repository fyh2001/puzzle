package websocket

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Client 客户端
type Client struct {
	ID            string
	Conn          *websocket.Conn // 连接
	Send          chan []byte     // 发送消息
	Message       *Message        // 消息
	LastHeartbeat time.Time       // 最后一次心跳时间
	BindUserId    string          // 绑定用户ID
	JoinGroup     []string        // 加入的群组
}

// ClientManager 客户端管理
type ClientManager struct {
	ID         string       // 客户端ID
	Clients    sync.Map     // 客户端列表
	Broadcast  chan []byte  // 广播消息
	Register   chan *Client // 新增客户端
	Unregister chan *Client // 删除客户端
	Timer      *time.Timer  // 定时器
	sync.Mutex
}

// WebSocketUser 用户基础信息
type WebSocketUser struct {
	Uid      string
	ClientID []string
}

// WebSocketGroup 分组基础信息
type WebSocketGroup struct {
	ClientID []string
}

type Message struct {
	Type    string `json:"type"`    // 消息类型
	Content string `json:"content"` // 消息
}

var GatewayUser, GatewayGroup, Managers sync.Map

var ClientManagerInstance = ClientManager{
	Clients:    sync.Map{},
	Broadcast:  make(chan []byte, 1024),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Timer:      time.NewTimer(5 * time.Second),
}
