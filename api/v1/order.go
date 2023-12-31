package api_v1

import (
	"mail/pkg/util"
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrders(ctx *gin.Context) {
	var orderssService service.OrderService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&orderssService); err == nil {
		res := orderssService.CreateOrders(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func GetOrders(ctx *gin.Context) {
	var orderssService service.OrderService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&orderssService); err == nil {
		res := orderssService.GetOrders(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func GetOrderById(ctx *gin.Context) {
	var orderssService service.OrderService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&orderssService); err == nil {
		res := orderssService.GetOrderById(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteOrderById(ctx *gin.Context) {
	var orderssService service.OrderService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&orderssService); err == nil {
		res := orderssService.DeleteOrderById(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 其实我觉得这个接口应该是输入订单号 order_num和支付密码 key就可以了
// 再一个token里的user_id 三者就可以完成支付 不知道接口里那么多参数干嘛
func OrderPay(ctx *gin.Context) {
	var orderssService service.OrderService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&orderssService); err == nil {
		res := orderssService.PayDown(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
