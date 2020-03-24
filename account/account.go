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
package account

import (
	"crypto/sha256"
	"github.com/niels1286/nerve-go-sdk/crypto/eckey"
	"golang.org/x/crypto/ripemd160"
)

type Account struct {
	address string
	chainId uint16
	accType uint8
	eckey   eckey.EcKey
}

//创建一个新账户
func NewNormalAccount(chainId uint16, prefix string) (Account, error) {
	ec, err := eckey.NewEcKey()
	if err != nil {
		return Account{}, err
	}
	pubBytes := ec.GetPubKeyBytes(true)
	digest := sha256.Sum256(pubBytes)
	digest = ripemd160.New().Sum([]byte{digest})
}

func NewNULSAccount() (Account, error) {

}

func NewNerveAccount() (Account, error) {

}

func ParseAccount(address string) Account {

}

func Valid(address string) bool {

}
