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

package eckey

import (
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcd/btcec"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"log"
	"math/big"
)

//用于管理公私钥的结构
//依赖于ethereum的基础包，但本结构总的数据都是golang基础包中的
type EcKey struct {
	privateKey *btcec.PrivateKey
	publicKey  *btcec.PublicKey
}

//随机生成新的公私钥对
func NewEcKey() (*EcKey, error) {
	prikey, err := btcec.NewPrivateKey(btcec.S256())
	if nil != err {
		return nil, err
	}
	pubkey := prikey.PubKey()
	return &EcKey{prikey, pubkey}, nil
}

//根据私钥字符串，还原EcKey
func FromPriKeyBytes(priBytes []byte) (*EcKey, error) {

	val := mathutils.BytesToBigInt(priBytes)
	if val == nil || val.Cmp(big.NewInt(1)) <= 0 {
		return nil, errors.New("Private key is wrong!")
	}

	prikey, pubkey := btcec.PrivKeyFromBytes(btcec.S256(), priBytes)
	return &EcKey{prikey, pubkey}, nil
}

//获取hex编码后的私钥
func (e *EcKey) GetPriKeyHex() string {
	return hex.EncodeToString(e.GetPriKeyBytes())
}

//获取私钥字节数组
func (e *EcKey) GetPriKeyBytes() []byte {
	return e.privateKey.Serialize()
}

//获取hex编码后的公钥
//compressed代表是否压缩，true为压缩，false为不压缩
func (e *EcKey) GetPubKeyHex(compressed bool) string {
	return hex.EncodeToString(e.GetPubKeyBytes(compressed))
}

//获取公钥的字节数组
//compressed代表是否压缩，true为压缩，false为不压缩
func (e *EcKey) GetPubKeyBytes(compressed bool) []byte {
	if compressed {
		return e.publicKey.SerializeCompressed()
	}
	return e.publicKey.SerializeUncompressed()
}

//根据公钥还原eckey，可以进行签名验证及数据解密
func FromPubKeyBytes(pubBytes []byte) (*EcKey, error) {
	val := mathutils.BytesToBigInt(pubBytes)
	if val == nil || val.Cmp(big.NewInt(1)) <= 0 {
		return nil, errors.New("Private key is wrong!")
	}

	pubkey, err := btcec.ParsePubKey(pubBytes, btcec.S256())
	if err != nil {
		return nil, err
	}
	return &EcKey{nil, pubkey}, nil
}

//判断结构中的私钥是否为空
func (e *EcKey) IsPriKeyNil() bool {
	return nil == e.privateKey
}

func (e *EcKey) Sign(data []byte) ([]byte, error) {
	sig, err := e.privateKey.Sign(data)
	if err != nil {
		return nil, err
	}
	return sig.Serialize(), nil
}

func (e *EcKey) Verify(data, signature []byte) bool {
	sig, err := btcec.ParseSignature(signature, btcec.S256())
	if err != nil {
		log.Print(err.Error())
		return false
	}
	return sig.Verify(data, e.publicKey)
}

//使用公钥加密数据
func (e *EcKey) Encrypt(data []byte) []byte {
	val, err := btcec.Encrypt(e.publicKey, data)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	return val
}

//使用私钥解密以公钥加密的数据
func (e *EcKey) Decrypt(in []byte) []byte {
	val, err := btcec.Decrypt(e.privateKey, in)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	return val
}
