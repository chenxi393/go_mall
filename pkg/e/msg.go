package e

var MsgFlag = map[int]string{
	Success:                 "ok",
	Error:                   "fail",
	InvaliedParams:          "参数错误",
	RegisteError_ExistUser:  "注册时用户已经存在",
	ErrorFailEncryption:     "密码加密失败",
	LoginError_No_User:      "用户名不存在",
	LoginError_Wrong_Secret: "密码错误",
	ErrorAuthToken:          "token 认证失败",
	ErrorAuthToken_TimeOut:  "token 已过期",
	ErrorUpLoadFailed:       "图片上传失败",
	ErrorSendEmail:          "邮件发送失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if !ok {
		return MsgFlag[Error]
	}
	return msg
}
