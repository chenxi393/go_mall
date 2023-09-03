package model

import "gorm.io/gorm"

// Order 订单信息
type Order struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not null"`
	AddressID uint `gorm:"not null"`
	Num       uint
	OrderNum  uint64
	Type      uint // 1未支付 2 已支付
	Money     float64
}
