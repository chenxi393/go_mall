package api_v1

import (
	"mail/pkg/util"
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFavorites(ctx *gin.Context) {
	favoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&favoriteService); err == nil {
		res := favoriteService.Create(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorites(ctx *gin.Context) {
	favoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&favoriteService); err == nil {
		res := favoriteService.Delete(ctx.Request.Context(), claims.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func GetFavorites(ctx *gin.Context) {
	favoriteService := service.FavoriteService{}
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&favoriteService); err == nil {
		res := favoriteService.GetAll(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
