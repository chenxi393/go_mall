package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	CategoryID    int
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"defalt:false"`
	Num           int
	BossID        int
	BossName      string
	BossAvatar    string
}
