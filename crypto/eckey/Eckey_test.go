package eckey

import (
	"fmt"
	"nerve-go-sdk/assert"
	"testing"
)

type eckeyCase struct {
	prikey string
	pubkey string
}

func TestEckey(t *testing.T) {
	t.Run("test random create.", func(t *testing.T) {
		eckey := NewEckey()
		fmt.Println(eckey)
		assert.NotNil(t, eckey.PubKey, "public key ")
		assert.NotNil(t, eckey.PriKey, "private key ")

		want := eckey.PubKey

		_eckey := FromPriKey(eckey.PriKey)

		assert.IsEquals(t, want, _eckey.PubKey)
	})
	cases := []eckeyCase{eckeyCase{"eebbb62294ab36f1336626074f1e81482ee17700f10e11a9875c76bd31d0a07f", "03182be87cd707c5965caab6ec808d8587260e1dae8d21a880f1ef584938725a34"},
		eckeyCase{"1ed08cc206d0a91f1cdd8913337bcd1ecd05308a4d18727aea6346ade78e863e", "03c29580ae7eba733e238ce9d6600ada110f4b83e9ded355d3796984a38a8a9690"},
		eckeyCase{"680be777d376e295a5e4cf81952f81de475f5c1a17dc22f79152460d5cc6b3c0", "02fe662c92d094d6d287cf42240a4ff6684775610f00e376041b1e5fda34f383b4"},
		eckeyCase{"5dd7e20845b5e1ff9805223bee5316f3457cbfb6538188dc9626ba20f58aa9f5", "029d4e5860752089f8bd92aa98058e8fa5d8633d58653417f9977a6e23d7271f7d"},
		eckeyCase{"1295c6f76978540d72cf1468a26311d21866d73c12e0375a20cdb81a51617316", "026e158e80c350dd27f01f93030fce0294fe8d1f97e6345d536270289c7bf8affb"},
		eckeyCase{"6d2d3e4ccea9767f2a04777788e3a118d890e9051c5529b141fcba3c63ae20cc", "02f3a8f536c47caa7c50dc98454fcd8405d40606c15a38e9bc633ab061beac5523"},
		eckeyCase{"93ba05ec77258b58139b0f1f56099760a6f2d559989d0f840d7e98b0a79f979e", "02c5df80aa8bda6e8e092d00d5ea194637946a751ece6c8fe70236404dd258e23d"},
		eckeyCase{"c8c6e57f78a2df72053f9773bf14cc17b0c26085d5c28c9fc24cb01bbd25e7d9", "021b45f07b4622afa489173249e4d1610dc65fcb2456394617102eddd05a7a0c1c"},
		eckeyCase{"9e26827d4b01c9ef1229286aee113c0dcc180e898ca8be488f2989cea445d600", "02e6035ee388cda1c28bf54ba1926d8ddeefa603382dc6cf20b31c3cd88cf4fe36"},
		eckeyCase{"4c04cfbb6dbfce896098fd51fc0983f109ca351efc280bd163c12635b9535a93", "02f322e62d2cf1b6014c3f25652f8cad057667b446bebdfa90f0039502a9356151"},
		eckeyCase{"ccf6826d330fb3bf57929f9cfce699a92d07248dec9073c8c5419fd7ab2b7e24", "0202977f0b00a3d25d27c1757d40ef6343c06323f8f74c2284b3ad9941ab01789b"},
		eckeyCase{"7af069e5b315b331e536437e7e3ca442b10bd19b76b1163d0c9cbe2820d012ae", "025428014cf428fd7160439d31cbae50799f42022ae12d0ccbdc0267f85e0cb46f"}}

	t.Run("test compatibility", func(t *testing.T) {
		for _, item := range cases {
			_eckey := FromPriKey(item.prikey)
			assert.IsEquals(t, _eckey.CompressedPubKey, item.pubkey)
		}
	})
}
