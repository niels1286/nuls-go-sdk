/*
 * MIT License
 * Copyright (c) 2019-2020 niels.wang
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package commands

import (
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"math/rand"
	"time"
)

type NetworkInfo struct {
	//本地高度
	LocalHeight int `json:""`
	//网络高度
	NetHeight int `json:""`
	//本地时间偏移值
	TimeOffset int `json:""`
	//主动连接节点数量
	OutCount int `json:""`
	//被动连接节点数量
	InCount int `json:""`
}

//获取api服务节点的网络状态
func GetNetworkInfo(client *jsonrpc.BasicClient) (*NetworkInfo, error) {

	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getNetworkInfo", []interface{}{})
	result, err := client.Request(param)
	if err != nil {
		return nil, err
	}
	resultMap := result.Result.(map[string]interface{})
	localHeight := resultMap["localBestHeight"].(int)
	netHeight := resultMap["netBestHeight"].(int)
	timeOffset := resultMap["timeOffset"].(int)
	outCount := resultMap["outCount"].(int)
	inCount := resultMap["inCount"].(int)

	return &NetworkInfo{
		LocalHeight: localHeight,
		NetHeight:   netHeight,
		TimeOffset:  timeOffset,
		OutCount:    outCount,
		InCount:     inCount,
	}, nil
}
