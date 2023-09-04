package util

// AES 对称加密
// 这里省略了 有兴趣自己去了解一下
// 视频里复制的代码 我还没找到
var Encrypt *Encryption

type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

func (e *Encryption) Setkey(key string) {
	e.key = key
}
