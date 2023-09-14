package service

import (
	"context"
	"errors"
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
	Key       string  `form:"key" json:"key"`
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
	condition := make(map[string]interface{})
	condition["user_id"] = uid
	condition["id"] = oid
	order, err := orderDao.GetOrderByCondition(condition)
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

func (service *OrderService) PayDown(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	condition["user_id"] = uid
	condition["order_num"] = service.OrderNum
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderByCondition(condition)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 这里的money我觉得就是总额 怎么可能再乘num
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money, err := util.Encrypt.GetOriginMoney(service.Key, user.Money)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	moneyFloat, err := strconv.ParseFloat(money, 64)
	util.LogrusObj.Infoln(err)
	remainMoney := moneyFloat - order.Money
	if remainMoney < 0 {
		tx.Rollback()
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("余额不足").Error(),
		}

	}
	mm := fmt.Sprintf("%f", remainMoney)
	util.Encrypt.Setkey(service.Key, mm)
	user.Money = util.Encrypt.Getkey()
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		tx.Rollback()
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	_, err = userDao.GetUserById(order.BossID)
	if err != nil {
		tx.Rollback()
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 完了 这里拿不到商家的支付密码 视频里是怎么解密的
	// 视频里也是拿用户的支付密码解密的 666
	// 拿不到就无法对商家的金额进行操作
	//moneyBoss ,err:= util.Encrypt.GetOriginMoney(service.Key, boss.Money)
	// 2. 商家加钱
	// 3. 商品数量减一 欸之前下单的时候有没有确定商品数量够不够
	// 视频弹幕说 事务写法有问题  上面都TODO
	// 4. 删除已支付订单 软删 改成已支付就行

	// 这里逻辑太拉了 商城应该是后续发货 而不是5. 自己的商品加一？？
	// 还是后续看看抖音的项目吧
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
