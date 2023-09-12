package serializer

import (
	"mail/model"
)

type Address struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(Addresss *model.Address) *Address {
	return &Address{
		ID:       Addresss.ID,
		UserID:   Addresss.UserID,
		Name:     Addresss.Name,
		Phone:    Addresss.Phone,
		Seen:     false,
		CreateAt: Addresss.CreatedAt.Unix(),
		Address:  Addresss.Address,
	}
}

func BuildAddresses(Addresss []*model.Address) []*Address {
	var items []*Address
	for _, item := range Addresss {
		items = append(items, BuildAddress(item))
	}
	return items
}
