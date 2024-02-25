package rabbitmq

import (
	commonService "puzzle/app/services"
)

type CallbackFunc[T any] func(T) error

type rabbitMQ[T any] struct {
	ExchangeName string          // 交换机名称
	QueueName    string          // 队列名称
	callback     CallbackFunc[T] // 回调函数
}

var Queues []*RabbitMQ

// 队列和更新操作的映射关系配置
var recordRankUpdateQueue = []rabbitMQ[any]{
	{
		QueueName:    "best_single_rank_update_queue",
		ExchangeName: "",
		callback:     commonService.UpdateRecordBestSingleRank,
	},
	{
		QueueName:    "best_average_rank_update_queue",
		ExchangeName: "",
		callback:     nil,
	},
	{
		QueueName:    "best_step_rank_update_queue",
		ExchangeName: "",
		callback:     nil,
	},
}

// 初始化队列和更新操作
func InitQueuesAndConsumers() {

	for _, q := range recordRankUpdateQueue {
		mq := NewRabbitMQ(q.QueueName, q.ExchangeName, "")
		go mq.Consume(q.callback)

		Queues = append(Queues, mq)
	}

}
