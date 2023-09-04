package serializer

import (
	"mail/config"
	"mail/model"
)

type User struct { //vo view objective
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   config.My_path.Host + config.HttpPort + config.My_path.Avatar + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
