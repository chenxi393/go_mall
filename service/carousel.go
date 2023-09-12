package service

import (
	"context"
	"mail/dao"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(ctx context.Context) serializer.Response {
	code := e.Success
	carouselDao := dao.NewCarouselDao(ctx)
	carrousels, err := carouselDao.GetCarousels()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carrousels), uint(len(carrousels)))
}
