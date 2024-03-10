package main

import (
	"puzzle/app/middlewares/rabbitmq"
	"puzzle/config"
	"puzzle/database"
	"puzzle/routes"
)

func main() {
	config.InitConfig() // 初始化配置文件

	database.InitMySQL() // 初始化MySQL数据库连接
	database.InitRedis() // 初始化Redis数据库连接

	// 初始化队列和消费者
	go rabbitmq.InitQueuesAndConsumers()
	defer func() {
		for _, mq := range rabbitmq.Queues {
			mq.Destory()
		}
	}()

	router := routes.InitRouter() // 初始化路由

	_ = router.Run(":8081")
}
