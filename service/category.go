package service

import (
	"context"
	"mail/dao"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
)

type CategoriesService struct {

}

func (service *CategoriesService) Get(ctx context.Context) serializer.Response {
	code := e.Success
	CategoryDao := dao.NewCategoryDao(ctx)
	categorys, err := CategoryDao.GetCategorys()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(categorys), uint(len(categorys)))
}
