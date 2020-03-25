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
	"github.com/niels1286/nuls-go-sdk/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestKeystore(t *testing.T) {
	t.Run("test create keystore ", func(t *testing.T) {

	})
	t.Run("test import keystore ", func(t *testing.T) {
		file, removeFile := createTempFile(t, `{
		"address":"tNULSeBaMtDQocJJiBGEDwFpPQPKXJo6pXmW7b",
		"encryptedPrivateKey":"9308015df78d2ee856710c459dc8166589172fdda356ce96277ba5ccbcbeca369cb8085c5326d500453a643d55e5269e",
		"pubKey":"0233dd5281a4e129dafeea8637b54806f667e56f654e098c5faab87fa7fe889d11"}`)
		defer removeFile()
		keystore, err := KeystoreFromFile(file)
		assert.IsNil(t, err, "keystore err")
		assert.IsEquals(t, keystore.Address, "tNULSeBaMtDQocJJiBGEDwFpPQPKXJo6pXmW7b")
		assert.IsEquals(t, keystore.EncryptedPrivateKey, "9308015df78d2ee856710c459dc8166589172fdda356ce96277ba5ccbcbeca369cb8085c5326d500453a643d55e5269e")
		assert.IsEquals(t, keystore.Pubkey, "0233dd5281a4e129dafeea8637b54806f667e56f654e098c5faab87fa7fe889d11")
	})
}

//创建一个临时文件，用以模拟keystore导入功能
func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
