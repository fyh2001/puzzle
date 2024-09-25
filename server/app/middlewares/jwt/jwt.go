package jwt

import (
	jwtUtil "puzzle/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 40001 // 无token，无权限访问
		} else {
			// 解析token
			claims, err := jwtUtil.ParseToken(token)
			if err != nil {
				code = 40002 // token不合法
			} else {
				c.Set("userID", claims.Id)
				c.Set("username", claims.Username)
			}
			//else if time.Now().Unix() > claims.ExpiresAt {
			//	code = result.TokenExpired.Code // token已过期
			//}

		}
		if code != 200 {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  "无权限访问",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
