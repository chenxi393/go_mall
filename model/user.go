package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string //`gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"unique"`
	Status         string //有没有被封禁什么的
	Avatar         string `gorm:"size:1000"`
	Monery         string
}
