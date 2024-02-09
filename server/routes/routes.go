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
			record.POST("/insert", controllers.InsertRecord)                     // 新增记录
			record.POST("/list-record", controllers.ListRecord)                  // 记录列表
			record.POST("/list-best-single", controllers.ListRecordBestSingle)   // 最佳单次记录列表
			record.POST("/list-best-average", controllers.ListRecordBestAverage) // 最佳平均记录列表
			record.POST("/list-best-step", controllers.ListRecordBestStep)       // 最佳步数记录列表
			record.POST("/update", controllers.UpdateRecord)                     // 更新记录
		}

		// 打乱
		scramble := root.Group("/scramble").Use(jwt.JWT())
		{
			scramble.POST("/get-new-scramble", controllers.GetNewScamble)    // 获取新打乱
			scramble.POST("/get-user-scramble", controllers.GetUserScramble) // 获取用户打乱
		}
	}

	return r
}
