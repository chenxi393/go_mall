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
		// 轮播图
		v1.GET("carousels", api_v1.ListCarousel)

		// 商品操作
		v1.GET("products",api_v1.ListProduct) //可以分页查看 
		v1.GET("product/:id", api_v1.GetProduct)
		v1.GET("imgs/:id", api_v1.GetProductImgs)
		v1.GET("categories", api_v1.GetCategories)
		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			// 有点疑惑的地方是 这里路由组会中间件鉴权 下面的函数又会鉴权
			authed.PUT("user", api_v1.UserUpdate)      //更新用户信息 这里只更新了昵称
			authed.POST("avatar", api_v1.UploadAvatar) //上传头像
			authed.POST("user/sending-email", api_v1.SendEmail)
			authed.POST("user/valid-email", api_v1.ValidEmail)
			//显示金额
			authed.POST("money", api_v1.ShowMoney)

			// 商品操作 这里用户应该既是买家也是卖家
			authed.POST("product", api_v1.CreatProduct)
			authed.POST("products",api_v1.SearchProduct)

			// 地址操作
			authed.POST("address",api_v1.CreateAddress)
			authed.GET("addresses",api_v1.GetAddresses)
			authed.GET("addresses/:id",api_v1.GetAddress)
			authed.DELETE("addresses/:id",api_v1.DeleteAddress)
			authed.PUT("addresses/:id",api_v1.ModifyAddress)

			// 收藏夹 这里有外键要注意 插入时 外键必须在别的表已经存在
			authed.POST("favorites",api_v1.CreateFavorites)
			authed.DELETE("favorites/:id",api_v1.DeleteFavorites)
			authed.GET("favorites",api_v1.GetFavorites)
		}
	}
	return r
}
