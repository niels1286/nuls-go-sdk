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
	"github.com/niels1286/nuls-go-sdk/account"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"reflect"
	"testing"
)

func TestCoinData_Serialize(t *testing.T) {
	cdHex := "021702000169103da2f21b1a7364566c4c5fe2c690241e952f01000100d202964900000000000000000000000000000000000000000000000000000000081111111111111111001702000169103da2f21b1a7364566c4c5fe2c690241e952f0200010023e88a08000000000000000000000000000000000000000000000000000000000800000000111111110002170200018f44b8662e78871f44ef1e1608282fd59560dcd001000100d2029649000000000000000000000000000000000000000000000000000000000000000000000000170200018f44b8662e78871f44ef1e1608282fd59560dcd00200010023bb5907000000000000000000000000000000000000000000000000000000000000000000000000"
	cdBytes, _ := hex.DecodeString(cdHex)
	cd := &CoinData{}
	cd.Parse(seria.NewByteBufReader(cdBytes, 0))
	tests := []struct {
		name       string
		fields     *CoinData
		want       []byte
		wantErr    bool
		wantAmount []string
	}{
		{name: "cd serialize.a", fields: cd, want: cdBytes, wantAmount: []string{"1234567890", "143321123", "1234567890", "123321123"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.Froms[0].Amount.String() != tt.wantAmount[0] ||
				tt.fields.Froms[1].Amount.String() != tt.wantAmount[1] ||
				tt.fields.Tos[0].Amount.String() != tt.wantAmount[2] ||
				tt.fields.Tos[1].Amount.String() != tt.wantAmount[3] {
				t.Errorf("CoinData serialize & parse failed.")
			}

			got, err := tt.fields.Serialize()
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

func Test_coinSerialize(t *testing.T) {
	type args struct {
		writer *seria.ByteBufWriter
		coin   Coin
	}
	coinHex, _ := hex.DecodeString("170200018f44b8662e78871f44ef1e1608282fd59560dcd001000100d202964900000000000000000000000000000000000000000000000000000000")
	val, _ := mathutils.StringToBigInt("1234567890")
	wantCoin := Coin{Address: account.AddressStrToBytes("tNULSeBaMoodYW7AqyJrgYdWiJ6nfwfVHHHyXm"), AssetsChainId: 1, AssetsId: 1, Amount: val}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test coin serialize.a", args: args{
			writer: seria.NewByteBufWriter(),
			coin:   wantCoin,
		}, want: coinHex},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coinSerialize(tt.args.writer, tt.args.coin)
			if !reflect.DeepEqual(tt.args.writer.Serialize(), coinHex) {
				t.Errorf("coin serializce failed.")
			}

		})
	}
}

func Test_parseCoin(t *testing.T) {
	type args struct {
		reader *seria.ByteBufReader
		coin   *Coin
	}
	coinHex, _ := hex.DecodeString("170200018f44b8662e78871f44ef1e1608282fd59560dcd001000100d202964900000000000000000000000000000000000000000000000000000000")
	val, _ := mathutils.StringToBigInt("1234567890")
	wantCoin := &Coin{Address: account.AddressStrToBytes("tNULSeBaMoodYW7AqyJrgYdWiJ6nfwfVHHHyXm"), AssetsChainId: 1, AssetsId: 1, Amount: val}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *Coin
	}{
		{name: "test parse coin.a", args: args{reader: seria.NewByteBufReader(coinHex, 0), coin: &Coin{}}, wantErr: false, want: wantCoin},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseCoin(tt.args.reader, tt.args.coin); (err != nil) != tt.wantErr {
				t.Errorf("parseCoin() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, tt.args.coin) {
				t.Errorf("parse coin failed.")
			}
		})
	}
}
