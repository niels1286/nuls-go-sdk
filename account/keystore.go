/*
 * MIT License
 * Copyright (c) 2019-2020 niels.wang
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
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
