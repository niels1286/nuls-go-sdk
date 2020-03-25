package aesutils

import (
	"encoding/hex"
	"github.com/niels1286/nuls-go-sdk/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	dataHex := "230cb8ebbf3a2c581d27f98f7a38f8b07c1ff170d605ca645db4ffa05ffa5505"
	want := "9308015df78d2ee856710c459dc8166589172fdda356ce96277ba5ccbcbeca369cb8085c5326d500453a643d55e5269e"
	data, _ := hex.DecodeString(dataHex)
	secret := []byte("qwer1234")
	got := Encrypt(data, secret)
	gotHex := hex.EncodeToString(got)
	assert.IsEquals(t, want, gotHex)
}
