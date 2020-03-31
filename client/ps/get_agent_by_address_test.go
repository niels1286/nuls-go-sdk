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
// @Author  Niels  2020/3/31
package ps

import (
	"fmt"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"testing"
)

func TestGetAgentByAddress(t *testing.T) {
	type args struct {
		client  *jsonrpc.NulsPSClient
		chainId uint16
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "getAgentByAddress.ps.a", args: args{
			client:  jsonrpc.NewNulsPSClient("https://public1.nuls.io"),
			chainId: 1,
			address: "NULSd6HgZkPDuWG7vZP8yQiLwMSxEFzY1rUNr",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAgentByAddress(tt.args.client, tt.args.chainId, tt.args.address)
			if err != nil {
				t.Errorf("GetAgentByAddress() error = %v, wantErr false", err)
				return
			}
			fmt.Println(got)
		})
	}
}
