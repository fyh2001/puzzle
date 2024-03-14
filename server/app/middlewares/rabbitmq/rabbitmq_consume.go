package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
)

// Consume 使用 goroutine 消费消息
func (r *RabbitMQ) Consume(callback CallbackFunc[any]) {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false, // 是否为自动删除
		false, // 是否具有排他性
		false, // 是否阻塞
		nil,   // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    // 用来区分多个消费者
		false, // 是否自动应答
		false, // 是否具有排他性
		false, // 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false, // 队列消费是否阻塞
		nil,   // 额外属性
	)

	if err != nil {
		fmt.Println(err)
	}

	// 设置Qos
	err = r.channel.Qos(1, 0, false)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协和处理消息
	go func() {
		for d := range msgs {
			// 处理消息
			result := RabbitMQMessage{}
			msg := json.Unmarshal(d.Body, &result)

			if msg != nil {
				log.Printf("消息解析失败: %s", msg)
			}

			if result.Message == "rank update" {
				callback(result.RankUpdate) // 调用回调函数处理消息
				log.Println("记录排名更新成功")
			}

			if result.Message == "notification-all" {
				callback(result.NotificationMsg) // 调用回调函数处理消息
				log.Println("发送全体通知成功")
			}

			d.Ack(false) // 手动应答
		}
	}()
	log.Printf("[%s] 正在等待消息, 按下 CTRL+C 退出", r.QueueName)
	<-forever
}
