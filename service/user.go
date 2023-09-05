package service

import (
	"context"
	"mail/dao"
	"mail/model"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
	"mime/multipart"
)

// 使用标签来映射表单数据和 JSON 数据到结构体的字段中，以及相应的数据转换和处理
// 可以将JSON数据存入结构体中
// 也可以将结构体序列化成JSON
// 实现前后端之前的数据交互 因为前后端要求的数据类型不同
type UserService struct {
	Nickname string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` //暂时前端验证 一般前后端都需要验证
}

func (service *UserService) Registe(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}

	// 10000 -> 密文存储 对称加密操作 不能直接存储金额
	util.Encrypt.Setkey(service.Key)

	// 把上下文信息传递进去
	userDao := dao.NewUserDao(ctx)

	_, exist, err := userDao.ExistByUserName(service.UserName)
	// 出错了
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//不存在 已经屏蔽不存在的错误
	if exist {
		code = e.RegisteError_ExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,
		NickName: service.Nickname,
		Avatar:   "default.jpg",
		Status:   model.Active,
		Monery:   "04 28.01 这里进行初始金额的加密", //
	}
	// 密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	err = userDao.CreatUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	user, exist, err := userDao.ExistByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if !exist {
		code = e.LoginError_No_User
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在 请先注册",
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.LoginError_Wrong_Secret
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误 请重新输入",
		}
	}
	// http是无状态的 不知道你是谁 需要返回一个 （认证：token）
	// token 签发 你看这里那拿了用户ID 放放到Token里了 后面用户做什么事情
	// 通过token里的ID就行
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
			Data:   "token 认证失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}
}

// / 用户修改信息
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 修改昵称

	if service.Nickname != "" {
		user.NickName = service.Nickname
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 头像更新 上传到本地
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, filesize int64) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	userdao := dao.NewUserDao(ctx)
	user, err = userdao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 保存图片到本地 修改数据库的文件路径
	path, err := UploadAvatarToLocal(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUpLoadFailed
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userdao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
