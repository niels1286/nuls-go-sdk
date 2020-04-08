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
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
)

//退出委托的交易扩展字段的具体协议
type Withdraw struct {
	//委托交易的hash
	DepositTxHash *txprotocal.NulsHash
}

//反序列化
func (w *Withdraw) Parse(reader *seria.ByteBufReader) error {
	bytes, err := reader.ReadBytes(txprotocal.HashLength)
	if err != nil {
		return err
	}
	w.DepositTxHash = txprotocal.NewNulsHash(bytes)
	return nil
}

//序列化方法
func (w *Withdraw) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	hash, _ := w.DepositTxHash.Serialize()
	writer.WriteBytes(hash)
	return writer.Serialize(), nil
}
