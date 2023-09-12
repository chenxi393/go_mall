package service

import (
	"context"
	"mail/dao"
	"mail/model"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
)

type AddressService struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
	Phone   string `form:"phone" json:"phone"`
}

func (service *AddressService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Address: service.Address,
		Phone:   service.Phone,
	}
	// 这里create是不是会回写creat_time 到address里
	// 测试还真会 具体见 Create inserts value, returning the inserted data's primary key in value's id
	err := addressDao.Create(address)
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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) GetAll(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addresses, err := addressDao.GetAddresss(uid)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildAddresses(addresses), uint(len(addresses)))
}

func (service *AddressService) GetAddressById(ctx context.Context, uid uint, addressId string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(uid, addressId)
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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) DeleteAddressById(ctx context.Context, uid uint, addressId string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	// 软删除了之后再删除 GORM也不会报错
	// 应该在删除之前 先检查有没有
	_, err := addressDao.GetAddressById(uid, addressId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = addressDao.DeleteAddressById(uid, addressId)
	if err != nil {
		code = e.ErrorDeleteAddress
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

func (service *AddressService) ModifyAddressById(ctx context.Context, uid uint, addressId string) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addr_map := make(map[string]interface{})
	addr_map["name"] = service.Name
	addr_map["phone"] = service.Phone
	addr_map["address"] = service.Address
	// 这里其实也应该判断存不存在
	err := addressDao.ModifyAddressById(addr_map, uid, addressId)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address, err := addressDao.GetAddressById(uid, addressId)
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
		Data:   serializer.BuildAddress(address),
	}
}
