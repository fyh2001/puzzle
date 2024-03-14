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
			user.POST("/register", controllers.User.Register)                     // 用户注册
			user.POST("/login", controllers.User.Login)                           // 用户登录
			user.POST("/list", controllers.User.List)                             // 列表查询
			user.POST("/get-user-info", jwt.JWT(), controllers.User.GetUserInfo)  // 获取用户信息
			user.POST("/update-avatar", jwt.JWT(), controllers.User.UpdateAvatar) // 更新用户头像
			user.POST("/update", jwt.JWT(), controllers.User.Update)              // 更新用户
		}

		// 记录
		record := root.Group("/record").Use(jwt.JWT())
		{
			record.POST("/insert", controllers.Record.Insert)                     // 新增记录
			record.POST("/list-record", controllers.Record.List)                  // 记录列表
			record.POST("/list-best-single", controllers.RecordBestSingle.List)   // 最佳单次记录列表
			record.POST("/list-best-average", controllers.RecordBestAverage.List) // 最佳平均记录列表
			record.POST("/list-best-step", controllers.RecordBestStep.List)       // 最佳步数记录列表
			record.POST("/update", controllers.Record.Update)                     // 更新记录
		}

		// 打乱
		scramble := root.Group("/scramble").Use(jwt.JWT())
		{
			scramble.POST("/get-new-scramble", controllers.Scramble.GetNewScamble)    // 获取新打乱
			scramble.POST("/get-user-scramble", controllers.Scramble.GetUserScramble) // 获取用户打乱
		}

		notification := root.Group("/notification")
		{
			notification.POST("/insert", controllers.Notification.Insert)        // 发布通知
			notification.POST("/list", jwt.JWT(), controllers.Notification.List) // 通知列表
		}

		// 管理员
		admin := root.Group("/admin")
		{
			// 鉴权
			authorization := admin.Group("/authorization")
			{
				authorization.POST("/reset-otp", controllers.AdminAuthorization.ResetOtp)          // 重置otp
				authorization.POST("/authorization", controllers.AdminAuthorization.Authorization) // 验证otp
				authorization.POST("/get-url", controllers.AdminAuthorization.GetUrl)              // 获取url
			}

			// 用户
			userManage := admin.Group("/user-manage").Use(jwt.AdminJWT())
			{
				userManage.POST("/list", controllers.Admin.ListUserData)     // 用户列表
				userManage.POST("/update", controllers.Admin.UpdateUserData) // 更新用户
			}

			// 记录
			recordManage := admin.Group("/record-manage").Use(jwt.AdminJWT())
			{
				recordManage.POST("/list", controllers.Admin.ListRecordData)     // 记录列表
				recordManage.POST("/update", controllers.Admin.UpdateRecordData) // 更新记录
			}

			// 最佳单次记录
			recordBestSingleManage := admin.Group("/record-best-single-manage").Use(jwt.AdminJWT())
			{
				recordBestSingleManage.POST("/list", controllers.Admin.ListRecordBestSingleData) // 最佳单次记录列表
			}

			// 最佳平均记录
			recordBestAverageManage := admin.Group("/record-best-average-manage").Use(jwt.AdminJWT())
			{
				recordBestAverageManage.POST("/list", controllers.Admin.ListRecordBestAverageData) // 最佳平均记录列表
			}

			// 最佳步数记录
			recordBestStepManage := admin.Group("/record-best-step-manage").Use(jwt.AdminJWT())
			{
				recordBestStepManage.POST("/list", controllers.Admin.ListRecordBestStepData) // 最佳步数记录列表
			}
		}
	}

	return r
}
