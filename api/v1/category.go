package api_v1

import (
	"mail/pkg/util"
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	var getCategories service.CategoriesService
	if err := ctx.ShouldBind(&getCategories); err == nil {
		res := getCategories.Get(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

