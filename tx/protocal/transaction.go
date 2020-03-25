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

import "github.com/niels1286/nuls-go-sdk/utils/seria"

//基础交易结构，NULS生态中所有的交易都是已次结构组成的
type Transaction struct {
	//交易类型，@txtype.go中对网络支持的交易类型做了常量
	txType uint16
	//交易发生时间，精确到秒
	time uint32
	//交易的备注，默认为UTF-8编码的字符串
	remark string
	//交易业务扩展字段，任何需要上链的数据都可以放在这里
	txData []byte
	//资产交易数据，资产的转入、转出都需要在这里进行
	coinData []byte
	//交易签名数据，支持多个签名，每个签名包含一个公钥和一个签名数据
	sigData []byte
}

//将交易序列化为字节slice
func (t Transaction) serialize() ([]byte, error) {

}

//将字节slice反序列化为交易结构体
//@bytes 包含交易的完整的序列化数据的字节slice
//@cursor 游标，从此开始解析交易
func ParseTransaction(bytes []byte, cursor int) (Transaction, error) {
	return ParseTransactionByReader(seria.NewByteBufReader(bytes, cursor))
}

//从Reader中解析交易
//@reader 字节slice数据阅读器，其中包含交易的完整序列化数据，且cursor刚好处于交易数据的起始点
func ParseTransactionByReader(reader seria.ByteBufReader) (Transaction, error) {
	tx := Transaction{}
	tx.txType = reader.ReadUint16()

	return tx, nil
}
