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
// @Author  Niels  2020/3/27
package txprotocal

import "github.com/niels1286/nuls-go-sdk/utils/seria"

type SignData struct {
	Signatures []P2PHKSignature
}

type P2PHKSignature struct {
	SignValue []byte
	PublicKey []byte
}

//序列化方法
func (s *SignData) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	for _, sig := range s.Signatures {
		writer.WriteByte(byte(len(sig.PublicKey)))
		writer.WriteBytes(sig.PublicKey)
		writer.WriteBytesWithLen(sig.SignValue)
	}
	return writer.Serialize(), nil
}

//反序列化方法
func (s *SignData) Parse(reader *seria.ByteBufReader) error {
	for !reader.IsFinished() {
		sig := P2PHKSignature{}
		length, err := reader.ReadByte()
		if err != nil {
			return err
		}
		sig.PublicKey, err = reader.ReadBytes(int(length))
		if err != nil {
			return err
		}
		sig.SignValue, err = reader.ReadBytesWithLen()
		if err != nil {
			return err
		}
		s.Signatures = append(s.Signatures, sig)
	}
	return nil
}
