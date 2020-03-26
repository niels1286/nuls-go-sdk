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
	"errors"
	"github.com/niels1286/nuls-go-sdk/crypto/base58"
	"github.com/niels1286/nuls-go-sdk/crypto/eckey"
	cryptoutils "github.com/niels1286/nuls-go-sdk/crypto/utils"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"log"
	"strings"
)

const (
	//NULS主网ID
	NULSChainId = uint16(1)
	//NULS主网的默认前缀
	NULSPrefix = "NULS"
	//NULS测试网id
	TNULSChainId = uint16(2)
	//NULS测试网的地址前缀
	TNULSPrefix = "tNULS"
	//默认的地址字节长度
	AddressBytesLength = 23
	//默认账户类型
	NormalAccountType = uint8(1)
	//合约账户类型
	ContractAccountType = uint8(2)
	//多签账户类型
	P2SHAccountType = uint8(3)
)

//用于前缀和实际地址的分隔符
var PrefixTable = [...]string{"", "a", "b", "c", "d", "e"}

//NULS大生态体系内的基本账户结构
type Account struct {
	//账户对应的公私钥对
	*eckey.EcKey
	//地址，是用户操作的载体，账户余额管理、转账、数据权限等，全有地址来识别
	//地址的格式为：address = prefix + Base58Encode(chainId+addressType+pkh+xor)
	Address string
	//地址的字节数组，在链上存储是，使用该格式
	AddressBytes []byte
	//NULS是一个多链生态，每一条区块链，都有自己的id，用于在多链交互时，识别所属身份
	//NULS主网为1，NULS测试网为2
	ChainId uint16
	//账户类型：1、普通账户，2、合约地址，3、多签地址
	AccType uint8
	//地址前缀
	Prefix string
}

//创建一个新账户
//chainId:账户所在链的ID，用于多链交互时，识别身份
//prefix:地址前缀，每条链都可以定义自己的地址前缀
func NewNormalAccount(chainId uint16, prefix string) (*Account, error) {
	ec, err := eckey.NewEcKey()
	if err != nil {
		return nil, err
	}
	return getAccountByEckey(ec, chainId, prefix)
}

//根据EcKey生成账户
func getAccountByEckey(ec *eckey.EcKey, chainId uint16, prefix string) (*Account, error) {
	pubBytes := ec.GetPubKeyBytes(true)
	addressBytes := GetAddressByPubBytes(pubBytes, chainId, NormalAccountType, prefix)
	address := GetStringAddress(addressBytes, prefix)
	return &Account{
		Address:      address,
		AddressBytes: addressBytes,
		ChainId:      chainId,
		AccType:      NormalAccountType,
		EcKey:        ec,
		Prefix:       prefix,
	}, nil
}

//去除前缀，获得真正的地址字符串
func getRealAddress(address string) (prefix, realAddress string) {
	if strings.HasPrefix(address, NULSPrefix) {
		return NULSPrefix, address[5:]
	}
	if strings.HasPrefix(address, TNULSPrefix) {
		return TNULSPrefix, address[6:]
	}
	for index, c := range address {
		if c >= 97 {
			return address[0 : index-1], address[index+1:]
		}
	}
	return "", ""
}

//根据地址字节数组，生成可以阅读的字符串地址
func GetStringAddress(bytes []byte, prefix string) string {
	//将之前得到的所有字节，进行异或操作，得到结果追加到
	xor := calcXor(bytes)
	bytes = append(bytes, xor)
	return prefix + PrefixTable[len(prefix)] + base58.Encode(bytes)
}

//根据公钥，生成账户地址
func GetAddressByPubBytes(bytes []byte, chainId uint16, accountType uint8, prefix string) []byte {
	hash160 := cryptoutils.Hash160(bytes)
	addressBytes := []byte{}
	addressBytes = append(addressBytes, mathutils.Uint16ToBytes(chainId)...)
	addressBytes = append(addressBytes, accountType)
	addressBytes = append(addressBytes, hash160...)
	return addressBytes
}

//计算异或字节
func calcXor(bytes []byte) byte {
	xor := byte(0)
	for _, one := range bytes {
		xor ^= one
	}
	return xor
}

//创建一个NULS主网账户
func NewNULSAccount(count int) ([]*Account, error) {
	return BatchNewAccount(count, NULSChainId, NULSPrefix)
}

//创建一个NULS测试网账户
func NewTNULSAccount(count int) ([]*Account, error) {
	return BatchNewAccount(count, TNULSChainId, TNULSPrefix)
}

//批量创建账户
func BatchNewAccount(count int, chainId uint16, prefix string) ([]*Account, error) {
	result := []*Account{}
	resultChannel := make(chan *Account)
	for i := 0; i < count; i++ {
		go func() {
			account, err := NewNormalAccount(chainId, prefix)
			if err != nil {
				log.Fatal("Create account failed.")
			}
			resultChannel <- account
		}()
	}
	for i := 0; i < count; i++ {
		account := <-resultChannel
		if account.Address == "" {
			continue
		}
		result = append(result, account)
	}
	return result, nil
}

//根据地址还原账户基本信息
func ParseAccount(address string) (Account, error) {
	if address == "" {
		return Account{}, errors.New("The address is blank.")
	}
	prefix, realAddressStr := getRealAddress(address)
	bytes := base58.Decode(realAddressStr)
	chainId := mathutils.BytesToUint16(bytes[0:2])
	accountType := bytes[2]
	addressBytes := bytes[0 : len(bytes)-1]
	return Account{
		Address:      address,
		AddressBytes: addressBytes,
		ChainId:      chainId,
		AccType:      accountType,
		EcKey:        nil,
		Prefix:       prefix,
	}, nil
}

//验证地址是否正确
func Valid(address string) bool {
	if address == "" {
		return false
	}
	prefix, realAddressStr := getRealAddress(address)
	bytes := base58.Decode(realAddressStr)
	//长度必须正确，默认长度+一个校验位（xor）
	if len(bytes) != AddressBytesLength+1 {
		return false
	}
	chainId := mathutils.BytesToUint16(bytes[0:2])
	//验证已知链的前缀是否正确
	if chainId == NULSChainId && prefix != NULSPrefix {
		return false
	}
	if chainId == TNULSChainId && prefix != TNULSPrefix {
		return false
	}
	accountType := bytes[2]
	if accountType > P2SHAccountType {
		return false
	}
	addressBytes := bytes[0 : len(bytes)-1]
	xor := calcXor(addressBytes)
	if xor != bytes[len(bytes)-1] {
		//校验位不正确
		return false
	}
	return true
}

//根据私钥生成账户
func GetAccountFromPrkey(priHex string, chainId uint16, prefix string) (*Account, error) {
	prikeyBytes, err := hex.DecodeString(priHex)
	if err != nil {
		return nil, err
	}
	return GetAccountFromPrkeyBytes(prikeyBytes, chainId, prefix)
}
func GetAccountFromPrkeyBytes(prikeyBytes []byte, chainId uint16, prefix string) (*Account, error) {
	ec, err := eckey.FromPriKeyBytes(prikeyBytes)
	if err != nil {
		return nil, err
	}
	return getAccountByEckey(ec, chainId, prefix)
}
