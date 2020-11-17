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
package api

import (
	"encoding/hex"
	"errors"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"math/big"
	"math/rand"
	"time"
)

type AccountStatus struct {
	//账户地址
	Address string
	//可用余额
	Balance *big.Int
	//当前nonce值
	Nonce []byte
	//该nonce值的类型，1：已确认的nonce值,0：未确认的nonce值
	NonceType int
	//总的余额：可用+锁定
	TotalBalance *big.Int
}

type TokenBalance struct {
	ContractAddress string
	Name            string
	Symbol          string
	Amount          *big.Int
	Decimals        int64
	Status          int
}

//获取指定链上指定地址的指定资产的相对应的余额和nonce状态
func GetAccountInfo(client *jsonrpc.NulsApiClient, address string, chainId uint16, assetsChainId, assetsId uint16) (*AccountStatus, error) {
	if client == nil || address == "" {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getAccountBalance", []interface{}{chainId, assetsChainId, assetsId, address})
	result, err := client.ApiRequest(param)
	if err != nil {
		return nil, err
	}
	if nil == result || nil == result.Result {
		return nil, errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	balance, err := mathutils.StringToBigInt(resultMap["balance"].(string))
	if err != nil {
		return nil, err
	}
	nonceHex := resultMap["nonce"].(string)
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return nil, err
	}
	nonceType := resultMap["nonceType"].(float64)
	totalBalance, err := mathutils.StringToBigInt(resultMap["totalBalance"].(string))
	if err != nil {
		return nil, err
	}
	return &AccountStatus{
		Address:      address,
		Balance:      balance,
		Nonce:        nonce,
		NonceType:    int(nonceType),
		TotalBalance: totalBalance,
	}, nil
}

func GetNRC20Balance(client *jsonrpc.NulsApiClient, chainId uint16, address, contractAddress string) (*TokenBalance, error) {
	if client == nil || address == "" {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getTokenBalance", []interface{}{chainId, contractAddress, address})
	result, err := client.ApiRequest(param)
	if err != nil {
		return nil, err
	}
	if nil == result || nil == result.Result {
		return nil, errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	balance, err := mathutils.StringToBigInt(resultMap["amount"].(string))
	if err != nil {
		return nil, err
	}
	name := resultMap["name"].(string)
	symbol := resultMap["symbol"].(string)
	decimals := resultMap["decimals"].(float64)
	status := resultMap["status"].(float64)

	if err != nil {
		return nil, err
	}
	return &TokenBalance{
		ContractAddress: contractAddress,
		Name:            name,
		Symbol:          symbol,
		Amount:          balance,
		Decimals:        int64(decimals),
		Status:          int(status),
	}, nil
}
