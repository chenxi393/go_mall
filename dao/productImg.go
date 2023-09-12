package dao

import (
	"context"
	"mail/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{newDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) Create(productImg *model.ProductImg) error {
	return dao.Model(&model.ProductImg{}).Create(productImg).Error
}

func (dao *ProductImgDao) FindById(id uint) ([]*model.ProductImg, error) {
	var productImags []*model.ProductImg
	err := dao.Model(&model.ProductImg{}).Where("product_id = ? ", id).Find(&productImags).Error
	return productImags, err
}
