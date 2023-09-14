package service

import (
	"context"
	"fmt"
	"mail/dao"
	"mail/model"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
	"math/rand"
	"strconv"
	"time"
)

// 这一块业务逻辑写的很烂 一个订单应该是包括多件商品
type OrderService struct {
	ProductId uint    `form:"product_id" json:"product_id"`
	Num       uint    `form:"num" json:"num"`
	AddressId uint    `form:"address_id" json:"address_id"`
	Money     float64 `form:"money" json:"money"`
	BossId    uint    `form:"boss_id" json:"boss_id"`
	Type      uint    `form:"type" json:"type"`
	OrderNum  uint    `form:"order_num" json:"order_num"`
	model.BasePage
}

func (service *OrderService) CreateOrders(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order := &model.Order{
		UserID:    uid,
		ProductID: service.ProductId,
		BossID:    service.BossId,
		AddressID: service.AddressId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1,
	}
	// 校验地址是否存在   这个接口也是有点 哪有让用户自己输boss_ID的 后端还得验证 不然也没啥用
	aid := strconv.Itoa(int(service.AddressId))
	addressDao := dao.NewAddressDao(ctx)
	_, err := addressDao.GetAddressById(uid, aid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    "地址不存在",
			Error:  err.Error(),
		}
	}
	// 是不是还应该去校验商品存不存在
	// 生成订单号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(uid))
	number = userNum + productNum + number
	temp, err := strconv.Atoi(number)
	util.LogrusObj.Infoln(err)
	order.OrderNum = uint64(temp)

	err = orderDao.CreateOrders(order)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) GetOrders(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	// 分页操作 一般大批量获取记录都需要进行分页操作 前面好像有些没有分页 TODO
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uid
	orderDao := dao.NewOrderDao(ctx)
	orders, cnt, err := orderDao.GetOrders(condition, service.BasePage)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 这里分页
	return serializer.BuildListResponse(serializer.BuildOrders(orders), uint(cnt))
}

func (service *OrderService) GetOrderById(ctx context.Context, uid uint, oid string) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderById(uid, oid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 拿到订单号 之后还得拿到地址 拿到商品的详细信息
	addressDao := dao.NewAddressDao(ctx)
	aid := strconv.Itoa(int(order.AddressID))
	address, err := addressDao.GetAddressById(uid, aid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    "不存在该地址",
			Error:  err.Error(),
		}
	}
	// 取商品信息
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductID)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    "不存在该商品",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, address, product),
	}
}

func (service *OrderService) DeleteOrderById(ctx context.Context, uid uint, oid string) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	err := orderDao.DeleteOrderById(uid, oid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
