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
package commands

import (
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"math/rand"
	"time"
)

//获取指定链上指定地址的指定资产的相对应的余额和nonce状态
func GetBlockHex(client *jsonrpc.BasicClient, chainId uint16, height uint64) (string, error) {
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getBlockSerializationByHeight", []interface{}{chainId, height})
	result, err := client.Request(param)
	if err != nil {
		return "", err
	}
	blockHex := result.Result.(string)
	return blockHex, nil
}
