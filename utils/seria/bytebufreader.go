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

// @Title bytebufreader.go
// @Description  辅助数据反序列化的工具
// @Author  Niels  2020/3/25
package seria

import (
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"math/big"
)

//反序列化工具，可以从byte slice中读取各种数据类型的数据
type ByteBufReader struct {
	//完整的字节slice
	payload []byte
	//游标，标识着当前读取的位置
	cursor int
}

//创建一个读取器
func NewByteBufReader(payload []byte, cursor int) ByteBufReader {
	return ByteBufReader{payload, cursor}
}

//从字节序列中读取一个uint16
func (reader *ByteBufReader) ReadUint16() uint16 {

}

//从字节序列中读取一个uint32
func (reader *ByteBufReader) ReadUint32() uint32 {

}
func (reader *ByteBufReader) ReadVarInt() int {

}
func (reader *ByteBufReader) ReadUint64() uint64 {

}
func (reader *ByteBufReader) ReadByte() byte {

}
func (reader *ByteBufReader) ReadBytes(length int) []byte {

}

func (reader *ByteBufReader) ReadBytesWithLen() []byte {

}
func (reader *ByteBufReader) ReadStringWithLen() []string {

}
func (reader *ByteBufReader) ReadBool() bool {

}
func (reader *ByteBufReader) ReadBigInt() big.Int {

}

func (reader *ByteBufReader) ReadFloat32() float32 {

}

func (reader *ByteBufReader) ReadFloat64() float64 {

}

func (reader *ByteBufReader) ReadTx() txprotocal.Transaction {

}

func (reader *ByteBufReader) ReadHash() []byte {

}
func (reader ByteBufReader) GetPayload() []byte {
	return reader.payload
}
func (reader ByteBufReader) GetCursor() int {
	return reader.cursor
}
