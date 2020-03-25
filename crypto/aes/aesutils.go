package aesutils

import (
	"crypto/aes"
	"crypto/cipher"
)

//使用AES算法进行加密
//data：需要加密的原始数据
//key：密码
//returns : 加密后的密文
func Encrypt(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}
