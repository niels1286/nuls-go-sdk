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
package nuls

import (
	"github.com/niels1286/nuls-go-sdk/client/commands"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
)

type NulsSdk struct {
	client  *jsonrpc.BasicClient
	chainId uint16
}

func NewNulsSdk(jsonrpcURL string, chainId uint16) *NulsSdk {
	return &NulsSdk{client: jsonrpc.NewJSONRPCClient(jsonrpcURL), chainId: chainId}
}

//获取账户对应资产的余额
func (sdk *NulsSdk) GetBalance(address string, assetsChainId, assetsId int) (*commands.AccountStatus, error) {
	return commands.GetAccountInfo(sdk.client, address, sdk.chainId, assetsChainId, assetsId)
}

//获取api节点的网络连接状况
func (sdk *NulsSdk) GetNetworkInfo() (*commands.NetworkInfo, error) {
	return commands.GetNetworkInfo(sdk.client)
}

//根据高度获取区块hex
func (sdk *NulsSdk) GetBlockHex(height uint64) (string, error) {
	return commands.GetBlockHex(sdk.client, sdk.chainId, height)
}

//获取最新区块高度˚
func (sdk *NulsSdk) GetBestHeight() (uint64, error) {
	return commands.GetBestHeight(sdk.client, sdk.chainId)
}
