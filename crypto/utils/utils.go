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

package cryptoutils

import (
	"crypto/sha256"
	aesutils "github.com/niels1286/nuls-go-sdk/crypto/aes"
	"golang.org/x/crypto/ripemd160"
)

//Calculates the SHA-256 hash of the given bytes, * and then hashes the resulting hash again.
func Sha256twice(data []byte) []byte {
	val := sha256.Sum256(Sha256h(data))
	return val[:]
}

//renturn SHA-256 hash of the given data, type : []byte
func Sha256h(data []byte) []byte {
	val := sha256.Sum256(data)
	return val[:]
}

// hash160 returns the RIPEMD160 hash of the SHA-256 HASH of the given data.
func Hash160(data []byte) []byte {
	h := Sha256h(data)
	return Ripemd160h(h[:])
}

//ripemd160h returns the RIPEMD160 hash of the given data.
func Ripemd160h(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	return h.Sum(nil)
}

//Use aes algorithm to encrypt data
func AESEncrpt(data, key []byte) []byte {
	return aesutils.Encrypt(data, key)
}

//Use aes algorithm to decrypt data
func AESDecrpt(data, key []byte) []byte {
	return aesutils.Decrypt(data, key)
}
