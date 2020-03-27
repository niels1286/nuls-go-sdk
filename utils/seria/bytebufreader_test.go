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
package seria

import (
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"math/big"
	"reflect"
	"testing"
)

func TestByteBufReader_GetCursor(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "just test", fields: fields{[]byte{0}, 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.GetCursor(); got != tt.want {
				t.Errorf("GetCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_GetPayload(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{name: "just test 2", fields: fields{[]byte{0, 1}, 0}, want: []byte{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.GetPayload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPayload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadBigInt(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}

	bytes1, _ := hex.DecodeString("b1fa764412178e1eb2cb92616de84300a50b4790f84901000000000000000000")
	want1, _ := mathutils.StringToBigInt("123456789987654321123456789987654321123456789987654321")
	tests := []struct {
		name    string
		fields  fields
		want    *big.Int
		wantErr bool
	}{
		{name: "test bigint 1", fields: fields{bytes1, 0}, want: want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadBigInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBigInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBigInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadBool(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{name: "TestReadBool.a", fields: fields{[]byte{1}, 0}, want: true, wantErr: false},
		{name: "TestReadBool.b", fields: fields{[]byte{0}, 0}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadByte(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    byte
		wantErr bool
	}{
		{name: "readByte.a", fields: fields{[]byte{1}, 0}, want: 1, wantErr: false},
		{name: "readByte.b", fields: fields{[]byte{127}, 0}, want: 127, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadBytes(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "raedBytes.a", fields: fields{[]byte{1, 2, 3, 4, 5, 6}, 2}, args: args{2}, want: []byte{3, 4}, wantErr: false},
		{name: "raedBytes.b", fields: fields{[]byte{1, 2, 3, 4, 5, 6}, 0}, args: args{4}, want: []byte{1, 2, 3, 4}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadBytes(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadBytesWithLen(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "readbyteswithlen.a", fields: fields{[]byte{1, 5}, 0}, want: []byte{5}, wantErr: false},
		{name: "readbyteswithlen.b", fields: fields{[]byte{3, 5, 6, 7}, 0}, want: []byte{5, 6, 7}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadBytesWithLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBytesWithLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBytesWithLen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadFloat64(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{name: "readFloat64.a", fields: fields{[]byte{6, 129, 149, 67, 139, 212, 94, 64}, 0}, wantErr: false, want: 123.321},
		{name: "readFloat64.b", fields: fields{[]byte{29, 108, 177, 244, 16, 34, 177, 67}, 0}, wantErr: false, want: 1234567890987654321.321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadFloat64()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadStringWithLen(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "readString.a", fields: fields{[]byte{12, 72, 101, 108, 108, 111, 44, 119, 111, 114, 108, 100, 33}, 0}, want: "Hello,world!", wantErr: false},
		{name: "readString.b", fields: fields{[]byte{29, 78, 117, 108, 115, 32, 105, 115, 32, 97, 32, 98, 108, 111, 99, 107, 99, 104, 97, 105, 110, 32, 112, 114, 111, 106, 101, 99, 116, 46}, 0}, want: "Nuls is a blockchain project.", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadStringWithLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadStringWithLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadStringWithLen() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadUint16(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint16
		wantErr bool
	}{
		{name: "readUint16.a", fields: fields{[]byte{0, 0}, 0}, want: 0, wantErr: false},
		{name: "readUint16.b", fields: fields{[]byte{255, 255}, 0}, want: 65535, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadUint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadUint32(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint32
		wantErr bool
	}{
		{name: "readUint32.a", fields: fields{[]byte{0, 0, 0, 0}, 0}, want: 0, wantErr: false},
		{name: "readUint32.b", fields: fields{[]byte{255, 255, 255, 255}, 0}, want: 4294967295, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadUint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadUint64(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{name: "readUint64.a", fields: fields{[]byte{0, 0, 0, 0, 0, 0, 0, 0}, 0}, want: 0, wantErr: false},
		{name: "readUint64.b", fields: fields{[]byte{255, 255, 255, 255, 255, 255, 255, 127}, 0}, want: 9223372036854775807, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadUint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadVarInt(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		{name: "readvarint.a", fields: fields{[]byte{100}, 0}, want: 100, wantErr: false},
		{name: "readvarint.b", fields: fields{[]byte{254, 0, 0, 1, 0}, 0}, want: 65536, wantErr: false},
		{name: "readvarint.c", fields: fields{[]byte{254, 255, 255, 255, 127}, 0}, want: 2147483647, wantErr: false},
		{name: "readvarint.d", fields: fields{[]byte{255, 255, 255, 255, 255, 255, 255, 255, 127}, 0}, want: 9223372036854775807, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			got, err := reader.ReadVarInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadVarInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadVarInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_canRead(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	type args struct {
		length int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "canRead.a", fields: fields{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 0}, args: args{1}, want: true},
		{name: "canRead.b", fields: fields{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 0}, args: args{5}, want: true},
		{name: "canRead.c", fields: fields{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 0}, args: args{15}, want: false},
		{name: "canRead.d", fields: fields{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 8}, args: args{5}, want: false},
		{name: "canRead.e", fields: fields{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 5}, args: args{5}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.canRead(tt.args.length); got != tt.want {
				t.Errorf("canRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewByteBufReader(t *testing.T) {
	type args struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name string
		args args
		want ByteBufReader
	}{
		{name: "newReader", args: args{[]byte{1, 2, 3}, 0}, want: ByteBufReader{[]byte{1, 2, 3}, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewByteBufReader(tt.args.payload, tt.args.cursor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewByteBufReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_IsFinished(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "test isFinished.a", fields: fields{
			payload: []byte{1, 2, 3, 4},
			cursor:  4,
		}, want: true},
		{name: "test isFinished.a", fields: fields{
			payload: []byte{1, 2, 3, 4},
			cursor:  3,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.IsFinished(); got != tt.want {
				t.Errorf("IsFinished() = %v, want %v", got, tt.want)
			}
		})
	}
}
