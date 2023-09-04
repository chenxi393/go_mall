package serializer

// Response 基础序列化器
// 空接口可以接受任何类型
// 使用类型断言
// 可以获得具体类型的值
// data=[]int{1,2,3}
// str, ok := data.([]int)
// 也可以直接打印空接口

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
