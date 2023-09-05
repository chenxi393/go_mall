package e

const (
	Success        = 200
	Error          = 500
	InvaliedParams = 400

	// user模块的错误用3开头区分
	RegisteError_ExistUser  = 30001
	ErrorFailEncryption     = 30002
	LoginError_No_User      = 30003
	LoginError_Wrong_Secret = 30004
	ErrorAuthToken          = 30005
	ErrorAuthToken_TimeOut  = 30006
	ErrorUpLoadFailed       = 30007
	ErrorSendEmail          = 30008
	// product模块的错误用  4XXXXX
)
