package dao

import (
	"context"
	"mail/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{newDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) Create(address *model.Address) error {
	return dao.Model(&model.Address{}).Create(address).Error
}

func (dao *AddressDao) GetAddresss(uid uint) ([]*model.Address, error) {
	var Addresses []*model.Address
	err := dao.Model(&model.Address{}).Where("user_id = ?", uid).Find(&Addresses).Error
	return Addresses, err
}

func (dao *AddressDao) GetAddressById(uid uint, id string) (*model.Address, error) {
	var Address *model.Address
	err := dao.Model(&model.Address{}).Where("id = ? AND user_id = ?", id, uid).First(&Address).Error
	return Address, err
}

func (dao *AddressDao) DeleteAddressById(uid uint, id string) error {
	var Address *model.Address
	return dao.Model(&model.Address{}).Where("id = ? AND user_id = ?", id, uid).Delete(&Address).Error
}

func (dao *AddressDao) ModifyAddressById(addr map[string]interface{}, uid uint, id string) error {
	return dao.Model(&model.Address{}).Where("id = ? AND user_id = ?", id, uid).Updates(&addr).Error
}
