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
package txprotocal

import (
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"math/big"
)

//trandaction的coindata的预定义结构体
type CoinData struct {
	//转出数据结构，可以包含多个账户，多种资产
	//每个账户可以重复出现多次，但必须是不同的资产
	//from中的每个账户都必须对交易签名，否则交易不合法
	Froms []CoinFrom
	//转出数据结构，可以包含多个账户，多种资产
	Tos []CoinTo
}

func (c *CoinData) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteVarint(uint64(len(c.Froms)))
	for i := 0; i < len(c.Froms); i++ {
		from := c.Froms[i]
		coinSerialize(writer, from.Coin)
		writer.WriteBytesWithLen(from.Nonce)
		writer.WriteByte(from.Locked)
	}
	writer.WriteVarint(uint64(len(c.Tos)))
	for i := 0; i < len(c.Tos); i++ {
		to := c.Tos[i]
		coinSerialize(writer, to.Coin)
		writer.WriteUInt64(to.LockValue)
	}
	return writer.Serialize(), nil
}
func (c *CoinData) Parse(reader *seria.ByteBufReader) error {
	fromCount, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	for i := 0; i < int(fromCount); i++ {
		from := CoinFrom{}
		err := parseCoin(reader, &from.Coin)
		if err != nil {
			return err
		}
		from.Nonce, err = reader.ReadBytesWithLen()
		if err != nil {
			return err
		}
		from.Locked, err = reader.ReadByte()
		if err != nil {
			return err
		}
		c.Froms = append(c.Froms, from)
	}
	toCount, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	for i := 0; i < int(toCount); i++ {
		to := CoinTo{}
		err := parseCoin(reader, &to.Coin)
		if err != nil {
			return err
		}
		to.LockValue, err = reader.ReadUint64()
		if err != nil {
			return err
		}
		c.Tos = append(c.Tos, to)
	}
	return nil
}

//从reader中读取Coin数据
func parseCoin(reader *seria.ByteBufReader, coin *Coin) error {
	var err error
	coin.Address, err = reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	coin.AssetsChainId, err = reader.ReadUint16()
	if err != nil {
		return err
	}

	coin.AssetsId, err = reader.ReadUint16()
	if err != nil {
		return err
	}
	coin.Amount, err = reader.ReadBigInt()
	if err != nil {
		return err
	}
	return nil
}

//将Coin数据，序列化到Writer中
func coinSerialize(writer *seria.ByteBufWriter, coin Coin) {
	writer.WriteBytesWithLen(coin.Address)
	writer.WriteUInt16(coin.AssetsChainId)
	writer.WriteUInt16(coin.AssetsId)
	writer.WriteBigint(coin.Amount)
}

//资产信息结构体
type Coin struct {
	//账户地址
	Address []byte
	//资产发行链的id
	AssetsChainId uint16
	//资产的唯一标识
	AssetsId uint16
	//资产的数量，将小数位左移为整数
	Amount *big.Int
}

//资产转出信息结构体
type CoinFrom struct {
	Coin
	//账户状态的nonce值，用来避免双花交易
	Nonce []byte
	//使用的资产是否是锁定资产
	//0普通交易，-1解锁金额交易（退出共识，退出委托）
	Locked byte
}

//资产转入结构体
type CoinTo struct {
	Coin

	//是否将资产锁定
	//0：代表不锁定
	//<0 :代表业务锁定
	//lockValue < 10,000,000时，代表区块高度，当网络高度超过这个数值时，自动解锁（可以使用）
	//10,000,000<lockValue<1,000,000,000,000时，代表当前时间（秒），当时间晚于此值时，自动解锁（可以使用）
	//lockValue>1,000,000,000,000时，代表当前时间（毫秒），当时间晚于此值时，自动解锁（可以使用）
	LockValue uint64
}
