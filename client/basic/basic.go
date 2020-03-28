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
package basic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//请求参数封装
type RequestParam struct {
	//固定值：2.0
	Jsonrpc string `json:"jsonrpc"`
	//接口名称
	Method string `json:"method"`
	//可以是slice，也可以是结构体
	Params interface{} `json:"params"`
	//请求的唯一标识，返回的结果中也包含该id
	Id int `json:"id"`
}

func (pp *RequestParam) ToJson() string {
	data, err := json.Marshal(pp)
	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return string(data)
}

func (pp *RequestParam) ToJsonBytes() []byte {
	data, err := json.Marshal(pp)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}
	return data
}

//请求返回结构体
type RequestResult struct {
	//固定值：2.0
	Jsonrpc string `json:"jsonrpc"`
	//对应请求的唯一标识
	Id string `json:"id"`

	Result interface{} `json:"result"`

	Error interface{} `json:"error"`
}

type JsonRpcClient interface {
	Request(param *RequestParam) (*RequestResult, error)
}

type BasicClient struct {
	client *http.Client
	url    string
}

//新建一个客户端，用来访问链上数据
func NewJSONRPCClient(url string) *BasicClient {
	return &BasicClient{
		client: &http.Client{},
		url:    url,
	}
}

//组装请求参数
//@id,请求的唯一标识，在返回的结构体中，也会包含本字段，并且等于请求时的值
//@method,请求的具体方法
//@params,实际参数，请参照接口文档进行组装
func NewRequestParam(id int, method string, params interface{}) *RequestParam {
	return &RequestParam{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      id,
	}
}

//接口请求
//请求的地址是client中的默认地址
//本工具针对NULS的api模块的jsonrpc接口进行设计，适用范围有限
func (c *BasicClient) Request(param *RequestParam) (*RequestResult, error) {
	if c.client == nil {
		c.client = &http.Client{}
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewReader(param.ToJsonBytes()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := c.client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &RequestResult{}
	json.Unmarshal(body, result)
	return result, nil
}
