package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string //`gorm:"unique"`
	PasswordDigest string
	NickName       string `gorm:"unique"`
	Status         string //有没有被封禁什么的
	Avatar         string `gorm:"size:1000"`
	Money          string
	// 字段是又规律的  UserName user_name Username username
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" // 激活用户
)

// 同一密码加密多次 不是一样的密文
func (user *User) SetPassword(password string) error {
	// 这里掉包进行加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
