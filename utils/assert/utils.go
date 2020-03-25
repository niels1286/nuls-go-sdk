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

package assert

import (
	"reflect"
	"testing"
)

//传入的数据必须是nil，否则失败
func IsNil(t *testing.T, val interface{}, title string) {
	t.Helper()
	if nil != val {
		t.Errorf("%s must be nil,but it's %s", title, val)
	}
}

//传入的数据必须不等于nil，否则失败
func NotNil(t *testing.T, val interface{}, title string) {
	t.Helper()
	if nil == val {
		t.Errorf("%s must not be nil,but it's nil", title)
	}
}

func IsEquals(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Fatalf("got %s bug want %s", got, want)
	}
}

func IsDeepEquals(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got and want not equals.")
	}
}
