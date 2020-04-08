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
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"math/big"
)

//委托交易的交易扩展字段的具体协议
type Deposit struct {
	//委托金额
	Amount *big.Int
	//对应节点的hash
	AgentHash *txprotocal.NulsHash
	//委托者的地址
	Address []byte
}

//反序列化
func (d *Deposit) Parse(reader *seria.ByteBufReader) error {
	var err error
	d.Amount, err = reader.ReadBigInt()
	if err != nil {
		return err
	}
	d.Address, err = reader.ReadBytes(account.AddressBytesLength)
	if err != nil {
		return err
	}
	bytes, err := reader.ReadBytes(txprotocal.HashLength)
	if err != nil {
		return err
	}
	d.AgentHash = txprotocal.NewNulsHash(bytes)
	return nil
}

//序列化方法
func (d *Deposit) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBigint(d.Amount)
	writer.WriteBytes(d.Address)
	hash, _ := d.AgentHash.Serialize()
	writer.WriteBytes(hash)
	return writer.Serialize(), nil
}
