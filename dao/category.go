package dao

import (
	"context"
	"mail/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{newDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

func (dao *CategoryDao) GetCategorys() ([]*model.Category, error) {
	var Categorys []*model.Category
	err := dao.Model(&model.Category{}).Find(&Categorys).Error
	return Categorys, err
}
