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
package txprotocal

import (
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"reflect"
	"testing"
)

func TestSignData_Parse(t *testing.T) {
	pubHex := "029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e"
	//priHex := "e894bf09ffda5dfd808b3cea3612ff494fdb1e621bda691462e5e1d15322183d"
	signValueHex := "3045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	allHex := "21029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e473045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	pub, _ := hex.DecodeString(pubHex)
	sv, _ := hex.DecodeString(signValueHex)
	all, _ := hex.DecodeString(allHex)
	type args struct {
		reader *seria.ByteBufReader
	}
	tests := []struct {
		name    string
		s       SignData
		args    args
		wantErr bool
		want    SignData
	}{
		{name: "test sing parse", s: SignData{}, args: args{reader: seria.NewByteBufReader(all, 0)}, wantErr: false, want: SignData{[]P2PHKSignature{P2PHKSignature{
			SignValue: sv,
			PublicKey: pub,
		}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Parse(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("sign parse failed.")
			}
		})
	}
}

func TestSignData_Serialize(t *testing.T) {
	pubHex := "029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e"
	//priHex := "e894bf09ffda5dfd808b3cea3612ff494fdb1e621bda691462e5e1d15322183d"
	signValueHex := "3045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	allHex := "21029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e473045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	pub, _ := hex.DecodeString(pubHex)
	sv, _ := hex.DecodeString(signValueHex)
	all, _ := hex.DecodeString(allHex)
	tests := []struct {
		name    string
		s       SignData
		want    []byte
		wantErr bool
	}{
		{name: "sign serialize", s: SignData{[]P2PHKSignature{P2PHKSignature{
			SignValue: sv,
			PublicKey: pub,
		}}}, want: all, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Serialize()
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
