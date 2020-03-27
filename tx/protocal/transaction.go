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

// @Title transaction.go
// @Description  定义交易的基本结构和常用方法
// @Author  Niels  2020/3/25
package txprotocal

import (
	cryptoutils "github.com/niels1286/nuls-go-sdk/crypto/utils"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
)

//基础交易结构，NULS生态中所有的交易都是已次结构组成的
type Transaction struct {
	//交易hash，唯一标识
	hash *NulsHash
	//交易类型，@txtype.go中对网络支持的交易类型做了常量
	TxType uint16
	//交易发生时间，精确到秒
	Time uint32
	//交易的备注，默认为UTF-8编码的字符串
	Remark []byte
	//交易业务扩展字段，任何需要上链的数据都可以放在这里
	Extend []byte
	//资产交易数据，资产的转入、转出都需要在这里进行
	CoinData []byte
	//交易签名数据，支持多个签名，每个签名包含一个公钥和一个签名数据
	SignData []byte
}

//获取交易hash
func (t *Transaction) GetHash() *NulsHash {
	if t.hash == nil || len(t.hash.bytes) == 0 {
		t.CalcHash()
	}
	return t.hash
}

//计算交易本身的hash，交易hash通过TxType，Time，Remark，Extend，CoinData共5个字段进行计算
func (t *Transaction) CalcHash() error {
	bytes, err := t.SerializeForHash()
	if err != nil {
		return err
	}
	bytes = cryptoutils.Sha256twice(bytes)
	t.hash = NewNulsHash(bytes)
	return nil
}

//将交易序列化为字节slice
func (t *Transaction) Serialize() ([]byte, error) {
	return t.serialize(true)
}

func (t *Transaction) SerializeForHash() ([]byte, error) {
	return t.serialize(false)
}

func (t *Transaction) serialize(withSign bool) ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteUInt16(t.TxType)
	writer.WriteUInt32(t.Time)
	writer.WriteBytesWithLen(t.Remark)
	writer.WriteBytesWithLen(t.Extend)
	writer.WriteBytesWithLen(t.CoinData)
	if withSign {
		writer.WriteBytesWithLen(t.SignData)
	}
	return writer.Serialize(), nil
}

func (t *Transaction) Parse(reader *seria.ByteBufReader) error {
	txType, err := reader.ReadUint16()
	if err != nil {
		return err
	}
	t.TxType = txType
	time, err := reader.ReadUint32()
	if err != nil {
		return err
	}
	t.Time = time
	remarkBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	t.Remark = remarkBytes

	dataBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	t.Extend = dataBytes

	coinBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	t.CoinData = coinBytes

	signBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	t.SignData = signBytes
	return nil
}

//将字节slice反序列化为交易结构体
//@bytes 包含交易的完整的序列化数据的字节slice
//@cursor 游标，从此开始解析交易
func ParseTransaction(bytes []byte, cursor int) *Transaction {
	return ParseTransactionByReader(seria.NewByteBufReader(bytes, cursor))
}

//从Reader中解析交易
//@reader 字节slice数据阅读器，其中包含交易的完整序列化数据，且cursor刚好处于交易数据的起始点
func ParseTransactionByReader(reader *seria.ByteBufReader) *Transaction {
	tx := Transaction{}
	tx.Parse(reader)
	return &tx
}
