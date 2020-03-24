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
	"testing"
)

func TestAccount(t *testing.T) {

	t.Run("test create account", func(t *testing.T) {
		account, err := NewNULSAccount()
		if err != nil {
			t.Fatalf(err.Error())
		}
		//fmt.Println(account.address+"::::::::"+account.eckey.GetPriKeyHex())
		if account.address == "" {
			t.Fatalf("create account failed")
		}
	})

	t.Run("valid right account", func(t *testing.T) {
		cases := []string{
			"NULSd6HgW1XmWzGsAtPaFbFKrfmV7gXmbLc55",
			"NULSd6HgYs5xFJnfwDyfbjiMBbWsJEswJsj55",
			"NULSd6HghYNHMHLGcXb4MucuUGmAxiomCnE55",
			"NULSd6HgdJ4MNKrSJmtT8XoWMLnUw7g6TFB55",
			"NULSd6HggTFtM7ZMy1D1YBqmbgNeDB8RCit55",
			"NULSd6HgV9pzR7aPyc5C1Lw1adZA38FUWdE55",
			"NULSd6HgeGRiHhK7qnpQkwjHj3zCcXLwxQf55",
			"NULSd6Hgh5jESTW7HJKhEBbNWQU9g164yhb55",
			"NULSd6HgUjZWkqNSiunmdpfoLw4wMdSAsvL55",
			"NULSd6HgVQQQndp5xTmESxjfyqk4JtXBfxC55",
			"NULSd6HgfYbUegUrfqimprUXovuMwU8BikU55",
			"NULSd6HgcEmBj6AJksg46ZWR1gPcoAZtK6C55",
		}
		for _, address := range cases {
			val := Valid(address)
			if !val {
				t.Fatalf("right address but valid failed:" + address)
			}
		}
	})
	t.Run("Test create and valid", func(t *testing.T) {
		//account, _ := NewAccount()
		//
		//success := Valid(account.Address)
		//
		//if !success {
		//	t.Errorf("create account and valid failed.")
		//}
	})
}

func BenchmarkNewNormalAccount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewNormalAccount(1, "NULS")
	}
}
