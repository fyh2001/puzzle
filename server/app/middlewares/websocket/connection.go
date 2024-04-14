package websocket

import (
	"fmt"
	"time"
)

// 设置心跳检测间隔时长（秒）
var HeartbeatTime = 180 * time.Second

// ClientManagerHearBeatCheck 服务端心跳检测
func ClientManagerHearBeatCheck(managerId string) {
	for {
		time.Sleep(5 * time.Minute)

		managerInterface, ok := Managers.Load(managerId)
		if !ok {
			break
		}

		manager, _ := managerInterface.(*ClientManager)

		len := 0
		if manager.Clients.Range(func(key, value interface{}) bool {
			len++
			return true
		}); len == 0 {
			fmt.Println("房间无人,自动关闭")
			Managers.Delete(managerId)
			break
		}
	}
}

// ClientHeartbeatCheck 客户端心跳检测
func (manager *ClientManager) ClientHeartbeatCheck(clientID string) {
	for {
		time.Sleep(5 * time.Second)

		clientInterface, ok := manager.Clients.Load(clientID)
		if !ok {
			break
		}

		client, _ := clientInterface.(*Client)
		if time.Since(client.LastHeartbeat) > HeartbeatTime {

			fmt.Printf("Client %s heartbeat timeout", clientID)

			// 注销客户端
			manager.Unregister <- client
			break
		}
	}
}

// clientBindUid 客户端断线时自动踢出Uid绑定列表
func clientUnBindUid(clientID string, uid string) {

	value, ok := GatewayUser.Load(uid)
	if ok {
		user := value.(*WebSocketUser)
		for k, v := range user.ClientID {
			if v == clientID {
				user.ClientID = append(user.ClientID[:k], user.ClientID[k+1:]...)
			}
		}

		if len(user.ClientID) == 0 {
			GatewayUser.Delete(uid)
		}
	}
}

// 客户端断线时自动踢出已加入的群组
func clientLeaveGroup(clientID string) {
	// 使用 Load 方法获取值
	value, ok := ClientManagerInstance.Clients.Load(clientID)
	if !ok {
		// 如果没有找到对应的值，处理相应的逻辑
		return
	}

	client := value.(*Client)

	// 遍历 JoinGroup
	for _, v := range client.JoinGroup {
		// 使用 Load 方法获取值
		groupValue, groupOK := GatewayGroup.Load(v)
		if !groupOK {
			// 如果没有找到对应的值，处理相应的逻辑
			continue
		}

		group := groupValue.(*WebSocketGroup)

		// 在群组中找到对应的 clientID，并删除
		for j, id := range group.ClientID {
			if id == clientID {
				copy(group.ClientID[j:], group.ClientID[j+1:])
				group.ClientID = group.ClientID[:len(group.ClientID)-1]

				// 如果群组中没有成员了，删除群组
				if len(group.ClientID) == 0 {
					GatewayGroup.Delete(v)
				}

				break
			}
		}
	}
}
