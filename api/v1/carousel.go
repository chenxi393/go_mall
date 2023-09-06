package api_v1

import (
	"mail/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCarousel(ctx *gin.Context) {
	var carousel service.CarouselService
	if err := ctx.ShouldBind(&carousel); err == nil {
		res := carousel.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
