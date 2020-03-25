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
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		want   big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadBigInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBigInt() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadBool(); got != tt.want {
				t.Errorf("ReadBool() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadByte(); got != tt.want {
				t.Errorf("ReadByte() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadBytes(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBytes() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadBytesWithLen(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBytesWithLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadFloat32(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadFloat32(); got != tt.want {
				t.Errorf("ReadFloat32() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadFloat64(); got != tt.want {
				t.Errorf("ReadFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadHash(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadHash() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadStringWithLen(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadStringWithLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBufReader_ReadTx(t *testing.T) {
	type fields struct {
		payload []byte
		cursor  int
	}
	tests := []struct {
		name   string
		fields fields
		want   txprotocal.Transaction
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadTx(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadTx() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadUint16(); got != tt.want {
				t.Errorf("ReadUint16() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadUint32(); got != tt.want {
				t.Errorf("ReadUint32() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadUint64(); got != tt.want {
				t.Errorf("ReadUint64() = %v, want %v", got, tt.want)
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
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &ByteBufReader{
				payload: tt.fields.payload,
				cursor:  tt.fields.cursor,
			}
			if got := reader.ReadVarInt(); got != tt.want {
				t.Errorf("ReadVarInt() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewByteBufReader(tt.args.payload, tt.args.cursor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewByteBufReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
