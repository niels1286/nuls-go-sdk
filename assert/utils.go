package assert

import "testing"

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
	t.Fatalf("stop todo ...")
}
