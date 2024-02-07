package routes

import (
	"puzzle/app/controllers"
	"puzzle/app/middlewares/cors"
	"puzzle/app/middlewares/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件服务
	r.Static("/static", "./static")

	//启用跨域中间件
	r.Use(cors.Cors())

	root := r.Group("/api")
	{
		// 用户
		user := root.Group("/user")
		{
			user.POST("/register", controllers.Register) // 用户注册
			user.POST("/login", controllers.Login)       // 用户登录
			user.POST("/list", controllers.ListUser)     // 列表查询
		}

		// 记录
		record := root.Group("/record").Use(jwt.JWT())
		{
			record.POST("/insert", controllers.InsertRecord) // 新增记录
			record.POST("/list", controllers.ListRecord)     // 记录列表
			record.POST("/update", controllers.UpdateRecord) // 更新记录
		}
	}

	return r
}
