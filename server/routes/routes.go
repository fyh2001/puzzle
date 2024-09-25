package routes

import (
	"puzzle/app/controllers"
	"puzzle/app/middlewares/cos"
	"puzzle/app/middlewares/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	r.Use(cos.Cors())

	root := r.Group("/api")
	{
		user := root.Group("/user")
		{
			user.POST("/register", controllers.User.Register)                      // 用户注册
			user.POST("/login", controllers.User.Login)                            // 用户登录
			user.POST("/list", jwt.JWT(), controllers.User.List)                   // 查询列表
			user.POST("/update", jwt.JWT(), controllers.User.Update)               // 用户更新
			user.POST("/delete", jwt.JWT(), controllers.User.Update)               // 用户注销
			user.POST("/listByAdmin", jwt.JWT(), controllers.User.ListByAdmin)     // 管理员查询列表
			user.POST("/updateByAdmin", jwt.JWT(), controllers.User.UpdateByAdmin) // 管理员更新
			user.POST("/deleteByAdmin", jwt.JWT(), controllers.User.DeleteByAdmin) // 管理员注销
		}
	}

	return r
}
