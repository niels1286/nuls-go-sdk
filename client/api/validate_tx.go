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
	"encoding/json"
	"errors"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"math/rand"
	"time"
)

//将组装好的交易广播到网络中
//@Returns
//	string : 验证成功的交易的hash，如果为空则为失败
//	error  : 验证失败的提示信息
func ValidateTx(client *jsonrpc.NulsApiClient, chainId uint16, txhex string) (string, error) {
	if client == nil || txhex == "" {
		return "", errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "validateTx", []interface{}{chainId, txhex})
	result, err := client.ApiRequest(param)
	if err != nil {
		return "", err
	}
	if nil == result || nil == result.Result {
		if result != nil && result.Error != nil {
			bytes, _ := json.Marshal(result.Error)
			return "", errors.New(string(bytes))
		}
		return "", errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	value := resultMap["value"].(string)
	return value, nil
}
