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
// @Author  Niels  2020/4/8
package txdata

import (
	"github.com/niels1286/nuls-go-sdk/utils/seria"
)

//设置别名交易的拓展字段的详细协议
type Alias struct {
	//地址
	Address []byte
	//别名
	Alias string
}

//反序列化
func (a *Alias) Parse(reader *seria.ByteBufReader) error {
	var err error
	a.Address, err = reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	a.Alias, err = reader.ReadStringWithLen()
	if err != nil {
		return err
	}
	return nil
}

//序列化方法
func (a *Alias) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBytesWithLen(a.Address)
	writer.WriteString(a.Alias)
	return writer.Serialize(), nil
}
