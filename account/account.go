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
	"fmt"
	"github.com/niels1286/nerve-go-sdk/crypto/base58"
	"github.com/niels1286/nerve-go-sdk/crypto/eckey"
	cryptoutils "github.com/niels1286/nerve-go-sdk/crypto/utils"
	"github.com/niels1286/nerve-go-sdk/math"
	"strings"
)

const (
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
	//地址，是用户操作的载体，账户余额管理、转账、数据权限等，全有地址来识别
	//地址的格式为：address = prefix + Base58Encode(chainId+addressType+pkh+xor)
	address string
	//NULS是一个多链生态，每一条区块链，都有自己的id，用于在多链交互时，识别所属身份
	//NULS主网为1，NULS测试网为2
	chainId uint16
	//账户类型：1、普通账户，2、合约地址，3、多签地址
	accType uint8
	//账户对应的公私钥对
	eckey eckey.EcKey
	//地址前缀
	prefix string
}

//创建一个新账户
//chainId:账户所在链的ID，用于多链交互时，识别身份
//prefix:地址前缀，每条链都可以定义自己的地址前缀
func NewNormalAccount(chainId uint16, prefix string) (Account, error) {
	ec, err := eckey.NewEcKey()
	if err != nil {
		return Account{}, err
	}
	pubBytes := ec.GetPubKeyBytes(true)
	addressBytes := GetAddressByPubBytes(pubBytes, chainId, NormalAccountType, prefix)
	address := GetStringAddress(addressBytes, prefix)
	return Account{
		address: address,
		chainId: chainId,
		accType: NormalAccountType,
		eckey:   ec,
		prefix:  prefix,
	}, nil
}

//去除前缀，获得真正的地址字符串
func getRealAddress(address string) string {
	if strings.HasPrefix(address, "NULS") {
		return address[5:]
	}
	if strings.HasPrefix(address, "tNULS") {
		return address[6:]
	}
	for index, c := range address {
		if c >= 97 {
			return address[index+1:]
		}
	}
	return ""
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
	addressBytes = append(addressBytes, math.Uint16ToBytes(chainId)...)
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

func NewNULSAccount() (Account, error) {
	return NewNormalAccount(1, "NULS")
}

func ParseAccount(address string) Account {
	//todo
	return Account{}
}

func Valid(address string) bool {
	if address == "" {
		return false
	}
	realAddressStr := getRealAddress(address)
	fmt.Println(realAddressStr)
	//todo
	return false
}
