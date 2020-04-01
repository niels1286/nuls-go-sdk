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

// @Title bytebufwriter.go
// @Description 数据序列化辅助工具
// @Author  Niels  2020/3/26
package seria

import (
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"log"
	"math/big"
)

//序列化辅助工具，可以将各种数据类型的数据写入到byte slice中
//这个对象是线程不安全的，使用时需要避免并发调用
type ByteBufWriter struct {
	stream []byte
}

func NewByteBufWriter() *ByteBufWriter {
	return &ByteBufWriter{}
}

func (writer *ByteBufWriter) WriteByte(b byte) {
	writer.stream = append(writer.stream, b)
}
func (writer *ByteBufWriter) WriteBytes(bytes []byte) {
	writer.stream = append(writer.stream, bytes...)
}
func (writer *ByteBufWriter) WriteBytesWithLen(bytes []byte) {
	writer.stream = append(writer.stream, mathutils.VarIntToBytes(uint64(len(bytes)))...)
	writer.stream = append(writer.stream, bytes...)
}

func (writer *ByteBufWriter) WriteUInt16(val uint16) {
	writer.stream = append(writer.stream, mathutils.Uint16ToBytes(val)...)

}
func (writer *ByteBufWriter) WriteUInt32(val uint32) {
	writer.stream = append(writer.stream, mathutils.Uint32ToBytes(val)...)
}
func (writer *ByteBufWriter) WriteUInt64(val uint64) {
	writer.stream = append(writer.stream, mathutils.Uint64ToBytes(val)...)
}
func (writer *ByteBufWriter) WriteVarint(val uint64) {
	writer.stream = append(writer.stream, mathutils.VarIntToBytes(val)...)
}
func (writer *ByteBufWriter) WriteString(val string) {
	writer.WriteBytesWithLen([]byte(val))
}
func (writer *ByteBufWriter) WriteBool(val bool) {
	if val {
		writer.WriteByte(1)
	} else {
		writer.WriteByte(0)
	}
}

func (writer *ByteBufWriter) WriteBigint(val *big.Int) {
	bytes, err := mathutils.BigIntToBytes(val)
	if err != nil {
		log.Println(err.Error())
		return
	}
	writer.WriteBytes(bytes)
}

func (writer *ByteBufWriter) WriteFloat64(val float64) {
	writer.stream = append(writer.stream, mathutils.Float64ToBytes(val)...)
}
func (writer *ByteBufWriter) WriteNulsData(val NulsData) {
	bytes, err := val.Serialize()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	writer.stream = append(writer.stream, bytes...)
}

func (writer *ByteBufWriter) Serialize() []byte {
	return writer.stream
}
