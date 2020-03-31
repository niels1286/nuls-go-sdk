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
// @Author  Niels  2020/3/31
package txdata

import (
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
)

//红牌惩罚
type RedPunish struct {
	//被惩罚的账户地址，对应节点的创建地址
	Address []byte
	//红牌原因:1,尝试分叉网络，2，双花等恶意交易，3，信用值过低
	ReasonCode byte
	//证据
	Evidence []byte
}

//反序列化
func (y *RedPunish) Parse(reader *seria.ByteBufReader) error {
	var err error
	y.Address, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	y.ReasonCode, err = reader.ReadByte()
	if err != nil {
		return err
	}
	y.Evidence, err = reader.ReadBytesWithLen()
	if err != nil {
		return err
	}

}

//序列化方法
func (y *RedPunish) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBytes(y.Address)
	writer.WriteByte(y.ReasonCode)
	writer.WriteBytesWithLen(y.Evidence)
	return writer.Serialize(), nil
}
