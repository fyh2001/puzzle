package database

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func InitRabbitMQ() {
	var err error
	RabbitConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "连接RabbitMQ失败")
	defer RabbitConn.Close()

	RabbitChannel, err = RabbitConn.Channel()
	failOnError(err, "获取RabbitMQ channel失败")
	defer RabbitChannel.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
