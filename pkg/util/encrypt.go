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

// 这里有似乎要对key 的密钥进行一个加解密的操作
// 这里省略了
func (e *Encryption) Setkey(key string) {
	e.key = key
}

func (e *Encryption) Getkey() string {
	return e.key
}
