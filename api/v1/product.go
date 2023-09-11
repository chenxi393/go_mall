package api_v1

import (
	"mail/pkg/util"
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatProduct(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var creatProductService service.ProductService
	if err := ctx.ShouldBind(&creatProductService); err == nil {
		res := creatProductService.Create(ctx.Request.Context(), claims.ID, files)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ListProduct(ctx *gin.Context) {
	var ListProductsService service.ProductService
	if err := ctx.ShouldBind(&ListProductsService); err == nil {
		res := ListProductsService.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func SearchProduct(ctx *gin.Context) {
	var searchProductService service.ProductService
	if err := ctx.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func GetProduct(ctx *gin.Context) {
	var getProductService service.ProductService
	if err := ctx.ShouldBind(&getProductService); err == nil {
		res := getProductService.Get(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
