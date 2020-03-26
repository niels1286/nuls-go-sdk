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
// @Author  Niels  2020/3/26
package seria

import (
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"math/big"
	"reflect"
	"testing"
)

func TestByteBufWriter_writeBigint(t *testing.T) {
	type fields struct {
		stream []byte
	}
	value1, _ := mathutils.StringToBigInt("123456789987654321123456789987654321123456789987654321")
	want, _ := hex.DecodeString("b1fa764412178e1eb2cb92616de84300a50b4790f84901000000000000000000")
	type args struct {
		val *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeBigint.a", fields: fields{stream: []byte{}}, args: args{val: value1}, want: want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteBigint(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeBool(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeBool.a", fields: fields{stream: []byte{1, 2, 1, 2, 1, 2}}, args: args{val: true}, want: []byte{1, 2, 1, 2, 1, 2, 1}},
		{name: "writeBool.b", fields: fields{stream: []byte{1, 2, 1, 2, 1, 2}}, args: args{val: false}, want: []byte{1, 2, 1, 2, 1, 2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteBool(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeByte(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		b byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "WriteByte.a", fields: fields{stream: []byte{1, 2, 1, 2, 1, 2}}, args: args{b: 123}, want: []byte{1, 2, 1, 2, 1, 2, 123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteByte(tt.args.b)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeBytes(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeBytes.a", fields: fields{stream: []byte{}}, args: args{bytes: []byte{1, 2, 1, 2, 1, 2}}, want: []byte{1, 2, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteBytes(tt.args.bytes)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeBytesWithLen(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeBytesWithLen.a", fields: fields{stream: []byte{}}, args: args{bytes: []byte{1, 2, 1, 2, 1, 2}}, want: []byte{6, 1, 2, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteBytesWithLen(tt.args.bytes)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeFloat64(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeFloat64.a", fields: fields{stream: []byte{}}, args: args{val: 123456.789}, want: []byte{201, 118, 190, 159, 12, 36, 254, 64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteFloat64(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeString(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeString.a", fields: fields{stream: []byte{}}, args: args{val: "Nuls is a blockchain project."}, want: []byte{29, 78, 117, 108, 115, 32, 105, 115, 32, 97, 32, 98, 108, 111, 99, 107, 99, 104, 97, 105, 110, 32, 112, 114, 111, 106, 101, 99, 116, 46}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteString(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeUInt16(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{

		{name: "writeUInt16.a", fields: fields{stream: []byte{}}, args: args{val: 1}, want: []byte{1, 0}},
		{name: "writeUInt16.b", fields: fields{stream: []byte{}}, args: args{val: 1000}, want: []byte{232, 3}},
		{name: "writeUInt16.c", fields: fields{stream: []byte{}}, args: args{val: 65535}, want: []byte{255, 255}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteUInt16(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeUInt32(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{

		{name: "writeUInt32.a", fields: fields{stream: []byte{}}, args: args{val: 1}, want: []byte{1, 0, 0, 0}},
		{name: "writeUInt32.b", fields: fields{stream: []byte{}}, args: args{val: 1000}, want: []byte{232, 3, 0, 0}},
		{name: "writeUInt32.c", fields: fields{stream: []byte{}}, args: args{val: 65536}, want: []byte{0, 0, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteUInt32(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeUInt64(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writeUInt64.a", fields: fields{stream: []byte{}}, args: args{val: 1}, want: []byte{1, 0, 0, 0, 0, 0, 0, 0}},
		{name: "writeUInt64.b", fields: fields{stream: []byte{}}, args: args{val: 1000}, want: []byte{232, 3, 0, 0, 0, 0, 0, 0}},
		{name: "writeUInt64.c", fields: fields{stream: []byte{}}, args: args{val: 65536}, want: []byte{0, 0, 1, 0, 0, 0, 0, 0}},
		{name: "writeUInt64.d", fields: fields{stream: []byte{}}, args: args{val: 9223372036854775807}, want: []byte{255, 255, 255, 255, 255, 255, 255, 127}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteUInt64(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_writeVarint(t *testing.T) {
	type fields struct {
		stream []byte
	}
	type args struct {
		val uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{name: "writevarint.a", fields: fields{stream: []byte{}}, args: args{val: 1}, want: []byte{1}},
		{name: "writevarint.b", fields: fields{stream: []byte{}}, args: args{val: 65536}, want: []byte{254, 0, 0, 1, 0}},
		{name: "writevarint.c", fields: fields{stream: []byte{}}, args: args{val: 9223372036854775807}, want: []byte{255, 255, 255, 255, 255, 255, 255, 255, 127}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			writer.WriteVarint(tt.args.val)
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestByteBufWriter_serialize(t *testing.T) {
	type fields struct {
		stream []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{name: "serialize.a", fields: fields{stream: []byte{1, 2, 3, 4, 5, 6, 7, 8}}, want: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &ByteBufWriter{
				stream: tt.fields.stream,
			}
			if got := writer.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
