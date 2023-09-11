package model
 

// 这里还不懂是干嘛的
type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}
