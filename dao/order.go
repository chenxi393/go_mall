package dao

import (
	"context"
	"mail/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{newDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) CreateOrders(Order *model.Order) error {
	return dao.Model(&model.Order{}).Create(Order).Error
}

func (dao *OrderDao) GetOrders(condition map[string]interface{}, page model.BasePage) ([]*model.Order, int64, error) {
	var Orderes []*model.Order
	var cnt int64
	err := dao.Model(&model.Order{}).Where(condition).Count(&cnt).Error
	if err != nil {
		return Orderes, 0, err
	}
	err = dao.Model(&model.Order{}).Offset((page.PageNum-1)*page.PageSize).Limit(page.PageSize).Find(&Orderes, condition).Error
	return Orderes, cnt, err
}

func (dao *OrderDao) GetOrderByCondition(condition map[string]interface{}) (*model.Order, error) {
	var Order *model.Order
	err := dao.Model(&model.Order{}).First(&Order, condition).Error
	return Order, err
}

func (dao *OrderDao) DeleteOrderById(uid uint, id string) error {
	var Order *model.Order
	return dao.Model(&model.Order{}).Where("id = ? AND user_id = ?", id, uid).Delete(&Order).Error
}
