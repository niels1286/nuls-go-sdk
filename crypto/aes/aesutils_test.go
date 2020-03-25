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

package aesutils

import (
	"encoding/hex"
	cryptoutils "github.com/niels1286/nuls-go-sdk/crypto/utils"
	"github.com/niels1286/nuls-go-sdk/utils/assert"
	"testing"
)

func TestAes(t *testing.T) {
	dataHex := "230cb8ebbf3a2c581d27f98f7a38f8b07c1ff170d605ca645db4ffa05ffa5505"
	key := []byte("qwer1234")
	key = cryptoutils.Sha256h(key)
	result := "9308015df78d2ee856710c459dc8166589172fdda356ce96277ba5ccbcbeca369cb8085c5326d500453a643d55e5269e"
	t.Run("test encrypt", func(t *testing.T) {
		data, _ := hex.DecodeString(dataHex)
		got := Encrypt(data, key)
		gotHex := hex.EncodeToString(got)
		assert.IsEquals(t, result, gotHex)
	})
	t.Run("test decrypt", func(t *testing.T) {
		ciphertext, _ := hex.DecodeString(result)
		got := Decrypt(ciphertext, key)
		gotHex := hex.EncodeToString(got)
		assert.IsEquals(t, dataHex, gotHex)
	})
}
