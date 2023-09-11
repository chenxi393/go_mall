package api_v1

import (
	"github.com/gin-gonic/gin"
	"mail/pkg/util"
	"mail/service"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Registe(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	if err := ctx.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UploadAvatar(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	filesize := fileHeader.Size
	var userLoadAvatar service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&userLoadAvatar); err == nil {
		res := userLoadAvatar.Post(ctx.Request.Context(), claims.ID, file, filesize)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func SendEmail(ctx *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ValidEmail(ctx *gin.Context) {
	var validEmail service.ValidEmailService
	if err := ctx.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ShowMoney(ctx *gin.Context) {
	var Show_Money service.ShowMoneyService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&Show_Money); err == nil {
		res := Show_Money.Show(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
