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
// @Author  Niels  2020/3/28
package api

import (
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	type args struct {
		client        *jsonrpc.NulsApiClient
		address       string
		chainId       uint16
		assetsChainId int
		assetsId      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test get account info.a", args: args{
			client:        jsonrpc.NewNulsApiClient("http://beta.api.nuls.io/jsonrpc"),
			address:       "tNULSeBaMoG1oaW1JZnh6Ly65Ttp6raeTFBfCG",
			chainId:       2,
			assetsChainId: 2,
			assetsId:      1,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccountInfo(tt.args.client, tt.args.address, tt.args.chainId, tt.args.assetsChainId, tt.args.assetsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Balance.String() == "" {
				t.Errorf("GetAccountInfo() failed.")
			}
		})
	}
}
