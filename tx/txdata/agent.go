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
	"math/big"
)

//创建节点的交易扩展字段的具体协议
type Agent struct {
	//节点保证金金额
	Amount *big.Int
	//佣金比例
	CommissionRate uint8
	//节点创建地址
	AgentAddress []byte
	//节点奖励接收地址
	RewardAddress []byte
	//节点打包地址
	PackingAddress []byte
}

//反序列化
func (a *Agent) Parse(reader *seria.ByteBufReader) error {
	var err error
	a.Amount, err = reader.ReadBigInt()
	if err != nil {
		return err
	}
	a.AgentAddress, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	a.PackingAddress, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	a.RewardAddress, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	a.CommissionRate, err = reader.ReadByte()
	if err != nil {
		return err
	}
	return nil
}

//序列化方法
func (a *Agent) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBigint(a.Amount)
	writer.WriteBytes(a.AgentAddress)
	writer.WriteBytes(a.PackingAddress)
	writer.WriteBytes(a.RewardAddress)
	writer.WriteByte(a.CommissionRate)
	return writer.Serialize(), nil
}
