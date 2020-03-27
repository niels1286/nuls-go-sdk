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
// @Author  Niels  2020/3/26
package assemble

import (
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"math/big"
	"time"
)

//交易的发出者，包含账户信息，资产信息，状态信息
type Sender struct {
	//转出账户
	//要求账户必须包含私钥
	Account account.Account
	//资产的链id，
	ChainId uint16
	//资产id
	AssetsId uint16
	//资产的数量，将小数位数左移到小数点之前，已整数的方式进行使用
	Amount *big.Int
	//账户当前的nonce值，
	//需要从区块链上获取，保证nonce值的有效性
	//Nonce值是该账户的上一条花费资产的交易hash的后8个字节
	//连续交易时，可以在应用端缓存nonce，用来组装下一个交易
	Nonce []byte
	//使用的资产是否是锁定资产
	//0普通交易，-1解锁金额交易（退出共识，退出委托）
	Locked byte
}

//交易的接收方，包含账户地址、资产信息、锁定信息
type Receiver struct {
	//接收方地址
	Address []byte
	//资产的链id
	ChainId uint16
	//资产id
	AssetsId uint16
	//资产数量
	Amount *big.Int
	//资产锁定标识
	//lockValue == -1时，代表业务锁定，只能通过对应的解锁交易进行解锁
	//lockValue < 10,000,000时，代表区块高度，当网络高度超过这个数值时，自动解锁（可以使用）
	//10,000,000<lockValue<1,000,000,000,000时，代表当前时间（秒），当时间晚于此值时，自动解锁（可以使用）
	//lockValue>1,000,000,000,000时，代表当前时间（毫秒），当时间晚于此值时，自动解锁（可以使用）
	LockValue uint64
}

//交易组装参数，用于普通组装转账交易
type TransferParams struct {
	//交易发送方信息
	Senders []Sender
	//交易接收方信息
	Receivers []Receiver
	//交易备注
	Remark string
	//交易扩展数据
	Extend []byte
}

//新建转账交易
func NewTransferTx(params *TransferParams) *txprotocal.Transaction {
	tx := &txprotocal.Transaction{}
	tx.TxType = txprotocal.TX_TYPE_TRANSFER
	tx.Time = uint32(time.Now().Second())
	if params.Remark != "" {
		tx.Remark = []byte(params.Remark)
	}
	tx.Extend = params.Extend
	tx.CoinData = NewCoinData(params.Senders, params.Receivers)
	tx.SignData = NewSignData(params.Senders, tx.GetHash())
	return tx
}

func NewSignData(senders []Sender, hash *txprotocal.NulsHash) []byte {

}

//根据sender和receiver生成coindata字节数据
func NewCoinData(senders []Sender, receivers []Receiver) []byte {
	cd := txprotocal.CoinData{}
	for _, sender := range senders {
		from := txprotocal.CoinFrom{}
		from.Address = sender.Account.AddressBytes
		from.AssetsChainId = sender.ChainId
		from.AssetsId = sender.AssetsId
		from.Amount = sender.Amount
		from.Nonce = sender.Nonce
		from.Locked = sender.Locked
		cd.Froms = append(cd.Froms, from)
	}
	for _, receiver := range receivers {
		to := txprotocal.CoinTo{
			Coin: txprotocal.Coin{
				Address:       receiver.Address,
				AssetsChainId: receiver.ChainId,
				AssetsId:      receiver.AssetsId,
				Amount:        receiver.Amount,
			},
			LockValue: receiver.LockValue,
		}
		cd.Tos = append(cd.Tos, to)
	}
	return cd
}
