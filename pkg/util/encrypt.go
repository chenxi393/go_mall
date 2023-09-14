package util

import (
	"bytes"
	"crypto/aes"
)

// AES 对称加密
// 这里省略了 有兴趣自己去了解一下
// 视频里复制的代码 我还没找到
var Encrypt *Encryption

type Encryption struct {
	secret []byte
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{
		secret: make([]byte, 16),
	}
}

// 这里有似乎要对key 的密钥进行一个加解密的操作
// 这里省略了
func (e *Encryption) Setkey(key string, originMoney string) error {
	originData := []byte(originMoney)
	k := []byte(key)
	c, err := aes.NewCipher(k)
	if err != nil {
		return err
	}
	blockSize := aes.BlockSize
	originData = ZeroPadding(originData, blockSize)
	out := make([]byte, len(originData))
	c.Encrypt(out, originData)
	copy(e.secret, out)
	return nil
}

func (e *Encryption) Getkey() []byte {
	return e.secret
}

func (e *Encryption) GetOriginMoney(key string, decrypt []byte) (string, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	out := make([]byte, len(decrypt))
	cipher.Decrypt(out, decrypt)
	var i int
	for i = range out {
		if out[i]!=byte(0){
			break
		}
	}
	out=out[i:]
	return string(out), nil
}

func ZeroPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(padText, data...)
}
