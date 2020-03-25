package account

import (
	"encoding/json"
	"fmt"
	"io"
)

//keystore 结构，所有NULS体系的keystore都遵循该基本协议
type KeyStore struct {
	//地址
	Address string `json:"address"`
	//使用密码对私钥进行加密后，得到的密文
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
	//公钥Hex
	Pubkey string `json:"pubkey"`
	//加密算法
	cipher string
	//版本号
	version int
	//加密算法相关参数
	cryptoParams CryptoParams
}

//todo 更多加密参数
type CryptoParams struct {
	iv      string
	kdf     string
	extends interface{}
	mac     string
}

//从文件中读取keystore内容，根据内容还原账户数据
func KeystoreFromFile(reader io.Reader) (KeyStore, error) {
	var store KeyStore
	err := json.NewDecoder(reader).Decode(&store)
	if err != nil {
		err = fmt.Errorf("problem parsing keystore, %v", err)
	}
	return store, err
}
