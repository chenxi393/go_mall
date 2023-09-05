package routes

import (
	api_v1 "mail/api/v1"
	"mail/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	// 程序可执行文件的位置不同 会导致这里失效
	// 若想同一个静态文件路由 到多个本地文件地址
	// 可能需要自己写中间件 可以自己再去了解一下
	// 也可以看看有没有别的实现方式

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})
		v1.POST("user/register", api_v1.UserRegister)
		v1.POST("user/login", api_v1.UserLogin)

		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api_v1.UserUpdate)      //更新用户信息 这里只更新了昵称
			authed.POST("avatar", api_v1.UploadAvatar) //上传头像
			authed.POST("user/sending-email", api_v1.SendEmail)
			authed.POST("user/valid-email", api_v1.ValidEmail)

			authed.POST("money", api_v1.ShowMoney)
		}
	}
	return r
}
