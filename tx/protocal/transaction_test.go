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

// @Title transaction_test.go
// @Description  交易结构的测试工具
// @Author  Niels  2020/3/25
package txprotocal

import (
	"encoding/hex"
	"testing"
)

func TestTransaction(t *testing.T) {
	txHex := "020056247b5e00008c0117010001b5d342e080f5ec95f901eb8ea11db03937410ff201000100a067f70500000000000000000000000000000000000000000000000000000000080fdef8a93176062b000117010001cd148bd3fe8bc4d62542e37ab4b6b84407bd7e450100010000e1f5050000000000000000000000000000000000000000000000000000000000000000000000006a21023b2c460aad25cbdcdc7091a47d101f09fde64003c3143053f2549494a8ca3c20473045022100f517849cae4622cb4a30b31a8b044f74962e73a42192aae4ba22df40b266bd8e022035df019356fccd79c24a7d98a24d5bd083931afe29f8bc7247ee430cbc825d30"
	txBytes, _ := hex.DecodeString(txHex)
	tx := ParseTransaction(txBytes, 0)
}
