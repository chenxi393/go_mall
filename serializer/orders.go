package serializer

import (
	"context"
	"mail/config"
	"mail/dao"
	"mail/model"
	"mail/pkg/util"
	"strconv"
)

type Order struct {
	ID           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	UserID       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossID       uint    `json:"boss_id"`
	Num          uint    `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImagPath     string  `json:"img_path"`
	Money        float64 `json:"money"`
}

func BuildOrder(Orders *model.Order, addr *model.Address, product *model.Product) *Order {
	return &Order{
		ID:           Orders.ID,
		OrderNum:     Orders.OrderNum,
		CreatedAt:    Orders.CreatedAt.Unix(),
		UpdatedAt:    Orders.UpdatedAt.Unix(),
		UserID:       Orders.UserID,
		ProductId:    Orders.ProductID,
		Num:          Orders.Num,
		AddressName:  addr.Address,
		AddressPhone: addr.Phone,
		Address:      addr.Address,
		Type:         Orders.Type,
		ProductName:  product.Name,
		ImagPath:     config.My_path.Host + config.HttpPort + config.My_path.Product + product.ImgPath,
		Money:        Orders.Money,
		BossID:       product.BossID,
	}
}

func BuildOrders(orders []*model.Order) []*Order {
	// 不要在循环里操作数据库 想一想有没有更好的方式
	var items []*Order
	productDao := dao.NewProductDao(context.Background())
	addressDao := dao.NewAddressDao(context.Background())
	for _, item := range orders {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			util.LogrusObj.Infoln(err)
			continue
		}
		address, err := addressDao.GetAddressById(item.UserID, strconv.Itoa(int(item.AddressID)))
		if err != nil {
			util.LogrusObj.Infoln(err)
			continue
		}
		items = append(items, BuildOrder(item, address, product))
	}
	return items
}
