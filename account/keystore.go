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
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	cryptoutils "github.com/niels1286/nuls-go-sdk/crypto/utils"
	"io"
	"regexp"
)

//keystore 结构，所有NULS体系的keystore都遵循该基本协议
//todo 下一个版本增加加密算法的说明和算法参数
type KeyStore struct {
	//地址
	Address string `json:"address"`
	//使用密码对私钥进行加密后，得到的密文
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
	//公钥Hex
	Pubkey string `json:"pubkey"`
	//版本号
	version int
}

//从文件中读取keystore内容，根据内容还原账户数据
func KeystoreFromFile(reader io.Reader) (KeyStore, error) {
	store := KeyStore{}
	err := json.NewDecoder(reader).Decode(&store)
	if err != nil {
		err = fmt.Errorf("problem parsing keystore, %v", err)
	}
	return store, err
}

//指定账户生成对应的keystore
func CreateKeyStore(account Account, password string) (KeyStore, error) {
	if !PasswordCheck(password) {
		return KeyStore{}, errors.New("Invalid password format")
	}
	key := cryptoutils.Sha256h([]byte(password))
	epk := cryptoutils.AESEncrypt(account.GetPriKeyBytes(), key)
	epkHex := hex.EncodeToString(epk)
	return KeyStore{Address: account.Address, EncryptedPrivateKey: epkHex, Pubkey: account.GetPubKeyHex(true), version: 1}, nil
}

//校验密码是否满足格式要求，如果不满足则返回false。
//密码至少8位，必须包含字母和数字
func PasswordCheck(password string) bool {
	if password == "" {
		return false
	}
	length := len(password)
	if length < 8 || length > 20 {
		return false
	}
	reg, _ := regexp.Compile("(.*)[a-zA-Z](.*)")
	if !reg.MatchString(password) {
		return false
	}
	reg, _ = regexp.Compile("(.*)\\d+(.*)")
	if !reg.MatchString(password) {
		return false
	}
	reg, _ = regexp.Compile("(.*)\\s+(.*)")
	if reg.MatchString(password) {
		return false
	}
	reg, _ = regexp.Compile("(.*)[\u4e00-\u9fa5\u3000]+(.*)")
	if reg.MatchString(password) {
		return false
	}
	return true
}
