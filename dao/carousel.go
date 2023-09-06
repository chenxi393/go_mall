package dao

import (
	"context"
	"mail/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

func (dao *CarouselDao) GetCarousels() ([]model.Carousel, error) {
	var carousels []model.Carousel
	err := dao.Model(&model.Carousel{}).Find(&carousels).Error
	return carousels, err
}
