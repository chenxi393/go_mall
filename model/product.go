package model

import (
	"context"
	"mail/cache"
	"mail/pkg/util"
	"strconv"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryID    uint
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"defalt:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, err := cache.RedisClient.Get(context.Background(), cache.ProductViewKey(product.ID)).Result()
	if err!=nil{
		util.LogrusObj.Infoln(err)
	}
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	cache.RedisClient.Incr(context.Background(), cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(context.Background(), cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
