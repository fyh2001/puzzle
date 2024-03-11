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
			user.POST("/register", controllers.Register)                     // 用户注册
			user.POST("/login", controllers.Login)                           // 用户登录
			user.POST("/list", controllers.ListUser)                         // 列表查询
			user.POST("/get-user-info", jwt.JWT(), controllers.GetUserInfo)  // 获取用户信息
			user.POST("/update-avatar", jwt.JWT(), controllers.UpdateAvatar) // 更新用户头像
			user.POST("/update", jwt.JWT(), controllers.UpdateUser)          // 更新用户
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

		notification := root.Group("/notification")
		{
			notification.POST("/insert", controllers.InsertNotification) // 发布通知
			notification.POST("/list", controllers.ListNotification)     // 通知列表
		}

		// 管理员
		admin := root.Group("/admin")
		{
			// 鉴权
			authorization := admin.Group("/authorization")
			{
				authorization.POST("/reset-otp", controllers.ResetOtp)          // 重置otp
				authorization.POST("/authorization", controllers.Authorization) // 验证otp
				authorization.POST("/get-url", controllers.GetUrl)              // 获取url
			}

			// 用户
			userManage := admin.Group("/user-manage").Use(jwt.AdminJWT())
			{
				userManage.POST("/list", controllers.ListUserData)     // 用户列表
				userManage.POST("/update", controllers.UpdateUserData) // 更新用户
			}

			// 记录
			recordManage := admin.Group("/record-manage").Use(jwt.AdminJWT())
			{
				recordManage.POST("/list", controllers.ListRecordData)     // 记录列表
				recordManage.POST("/update", controllers.UpdateRecordData) // 更新记录
			}

			// 最佳单次记录
			recordBestSingleManage := admin.Group("/record-best-single-manage").Use(jwt.AdminJWT())
			{
				recordBestSingleManage.POST("/list", controllers.ListRecordBestSingleData) // 最佳单次记录列表
			}

			// 最佳平均记录
			recordBestAverageManage := admin.Group("/record-best-average-manage").Use(jwt.AdminJWT())
			{
				recordBestAverageManage.POST("/list", controllers.ListRecordBestAverageData) // 最佳平均记录列表
			}

			// 最佳步数记录
			recordBestStepManage := admin.Group("/record-best-step-manage").Use(jwt.AdminJWT())
			{
				recordBestStepManage.POST("/list", controllers.ListRecordBestStepData) // 最佳步数记录列表
			}
		}
	}

	return r
}
