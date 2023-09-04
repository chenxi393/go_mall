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
	r.StaticFS("/static", http.Dir("../static"))

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})
		v1.POST("user/register", api_v1.UserRegister)
		v1.POST("user/login",api_v1.UserLogin)
	}
	return r
}
