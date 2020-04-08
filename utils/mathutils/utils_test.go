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
// @Author  Niels  2020/3/25
package mathutils

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestBytesToUint16(t *testing.T) {
	type args struct {
		array []byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "bytes to uint16.a", args: args{[]byte{103, 43}}, want: 11111},
		{name: "bytes to uint16.b", args: args{[]byte{0, 0}}, want: 0},
		{name: "bytes to uint16.c", args: args{[]byte{255, 255}}, want: 65535},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToUint16(tt.args.array); got != tt.want {
				t.Errorf("BytesToUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBigInt(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestStringToBigInt.a", args: args{"123456789987654321123456789987654321123456789987654321"}, want: "123456789987654321123456789987654321123456789987654321"},
		{name: "TestStringToBigInt.b", args: args{"0"}, want: "0"},
		{name: "TestStringToBigInt.c", args: args{"11111"}, want: "11111"},
		{name: "TestStringToBigInt.d", args: args{"2222222222222222222222222222"}, want: "2222222222222222222222222222"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := StringToBigInt(tt.args.val); got.String() != tt.want {
				t.Errorf("StringToBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16ToBytes(t *testing.T) {
	type args struct {
		n uint16
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "TestUint16ToBytes.a", args: args{0}, want: []byte{0, 0}},
		{name: "TestUint16ToBytes.b", args: args{65535}, want: []byte{255, 255}},
		{name: "TestUint16ToBytes.c", args: args{11111}, want: []byte{103, 43}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16ToBytes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32ToBytes(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "TestUint32ToBytes.a", args: args{0}, want: []byte{0, 0, 0, 0}},
		{name: "TestUint32ToBytes.b", args: args{1234567890}, want: []byte{210, 2, 150, 73}},
		{name: "TestUint32ToBytes.b", args: args{4294967295}, want: []byte{255, 255, 255, 255}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32ToBytes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64ToBytes(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "TestUint64ToBytes.a", args: args{0}, want: []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{name: "TestUint64ToBytes.b", args: args{9223372036854775807}, want: []byte{255, 255, 255, 255, 255, 255, 255, 127}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64ToBytes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToBigInt(t *testing.T) {
	type args struct {
		bytes []byte
	}

	bytes1, _ := hex.DecodeString("b1fa764412178e1eb2cb92616de84300a50b4790f84901000000000000000000")
	want1, _ := StringToBigInt("123456789987654321123456789987654321123456789987654321")
	fmt.Println(hex.EncodeToString(want1.Bytes()))
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{name: "bytestobigint.a", args: args{bytes1}, want: want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToBigInt(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToUint32(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "bytestouint32.a", args: args{bytes: Uint32ToBytes(1234567)}, want: 1234567},
		{name: "bytestouint32.b", args: args{bytes: Uint32ToBytes(0)}, want: 0},
		{name: "bytestouint32.c", args: args{bytes: Uint32ToBytes(10000)}, want: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToUint32(tt.args.bytes); got != tt.want {
				t.Errorf("BytesToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToUint64(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "bytestouint32.a", args: args{bytes: Uint64ToBytes(1234567890987654321)}, want: 1234567890987654321},
		{name: "bytestouint32.b", args: args{bytes: Uint64ToBytes(0)}, want: 0},
		{name: "bytestouint32.c", args: args{bytes: Uint64ToBytes(10000)}, want: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToUint64(tt.args.bytes); got != tt.want {
				t.Errorf("BytesToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToFloat64(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "toFload64.a", args: args{bytes: Float64ToBytes(123.321)}, want: 123.321},
		{name: "toFload64.b", args: args{bytes: Float64ToBytes(0)}, want: 0},
		{name: "toFload64.c", args: args{bytes: Float64ToBytes(123456789.123)}, want: 123456789.123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToFloat64(tt.args.bytes); got != tt.want {
				t.Errorf("BytesToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ToBytes(t *testing.T) {
	type args struct {
		float float64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "float64ToBytes.a", args: args{123.321}, want: Float64ToBytes(123.321)},
		{name: "float64ToBytes.b", args: args{123456789.321}, want: Float64ToBytes(123456789.321)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ToBytes(tt.args.float); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarIntToBytes(t *testing.T) {
	type args struct {
		varint uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "varinttobytes.a", args: args{varint: 100}, want: []byte{100}},
		{name: "varinttobytes.b", args: args{varint: 65536}, want: []byte{254, 0, 0, 1, 0}},
		{name: "varinttobytes.c", args: args{varint: 2147483647}, want: []byte{254, 255, 255, 255, 127}},
		{name: "varinttobytes.d", args: args{varint: 9223372036854775807}, want: []byte{255, 255, 255, 255, 255, 255, 255, 255, 127}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VarIntToBytes(tt.args.varint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VarIntToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToVarInt(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "bytesToVarInt.a", args: args{bytes: VarIntToBytes(0)}, want: 0},
		{name: "bytesToVarInt.b", args: args{bytes: VarIntToBytes(1)}, want: 1},
		{name: "bytesToVarInt.c", args: args{bytes: VarIntToBytes(65536)}, want: 65536},
		{name: "bytesToVarInt.d", args: args{bytes: VarIntToBytes(2147483647)}, want: 2147483647},
		{name: "bytesToVarInt.e", args: args{bytes: VarIntToBytes(9223372036854775807)}, want: 9223372036854775807},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToVarInt(tt.args.bytes); got != tt.want {
				t.Errorf("BytesToVarInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigIntToBytes(t *testing.T) {
	type args struct {
		val *big.Int
	}
	val, _ := StringToBigInt("123456789987654321123456789987654321123456789987654321")
	want, _ := hex.DecodeString("b1fa764412178e1eb2cb92616de84300a50b4790f84901000000000000000000")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "BigIntToBytes.a", args: args{val: val}, want: want, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BigIntToBytes(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("BigIntToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigIntToBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aliceReverse(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "aliceREverse.a", args: args{bytes: []byte{1, 2, 3, 4, 5}}, want: []byte{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aliceReverse(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aliceReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVarIntLen(t *testing.T) {
	type args struct {
		bytes  []byte
		cursor int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "varintlen.a", args: args{bytes: VarIntToBytes(100), cursor: 0}, want: 1},
		{name: "varintlen.b", args: args{bytes: VarIntToBytes(1000), cursor: 0}, want: 3},
		{name: "varintlen.c", args: args{bytes: VarIntToBytes(65536), cursor: 0}, want: 5},
		{name: "varintlen.d", args: args{bytes: VarIntToBytes(9223372036854775807), cursor: 0}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVarIntLen(tt.args.bytes, tt.args.cursor); got != tt.want {
				t.Errorf("GetVarIntLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringAmount(t *testing.T) {
	type args struct {
		val   *big.Int
		scale int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test_decimal.a", args: args{
			val:   big.NewInt(123456789),
			scale: 8,
		}, want: "1.23456789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringAmount(tt.args.val, tt.args.scale); got != tt.want {
				t.Errorf("GetStringAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
