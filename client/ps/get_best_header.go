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
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"math/rand"
	"time"
)

type HeaderInfo struct {
	Hash    string `json:"hash"`
	Height  uint64 `json:"height"`
	TxCount int    `json:"txCount"`
}

//获取当前网络最新高度
func GetBestHeader(client *jsonrpc.NulsPSClient, chainId uint16) (*HeaderInfo, error) {
	if client == nil {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getBestBlockHeader", []interface{}{chainId})
	result, err := client.PSRequest(param)
	if err != nil {
		return nil, err
	}
	resultMap := result.Result.(map[string]interface{})
	header := &HeaderInfo{}
	err = mapstructure.Decode(resultMap, header)
	if err != nil {
		return nil, err
	}
	return header, nil
}
