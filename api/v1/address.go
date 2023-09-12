package api_v1

import (
	"mail/pkg/util"
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(ctx *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&address); err == nil {
		res := address.Create(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func GetAddresses(ctx *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&address); err == nil {
		res := address.GetAll(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func GetAddress(ctx *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&address); err == nil {
		res := address.GetAddressById(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func DeleteAddress(ctx *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&address); err == nil {
		res := address.DeleteAddressById(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func ModifyAddress(ctx *gin.Context) {
	var address service.AddressService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&address); err == nil {
		res := address.ModifyAddressById(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
