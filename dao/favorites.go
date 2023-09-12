package dao

import (
	"context"
	"mail/model"

	"gorm.io/gorm"
)

type FavoritesDao struct {
	*gorm.DB
}

func NewFavoritesDao(ctx context.Context) *FavoritesDao {
	return &FavoritesDao{newDBClient(ctx)}
}

func NewfavoritesDaoByDB(db *gorm.DB) *FavoritesDao {
	return &FavoritesDao{db}
}

func (dao *FavoritesDao) IsExist(pid, uid uint) (bool, error) {
	var cnt int64
	err := dao.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", pid, uid).Count(&cnt).Error
	if err != nil {
		return false, err
	}
	return cnt >= 1, err
}

func (dao *FavoritesDao) Create(favorite *model.Favorite) error {
	return dao.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoritesDao) Getfavorites(uid uint) ([]*model.Favorite, error) {
	var favoritess []*model.Favorite
	err := dao.Model(&model.Favorite{}).Where("user_id = ? ", uid).Find(&favoritess).Error
	return favoritess, err
}

func (dao *FavoritesDao) DeleteById(uid uint, fid string) error {
	var favorite model.Favorite
	// 一直有一个很奇怪的事情 上面用了指针 下面不用 就会出错
	// 上面不是指针 下面取地址就不会  原因是！！！ 上面用指针会初始化为NULL 不能传null进函数里
	// 这里就没管重复删除的事情了 反正也删不掉 虽然返回的ok 实际没删 这个以后再说 TODO
	return dao.Model(&model.Favorite{}).Where("user_id = ? AND id = ? ", uid, fid).Delete(&favorite).Error
}
