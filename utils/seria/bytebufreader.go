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
	"errors"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"math/big"
)

//反序列化工具，可以从byte slice中读取各种数据类型的数据
//这个对象是线程不安全的，使用时需要避免并发调用
type ByteBufReader struct {
	//完整的字节slice
	payload []byte
	//游标，标识着当前读取的位置
	cursor int
}

//创建一个读取器
func NewByteBufReader(payload []byte, cursor int) *ByteBufReader {
	return &ByteBufReader{payload, cursor}
}

//判断当前状态下，是否还可以继续读取length个字节的数据
func (reader *ByteBufReader) canRead(length int) bool {
	if length == 0 {
		return false
	}
	return length+reader.cursor <= len(reader.payload)
}

//从当前的字节序列中读取一个字节，并在游标加1
func (reader *ByteBufReader) ReadByte() (byte, error) {
	if !reader.canRead(1) {
		return 0, errors.New("Thers is not enough bytes.")
	}
	reader.cursor++
	return reader.payload[reader.cursor-1], nil
}

//从字节序列中读取指定数量的字节
func (reader *ByteBufReader) ReadBytes(length int) ([]byte, error) {
	if !reader.canRead(length) {
		return []byte{}, nil
	}
	bytes := reader.payload[reader.cursor : reader.cursor+length]
	reader.cursor += length
	return bytes, nil
}

//从字节序列中读取一个uint16,2个字节
func (reader *ByteBufReader) ReadUint16() (uint16, error) {
	length := 2
	if !reader.canRead(length) {
		return 0, nil
	}
	bytes := reader.payload[reader.cursor : reader.cursor+length]
	reader.cursor += length
	return mathutils.BytesToUint16(bytes), nil
}

//从字节序列中读取一个uint32
func (reader *ByteBufReader) ReadUint32() (uint32, error) {
	length := 4
	if !reader.canRead(length) {
		return 0, nil
	}
	bytes := reader.payload[reader.cursor : reader.cursor+length]
	reader.cursor += length
	return mathutils.BytesToUint32(bytes), nil
}

func (reader *ByteBufReader) ReadVarInt() (uint64, error) {
	length := mathutils.GetVarIntLen(reader.payload, reader.cursor)
	val := mathutils.BytesToVarIntByCursor(reader.payload, reader.cursor)
	reader.cursor += length
	return val, nil
}

func (reader *ByteBufReader) ReadUint64() (uint64, error) {
	length := 8
	if !reader.canRead(length) {
		return 0, nil
	}
	bytes := reader.payload[reader.cursor : reader.cursor+length]
	reader.cursor += length
	return mathutils.BytesToUint64(bytes), nil
}

func (reader *ByteBufReader) ReadBytesWithLen() ([]byte, error) {
	length, err := reader.ReadVarInt()
	if err != nil {
		return []byte{}, err
	}
	//存在数据精度丢失的可能性
	return reader.ReadBytes(int(length))
}

func (reader *ByteBufReader) ReadStringWithLen() (string, error) {
	bytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (reader *ByteBufReader) ReadBool() (bool, error) {
	val, err := reader.ReadByte()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

func (reader *ByteBufReader) ReadBigInt() (*big.Int, error) {
	length := 32
	if !reader.canRead(length) {
		return big.NewInt(0), nil
	}
	bytes := reader.payload[reader.cursor : reader.cursor+length]
	reader.cursor += length
	return mathutils.BytesToBigInt(bytes), nil
}

func (reader *ByteBufReader) ReadFloat64() (float64, error) {
	length := 8
	if !reader.canRead(length) {
		return 0, nil
	}
	bytes := reader.payload[reader.cursor : reader.cursor+length]
	reader.cursor += length
	return mathutils.BytesToFloat64(bytes), nil
}

func (reader *ByteBufReader) GetPayload() []byte {
	return reader.payload
}
func (reader *ByteBufReader) GetCursor() int {
	return reader.cursor
}

func (reader *ByteBufReader) IsFinished() bool {
	return len(reader.payload) == reader.cursor
}
