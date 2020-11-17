// @Title
// @Description
// @Author  Niels  2020/11/17
package txdata

import (
	"encoding/hex"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"testing"
)

func TestCallContrace_Parse(t *testing.T) {
	txHex := "10001535a85d006a0200013452fcc77369361d225f61ea34b9930f98d2b536020002a6d26cb6b330c7a0aaf47848e5c27e3abe1a980400000000000000000000000000000000000000000000000000000000000000007a500000000000001900000000000000087365744167656e740001004801170200013452fcc77369361d225f61ea34b9930f98d2b536020001008a62090000000000000000000000000000000000000000000000000000000000087e8ac0a9028cd89b00006a2103e56a675cd355d11bc7667d53adfc7c70d50abb4df7cf438df55253db5e10e22d473045022100b293a59aa4db375e6ccee44033bb115f535ac1f098e992ace4e65cb87581dd0402203ac0896e5e0fba202279d4aacca2bd99af51789884b6d15e9195a3aa0625f317"
	txBytes, _ := hex.DecodeString(txHex)
	tx := txprotocal.ParseTransaction(txBytes, 0)

	data := CallContract{}
	data.Parse(seria.NewByteBufReader(tx.Extend, 0))

	val1 := hex.EncodeToString(tx.Extend)
	bytes, _ := data.Serialize()
	val2 := hex.EncodeToString(bytes)
	if val1 != val2 {
		t.Error("错误")
	}

}
