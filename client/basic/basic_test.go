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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {

	user := &RequestParam{
		Jsonrpc: "2.0",
		Method:  "getAccountBalance",
		Params:  []interface{}{1, 1, 1, "NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"},
		Id:      123,
	}

	data, err := json.Marshal(user)

	if err != nil {
		t.Errorf(err.Error())
	}

	if `{"jsonrpc":"2.0","method":"getAccountBalance","params":[1,1,1,"NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"],"id":123}` != string(data) {
		t.Errorf("to json failed.")
	}
}

func TestBasicClient_Request(t *testing.T) {
	type fields struct {
		client *http.Client
		url    string
	}
	type args struct {
		param *RequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "test request.a", fields: fields{
			client: &http.Client{},
			url:    "https://api.nuls.io/jsonrpc",
		}, args: args{param: NewRequestParam(123, "getAccountBalance", []interface{}{1, 1, 1, "NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"})},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BasicClient{
				client: tt.fields.client,
				url:    tt.fields.url,
			}
			got, err := c.Request(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Result == nil {
				t.Errorf("Request() got = %v", got)
			}
		})
	}
}

func TestNewHttpClient(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *BasicClient
	}{
		{name: "test new client.a", args: args{"https://api.nuls.io/jsonrpc"}, want: &BasicClient{
			url:    "https://api.nuls.io/jsonrpc",
			client: &http.Client{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONRPCClient(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONRPCClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRequestParam(t *testing.T) {
	type args struct {
		id     int
		method string
		params interface{}
	}
	tests := []struct {
		name string
		args args
		want *RequestParam
	}{
		{name: "test new params.a", args: args{
			id:     123,
			method: "getAccountBalance",
			params: []interface{}{1, 1, 1, "NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"},
		}, want: &RequestParam{
			Id:      123,
			Method:  "getAccountBalance",
			Params:  []interface{}{1, 1, 1, "NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"},
			Jsonrpc: "2.0",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequestParam(tt.args.id, tt.args.method, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequestParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestParam_ToJson(t *testing.T) {
	type fields struct {
		jsonrpc string
		method  string
		params  interface{}
		id      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "test to json.a", fields: fields{
			jsonrpc: "2.0",
			method:  "getAccountBalance",
			params:  []interface{}{1, 1, 1, "NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"},
			id:      123,
		}, want: `{"jsonrpc":"2.0","method":"getAccountBalance","params":[1,1,1,"NULSd6HgcfGtsmm79QDoBK1MAjqNmm3rgKXSj"],"id":123}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := &RequestParam{
				Jsonrpc: tt.fields.jsonrpc,
				Method:  tt.fields.method,
				Params:  tt.fields.params,
				Id:      tt.fields.id,
			}
			if got := pp.ToJson(); got != tt.want {
				t.Errorf("ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestExample(t *testing.T) {
	ExampleBasicClient_Request()
}
func ExampleBasicClient_Request() {
	//step 1:首先创建一个client,用于访问测试网公共服务
	client := NewJSONRPCClient("http://beta.api.nuls.io/jsonrpc")
	//step 2: 组装参数
	params := NewRequestParam(
		//id，是一个随机数，用于标识本次请求
		1286,
		//将要请求的方法
		"getAccountBalance",
		//实际的参数，可以参考接口文档获得
		[]interface{}{2, 2, 1, "tNULSeBaMoG1oaW1JZnh6Ly65Ttp6raeTFBfCG"})
	//step 3: 调用请求
	result, err := client.Request(params)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	//step 4:处理结果，示例中只进行打印，如下（请求时间：2020-3-28)）
	val, _ := json.Marshal(result)
	fmt.Println(string(val))
	//结果示例，如下：
	//{
	//	"jsonrpc": "2.0",
	//	"id": "1286",
	//	"result": {
	//	"balance": "2967697605700",
	//		"consensusLock": "0",
	//		"freeze": "0",
	//		"nonce": "d311279375bcc2d7",
	//		"nonceType": 1,
	//		"timeLock": "0",
	//		"totalBalance": "2967697605700"
	//},
	//	"error": null
	//}
}
