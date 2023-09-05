package serializer

import (
	"mail/model"
	"mail/pkg/util"
)

type Monery struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) Monery {
	util.Encrypt.Setkey(key) //把当前的key放进去
	return Monery{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: user.Money, //这里省略了上面解密的过程
	}
}
