package main

import (
	"puzzle/database"
	"puzzle/routes"
)

func main() {
	database.InitMySQL() // 初始化MySQL数据库连接

	router := routes.InitRouter() // 初始化路由

	_ = router.Run(":8081")
}
