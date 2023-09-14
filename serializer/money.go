package serializer

import (
	"mail/model"
	"mail/pkg/util"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) Money {
	m,err:=util.Encrypt.GetOriginMoney(key,user.Money)
	if err!=nil{
		util.LogrusObj.Infoln(err)
		return Money{}
	}
	return Money{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: m,
	}
}
