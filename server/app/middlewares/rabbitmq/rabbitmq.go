package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn         *amqp.Connection // 连接
	channel      *amqp.Channel    // 通道
	ExchangeName string           // 交换机名称
	QueueName    string           // 队列名称
	RouteKey     string           // 路由键
	Mqurl        string           // 连接信息
}

type RabbitMQMessage struct {
	Dimension int
	Message   string
}

// failOnErr 错误处理函数
func (r *RabbitMQ) failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMQ 创建RabbitMQ实例
func NewRabbitMQ(queueName, exchangeName, routeKey string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName:    queueName,
		ExchangeName: exchangeName,
		RouteKey:     routeKey,
		Mqurl:        "amqp://guest:guest@localhost:5672/",
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err, "连接RabbitMQ失败")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "获取RabbitMQ channel失败")

	return rabbitmq
}

// Destory 断开channel和connection
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}
