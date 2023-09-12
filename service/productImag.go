package service

import (
	"context"
	"mail/dao"
	"mail/serializer"
	"strconv"
)

type ProductImagService struct {
}

func (service *ProductImagService) GetImagsById(ctx context.Context, id string) serializer.Response {
	productImagDao := dao.NewProductImgDao(ctx)
	productid, _ := strconv.Atoi(id)
	productImags, _ := productImagDao.FindById(uint(productid))
	return serializer.BuildListResponse(serializer.BuildImags(productImags), uint(len(productImags)))
}
