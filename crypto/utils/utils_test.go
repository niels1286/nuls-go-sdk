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
	"encoding/hex"
	"testing"
)

func TestUtils(t *testing.T) {
	//data : 2d3aef2c4abba64e23bdff0c626483e7d7062255443cd627989cd35bc2a94dcf
	//hash256:: 51f06ec8c4bb2d143e47dbcedc7e99a888936742508eec6a5d6b4ee57f4997d2
	//	hash256 twice:: 6ef53f7b750355fb89520fb995ebbdd230126dbfbe6a47d25cdb3e6d10749c62
	//hash160:: d3e5ff3a3493c0127cd038ad517f7c3ff51a2fc9
	//ripemd160h:: 7c0d7bb7a616e9072ceb66dddbf87f071ac65690
	bytes, err := hex.DecodeString("2d3aef2c4abba64e23bdff0c626483e7d7062255443cd627989cd35bc2a94dcf")
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Run("test sha256", func(t *testing.T) {
		val := Sha256h(bytes)
		if hex.EncodeToString(val) != "51f06ec8c4bb2d143e47dbcedc7e99a888936742508eec6a5d6b4ee57f4997d2" {
			t.Fatalf("sha256 failed")
		}
	})
	t.Run("sha 256 twice", func(t *testing.T) {
		val := Sha256twice(bytes)
		if hex.EncodeToString(val) != "6ef53f7b750355fb89520fb995ebbdd230126dbfbe6a47d25cdb3e6d10749c62" {
			t.Fatalf("sha256 twice failed")
		}
	})
	t.Run("test ripemd160", func(t *testing.T) {
		val := Ripemd160h(bytes)
		if hex.EncodeToString(val) != "7c0d7bb7a616e9072ceb66dddbf87f071ac65690" {
			t.Fatalf("ripemd160 failed")
		}
	})
	t.Run("test hash160", func(t *testing.T) {
		val := Hash160(bytes)
		if hex.EncodeToString(val) != "d3e5ff3a3493c0127cd038ad517f7c3ff51a2fc9" {
			t.Fatalf("hash160 failed")
		}
	})

}
