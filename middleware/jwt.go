package middleware

import (
	"github.com/gin-gonic/gin"
	"mail/pkg/e"
	"mail/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.Success
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthToken_TimeOut
			}
		}

		if code != e.Success {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort() // 鉴权失败 停止后续调用
			return
		}
		ctx.Next()
	}
}
