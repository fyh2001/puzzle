package rabbitmq

import (
	"fmt"
	"log"
	"puzzle/app/services"
	"puzzle/config"

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
	RankUpdate services.RankUpdate
	Message    string
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
		Mqurl: fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.Settings.RabbitMQ.Username,
			config.Settings.RabbitMQ.Password,
			config.Settings.RabbitMQ.Host,
			config.Settings.RabbitMQ.Port),
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
