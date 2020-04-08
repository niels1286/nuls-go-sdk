/*
 * MIT License
 * Copyright (c) 2019-2020 niels.wang
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package mathutils

import (
	"encoding/binary"
	"errors"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}

func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

func StringToBigInt(val string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(val, 10)
	if !ok {
		return big.NewInt(0), errors.New("String to big int failed.")
	}
	return n, nil
}

func BytesToBigInt(bytes []byte) *big.Int {
	i := new(big.Int)
	i.SetBytes(aliceReverse(bytes))
	return i
}

const BigIntBytesLength = 32

func BigIntToBytes(val *big.Int) ([]byte, error) {
	bytes := val.Bytes()
	bytes = aliceReverse(bytes)
	length := len(bytes)
	if length > BigIntBytesLength {
		return []byte{}, errors.New("It's too long(big.int bytes)")
	} else if length == BigIntBytesLength {
		return bytes, nil
	}
	for i := BigIntBytesLength - length; i > 0; i-- {
		bytes = append(bytes, 0)
	}
	return bytes, nil
}

//反转slice中的元素顺序
func aliceReverse(bytes []byte) []byte {
	newBytes := []byte{}
	for i := len(bytes) - 1; i >= 0; i-- {
		newBytes = append(newBytes, bytes[i])
	}
	return newBytes
}

func BytesToUint32(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}

func BytesToUint16(bytes []byte) uint16 {
	return binary.LittleEndian.Uint16(bytes)
}

func BytesToUint64(bytes []byte) uint64 {
	return binary.LittleEndian.Uint64(bytes)
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//Float64ToByte Float64转byte
func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func VarIntToBytes(value uint64) []byte {
	//bytes := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0,0}
	//count := binary.PutUvarint(bytes, varint)
	//return bytes[0:count]
	bytes := []byte{}
	if value < 253 {
		bytes = append(bytes, uint8(value))
		return bytes
	}
	if value <= 0xFFFF {
		// 1 marker + 2 entity bytes
		bytes = append(bytes, 253)
		bytes = append(bytes, Uint16ToBytes(uint16(value))...)
		return bytes
	}
	if value <= 0xFFFFFFFF {
		// 1 marker + 4 entity bytes
		bytes = append(bytes, 254)
		bytes = append(bytes, Uint32ToBytes(uint32(value))...)
		return bytes
	}
	// 1 marker + 8 entity bytes
	bytes = append(bytes, 255)
	bytes = append(bytes, Uint64ToBytes(uint64(value))...)
	return bytes
}

func BytesToVarInt(bytes []byte) uint64 {
	return BytesToVarIntByCursor(bytes, 0)
}
func BytesToVarIntByCursor(bytes []byte, cursor int) uint64 {
	//val, _ := binary.Uvarint(bytes)
	//return val
	var value uint64
	first := bytes[cursor]
	if first < 253 {
		value = uint64(first)
	} else if first == 253 {
		value = uint64(BytesToUint16(bytes[cursor+1 : cursor+3]))
	} else if first == 254 {
		value = uint64(BytesToUint32(bytes[cursor+1 : cursor+5]))
	} else {
		value = uint64(BytesToUint64(bytes[cursor+1 : cursor+9]))
	}
	return value
}

func GetVarIntLen(bytes []byte, cursor int) int {
	first := bytes[cursor]
	if first < 253 {
		return 1
	} else if first == 253 {
		return 3
	} else if first == 254 {
		return 5
	}
	return 9
}

func GetStringAmount(val *big.Int, scale int) string {
	decimal.DivisionPrecision = scale
	value := decimal.NewFromBigInt(val, 0)
	value = value.Div(decimal.NewFromBigInt(Pow(big.NewInt(10), scale), 0))
	return value.Round(int32(scale)).String()
}

func Pow(x *big.Int, n int) *big.Int {
	ret := big.NewInt(1) // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret.Mul(ret, x)
		}
		n /= 2
		x = x.Mul(x, x)
	}
	return ret
}
