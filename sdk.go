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
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/client/api"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"github.com/niels1286/nuls-go-sdk/client/ps"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
)

type NulsSdk struct {
	apiClient *jsonrpc.NulsApiClient
	psClient  *jsonrpc.NulsPSClient
	chainId   uint16
}

func NewNulsSdk(apiUrl string, psUrl string, chainId uint16) *NulsSdk {
	return &NulsSdk{apiClient: jsonrpc.NewNulsApiClient(apiUrl), psClient: jsonrpc.NewNulsPSClient(psUrl), chainId: chainId}
}

//获取账户对应资产的余额
func (sdk *NulsSdk) GetBalance(address string, assetsChainId, assetsId uint16) (*api.AccountStatus, error) {
	return api.GetAccountInfo(sdk.apiClient, address, sdk.chainId, assetsChainId, assetsId)
}

//获取账户对应资产的余额
func (sdk *NulsSdk) GetNRC20Balance(address, contractAddress string) (*api.TokenBalance, error) {
	return api.GetNRC20Balance(sdk.apiClient, sdk.chainId, address, contractAddress)
}

//获取api节点的网络连接状况
func (sdk *NulsSdk) GetNetworkInfo() (*api.NetworkInfo, error) {
	return api.GetNetworkInfo(sdk.apiClient)
}

//根据高度获取区块hex
func (sdk *NulsSdk) GetBlockHex(height uint64) (string, error) {
	return api.GetBlockHex(sdk.apiClient, sdk.chainId, height)
}

//获取最新区块高度˚
func (sdk *NulsSdk) GetBestHeight() (uint64, error) {
	return api.GetBestHeight(sdk.apiClient, sdk.chainId)
}

//获取交易的详细信息
func (sdk *NulsSdk) GetTxJson(txHash *txprotocal.NulsHash) (string, error) {
	return api.GetTransactionJson(sdk.apiClient, sdk.chainId, txHash)
}

//根据地址获取节点信息
func (sdk *NulsSdk) GetAgentByAddress(address string) (*ps.AgentInfo, error) {
	return ps.GetAgentByAddress(sdk.psClient, sdk.chainId, address)
}

//获取全部节点信息
func (sdk *NulsSdk) GetAllAgent() ([]*ps.AgentInfo, error) {
	return ps.GetAllAgents(sdk.psClient, sdk.chainId)
}

//获取最新高度和hash
func (sdk *NulsSdk) GetBestHeader() (*ps.HeaderInfo, error) {
	return ps.GetBestHeader(sdk.psClient, sdk.chainId)
}

//广播交易
func (sdk *NulsSdk) BroadcastTx(bytes []byte) (string, error) {
	return api.BroadcastTx(sdk.apiClient, sdk.chainId, hex.EncodeToString(bytes))
}

//广播交易
func (sdk *NulsSdk) BroadcastTxHex(txHex string) (string, error) {
	return api.BroadcastTx(sdk.apiClient, sdk.chainId, txHex)
}

//验证交易
func (sdk *NulsSdk) ValidateTx(bytes []byte) (string, error) {
	return api.ValidateTx(sdk.apiClient, sdk.chainId, hex.EncodeToString(bytes))
}

//验证交易
func (sdk *NulsSdk) ValidateTxHex(txHex string) (string, error) {
	return api.ValidateTx(sdk.apiClient, sdk.chainId, txHex)
}
