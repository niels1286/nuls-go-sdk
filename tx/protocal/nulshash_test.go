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
package txprotocal

import (
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"reflect"
	"testing"
)

func TestNulsHash_Parse(t *testing.T) {
	type fields struct {
		bytes   []byte
		hashHex string
	}
	type args struct {
		reader seria.ByteBufReader
	}
	data, _ := hex.DecodeString("b1fa764412178e1eb2cb92616de84300a50b4790f84901000000000000000000")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "testParse.a", fields: fields{bytes: []byte{}, hashHex: ""}, args: args{reader: seria.NewByteBufReader(data, 0)}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := &NulsHash{
				bytes:   tt.fields.bytes,
				hashHex: tt.fields.hashHex,
			}
			if err := hash.Parse(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if hash.String() != "b1fa764412178e1eb2cb92616de84300a50b4790f84901000000000000000000" {
				t.Fatalf("NulsHash Parse failed.")
			}
		})
	}
}

func TestNulsHash_Serialize(t *testing.T) {
	type fields struct {
		bytes   []byte
		hashHex string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "testSerialize.a", fields: fields{bytes: []byte{1, 2, 3, 4, 5, 6}, hashHex: ""}, want: []byte{1, 2, 3, 4, 5, 6}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := &NulsHash{
				bytes:   tt.fields.bytes,
				hashHex: tt.fields.hashHex,
			}
			got, err := hash.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNulsHash_String(t *testing.T) {
	type fields struct {
		bytes   []byte
		hashHex string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "TestString.a", fields: fields{
			bytes:   []byte{0, 1, 2, 3},
			hashHex: "",
		}, want: "00010203"},
		{name: "TestString.b", fields: fields{
			bytes:   []byte{0, 1, 2, 3},
			hashHex: "aaaaaaaa",
		}, want: "aaaaaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := &NulsHash{
				bytes:   tt.fields.bytes,
				hashHex: tt.fields.hashHex,
			}
			if got := hash.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
