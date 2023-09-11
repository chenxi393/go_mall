package api_v1

import (
	"mail/pkg/util"
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductImgs(ctx *gin.Context) {
	var getProductImagsService service.ProductImagService
	if err := ctx.ShouldBind(&getProductImagsService); err == nil {
		res := getProductImagsService.GetImagsById(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
