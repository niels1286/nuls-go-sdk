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
// @Author  Niels  2020/3/27
package assemble

import (
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestNewCoinData(t *testing.T) {
	type args struct {
		senders   []Sender
		receivers []Receiver
	}
	priHex := "230cb8ebbf3a2c581d27f98f7a38f8b07c1ff170d605ca645db4ffa05ffa5505"
	acc, _ := account.GetAccountFromPrkey(priHex, account.NULSChainId, account.NULSPrefix)
	amount1, _ := mathutils.StringToBigInt("100100000")
	amount2, _ := mathutils.StringToBigInt("100000000")
	wantHex := "0117010001d6fc56e4dbf5417e9eb6041450872a600feddbe401000100a067f705000000000000000000000000000000000000000000000000000000000801020304050607080001170100019aa6bccb9e3cba60c95b409701b2417989da208b0100010000e1f505000000000000000000000000000000000000000000000000000000000000000000000000"
	want, _ := hex.DecodeString(wantHex)
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test new coindata.a", args: args{
			senders: []Sender{{
				Account:  acc,
				ChainId:  1,
				AssetsId: 1,
				Amount:   amount1,
				Nonce:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				Locked:   0,
			}},
			receivers: []Receiver{{
				Address:   account.AddressStrToBytes("NULSd6HgdNumANdW3LxB7NEZd4oa7otR4LkPN"),
				ChainId:   1,
				AssetsId:  1,
				Amount:    amount2,
				LockValue: 0,
			}},
		}, want: want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoinData(tt.args.senders, tt.args.receivers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoinData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSignData(t *testing.T) {
	priHex := "230cb8ebbf3a2c581d27f98f7a38f8b07c1ff170d605ca645db4ffa05ffa5505"
	acc, _ := account.GetAccountFromPrkey(priHex, account.NULSChainId, account.NULSPrefix)
	amount1, _ := mathutils.StringToBigInt("100100000")
	hashBytes, _ := txprotocal.ImportNulsHash("510a54c68cc64c6f131c8c7dc1ac59153d9b81e7dd1dc6f565f23ab2f9a6fcaf").Serialize()
	wantBytes, _ := hex.DecodeString("210233dd5281a4e129dafeea8637b54806f667e56f654e098c5faab87fa7fe889d1146304402203b26b0e057f5668af423b163125a2de5a7461066222bdb103574be5a304650e9022064dc1b9a59cfe22e871307f547834f8cfcb2de16f1322f3afdd376978f48f5d3")
	type args struct {
		senders   []Sender
		hashBytes []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "test new signData.a", args: args{
			senders: []Sender{{
				Account:  acc,
				ChainId:  1,
				AssetsId: 1,
				Amount:   amount1,
				Nonce:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				Locked:   0,
			}},
			hashBytes: hashBytes,
		}, want: wantBytes},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSignData(tt.args.senders, tt.args.hashBytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSignData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransferTx(t *testing.T) {
	priHex := "230cb8ebbf3a2c581d27f98f7a38f8b07c1ff170d605ca645db4ffa05ffa5505"
	acc, _ := account.GetAccountFromPrkey(priHex, account.NULSChainId, account.NULSPrefix)
	amount1, _ := mathutils.StringToBigInt("100100000")
	amount2, _ := mathutils.StringToBigInt("100000000")
	type args struct {
		params *TransferParams
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test new tx.a", args: args{params: &TransferParams{
			Senders: []Sender{{
				Account:  acc,
				ChainId:  1,
				AssetsId: 1,
				Amount:   amount1,
				Nonce:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
				Locked:   0,
			}},
			Receivers: []Receiver{{
				Address:   account.AddressStrToBytes("NULSd6HgdNumANdW3LxB7NEZd4oa7otR4LkPN"),
				ChainId:   1,
				AssetsId:  1,
				Amount:    amount2,
				LockValue: 0,
			}},
			TimeUnix: 1111111111, //uint32(time.Now().Unix()),
			Remark:   "test create transfer tx",
			Extend:   []byte("test extend data."),
		}}, want: "0200c7353a42177465737420637265617465207472616e73666572207478117465737420657874656e6420646174612e8c0117010001d6fc56e4dbf5417e9eb6041450872a600feddbe401000100a067f705000000000000000000000000000000000000000000000000000000000801020304050607080001170100019aa6bccb9e3cba60c95b409701b2417989da208b0100010000e1f50500000000000000000000000000000000000000000000000000000000000000000000000069210233dd5281a4e129dafeea8637b54806f667e56f654e098c5faab87fa7fe889d11463044022063369d66d7e72ec11c9c262be027876e6d25e366d7137f67968dc24bfdabeb9202206fe08e179e83d69c14f35e82d6d323bf30424df4e43656b2ae54956ab53707a7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gottx := NewTransferTx(tt.args.params)
			bytes, _ := gottx.Serialize()
			got := hex.EncodeToString(bytes)
			if got != tt.want {
				t.Errorf("NewTransferTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

//组装转账交易示例
func ExampleNewTransferTx() {
	//需求：用户a需要转100个nuls给用户B，花费手续费0.001个nuls。交易组装过程如下：
	//step 1: 首先要准备账户A的私钥或者keystore，保证账户A，拥有签名能力
	priHex := "230cb8ebbf3a2c581d27f98f7a38f8b07c1ff170d605ca645db4ffa05ffa5505"
	accountA, _ := account.GetAccountFromPrkey(priHex, account.NULSChainId, account.NULSPrefix)
	//step 2: 准备接收人地址，账户B地址
	addressB := "NULSd6HgdNumANdW3LxB7NEZd4oa7otR4LkPN"
	//step 3: 确认账户A，有足够转账的资产余额，并获取当前的nonce值
	//todo 未来通过client中的请求，从节点处获取真实数据，此处暂时模拟
	nonce := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	//step 4 :准备组装参数
	//提前准备好两个数值，nuls的小数位数为8位，所以100.001(加手续费)个NULS，用bigint表示，就是下面情况：
	amount1, _ := mathutils.StringToBigInt("10000100000")
	amount2, _ := mathutils.StringToBigInt("10000000000")
	params := &TransferParams{
		Senders: []Sender{{
			Account:  accountA,
			ChainId:  1,
			AssetsId: 1,
			Amount:   amount1,
			Nonce:    nonce,
			Locked:   0,
		}},
		Receivers: []Receiver{{
			Address:   account.AddressStrToBytes(addressB),
			ChainId:   1,
			AssetsId:  1,
			Amount:    amount2,
			LockValue: 0,
		}},
		//这里是系统时间，代表交易发生时间
		TimeUnix: uint32(time.Now().Unix()),
		//备注可以设置为业务数据，也可以是任何字符串
		Remark: "test create transfer tx",
		//在做系统扩展是，可以通过Extend进行业务数据的存储
		Extend: []byte("test extend data."),
	}
	//调用交易生成函数
	tx := NewTransferTx(params)
	//step 5:将交易转换为hex
	bytes, err := tx.Serialize()
	if err != nil {
		//序列化失败，则组装失败
		log.Println("Transaction create failed,case:" + err.Error())
		return
	}
	txHex := hex.EncodeToString(bytes)
	//step 6:调用api接口，将交易广播到区块链网络中
	//todo 待实现，用打印进行模拟
	log.Println("New transfer tx:" + txHex)
	//step 7: 等待交易确认，在交易确认后，应用可以进行后续业务操作
	//Done.
}
