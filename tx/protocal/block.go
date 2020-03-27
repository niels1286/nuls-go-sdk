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

// @Title block.go
// @Description  区块头数据结构
// @Author  Niels  2020/3/26
package txprotocal

import (
	cryptoutils "github.com/niels1286/nuls-go-sdk/crypto/utils"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"log"
)

//区块结构体，包含区块头和交易列表两个部分
type Block struct {
	//区块头
	Header *BlockHeader
	//交易列表
	Txs []*Transaction
}

//根据序列化数据，反序列化区块结构体
func ParseBlock(bytes []byte) *Block {
	block := new(Block)
	block.Parse(seria.NewByteBufReader(bytes, 0))
	return block
}

//将区块完整数据序列化为字节数组
func (b *Block) Serialize() []byte {
	writer := seria.NewByteBufWriter()
	writer.WriteNulsData(b.Header)
	for _, tx := range b.Txs {
		writer.WriteNulsData(tx)
	}
	return writer.Serialize()
}

//从reader中读取所有字段的值，并赋值到结构体中
func (b *Block) Parse(reader *seria.ByteBufReader) {
	b.Header = new(BlockHeader)
	b.Header.Parse(reader)
	for i := uint32(0); i < b.Header.TxCount; i++ {
		b.Txs = append(b.Txs, ParseTransactionByReader(reader))
	}
}

//区块头结构图
type BlockHeader struct {
	hash           *NulsHash
	PreHash        *NulsHash
	MerkleRoot     *NulsHash
	Time           uint32
	Height         uint32
	TxCount        uint32
	Extend         []byte
	PackagerPubkey []byte
	SignData       []byte
}

func (h *BlockHeader) GetHash() *NulsHash {
	if h.hash == nil || len(h.hash.bytes) == 0 {
		bytes, err := h.SerializeForHash()
		if err != nil {
			log.Fatalln(err.Error())
			return nil
		}
		bytes = cryptoutils.Sha256twice(bytes)
		h.hash = NewNulsHash(bytes)
	}
	return h.hash
}

//将区块头完整数据序列化为字节数组
func (h *BlockHeader) Serialize() ([]byte, error) {
	return h.serialize(true)
}
func (h *BlockHeader) SerializeForHash() ([]byte, error) {
	return h.serialize(false)
}
func (h *BlockHeader) serialize(withSign bool) ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteNulsData(h.PreHash)
	writer.WriteNulsData(h.MerkleRoot)
	writer.WriteUInt32(h.Time)
	writer.WriteUInt32(h.Height)
	writer.WriteUInt32(h.TxCount)
	writer.WriteBytesWithLen(h.Extend)
	if withSign {
		writer.WriteBytesWithLen(h.PackagerPubkey)
		writer.WriteBytesWithLen(h.SignData)
	}
	return writer.Serialize(), nil
}

//从reader中读取所有字段的值，并赋值到结构体中
func (h *BlockHeader) Parse(reader *seria.ByteBufReader) error {
	h.PreHash = new(NulsHash)
	h.PreHash.Parse(reader)
	h.MerkleRoot = new(NulsHash)
	h.MerkleRoot.Parse(reader)
	time, err := reader.ReadUint32()
	if err != nil {
		return err
	}
	h.Time = time
	height, err := reader.ReadUint32()
	if err != nil {
		return err
	}
	h.Height = height
	txCount, err := reader.ReadUint32()
	if err != nil {
		return err
	}
	h.TxCount = txCount
	extendBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	h.Extend = extendBytes

	pubBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	h.PackagerPubkey = pubBytes

	signBytes, err := reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	h.SignData = signBytes
	return nil
}
