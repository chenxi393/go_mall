package service

import (
	"context"
	"mail/config"
	"mail/dao"
	"mail/model"
	"mail/pkg/e"
	"mail/pkg/util"
	"mail/serializer"
	"mime/multipart"
	"strings"
	"time"

	"gopkg.in/mail.v2"
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

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	//1 绑定邮箱 2 解绑邮箱 3 更改密码
}

type ValidEmailService struct {
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
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
	// TODO 这里是那支付密码来加密10000？？余额
	err:=util.Encrypt.Setkey(service.Key, "10000")
	if err!=nil{
		util.LogrusObj.Infoln(err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "加密失败",
		}
	}
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
		Money:    util.Encrypt.Getkey(),
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

func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice
	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	// 这里时根据类型 去数据库拿通知文本 总感觉写法很奇怪
	//还需要手动在数据库存入文本
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = config.Email.ValidEmail + token //发送方
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	// -1表示替换所有 可以看If n < 0, there is no limit on the number of replacements.
	m := mail.NewMessage()
	m.SetHeader("From", config.Email.SMTPEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "go-mail 你正在进行邮箱操作")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(config.Email.SMTPHost, 465, config.Email.SMTPEmail, config.Email.SMTPPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
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

func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	// 验证token
	if token == "" {
		code = e.InvaliedParams
	} else {
		claims, err := util.ParseEmailToken(token)
		//ParseEmailToken 写错了 导致err确实等于nil 但是claim 也等于nil
		//下面时间比较会导致空指针错误
		if err != nil || claims == nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthToken_TimeOut
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}

	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//获取用户的信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if operationType == 1 {
		user.Email = email
	} else if operationType == 2 {
		user.Email = ""
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
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

func (service *ShowMoneyService) Show(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
