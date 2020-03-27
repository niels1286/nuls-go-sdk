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
// @Author  Niels  2020/3/27
package account

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestAddressBytesToStr(t *testing.T) {
	address := "tNULSeBaMoodYW7AqyJrgYdWiJ6nfwfVHHHyXm"
	addressBytes, _ := hex.DecodeString("0200018f44b8662e78871f44ef1e1608282fd59560dcd0")
	type args struct {
		address []byte
		prefix  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test to string", args: args{address: addressBytes, prefix: "tNULS"}, want: address},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddressBytesToStr(tt.args.address, tt.args.prefix); got != tt.want {
				t.Errorf("AddressBytesToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressStrToBytes(t *testing.T) {
	address := "tNULSeBaMoodYW7AqyJrgYdWiJ6nfwfVHHHyXm"
	addressBytes, _ := hex.DecodeString("0200018f44b8662e78871f44ef1e1608282fd59560dcd0")
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "str to bytes", args: args{address: address}, want: addressBytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddressStrToBytes(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressStrToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
