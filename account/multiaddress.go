/*
 *  MIT License
 *  Copyright (c) 2019-2020 niels.wang
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 *
 */

// @Title
// @Description
// @Author  Niels  2020/4/8
package account

import (
	"encoding/hex"
	"fmt"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"sort"
)

//生成多签地址字符串的工具方法
//chainId : 链标识
//m ：最小签名数量
//pxHexSlice : 公钥字符串，每个公钥间用','分隔
//prefix：地址前缀
func CreateMultiAddress(chainId uint16, m uint8, pkHexSlice []string, prefix string) string {
	sort.Slice(pkHexSlice, func(i, j int) bool {
		return pkHexSlice[i] < pkHexSlice[j]
	})
	writer := seria.NewByteBufWriter()
	writer.WriteByte(byte(chainId))
	writer.WriteByte(m)
	for _, pk := range pkHexSlice {
		bytes, err := hex.DecodeString(pk)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		writer.WriteBytes(bytes)
	}
	bytes := writer.Serialize()
	addressBytes := GetAddressByPubBytes(bytes, chainId, P2SHAccountType, prefix)
	return GetStringAddress(addressBytes, prefix)
}
