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
	"fmt"
	"github.com/niels1286/nuls-go-sdk/crypto/eckey"
	"log"
	"testing"
)

func TestAccount(t *testing.T) {

	t.Run("test create account", func(t *testing.T) {
		account, err := NewNormalAccount(NULSChainId, NULSPrefix)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if account.Address == "" {
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
		account, _ := NewNormalAccount(NULSChainId, NULSPrefix)

		success := Valid(account.Address)

		if !success {
			t.Errorf("create account and valid failed.")
		}
	})
	t.Run("test parse address", func(t *testing.T) {
		cases := []string{
			"NULSd6HgUjZWkqNSiunmdpfoLw4wMdSAsvL55",
			"tNULSeBaMvEtDfvZuukDf2mVyfGo3DdiN8KLRG",
		}
		values := []Account{{Address: "NULSd6HgUjZWkqNSiunmdpfoLw4wMdSAsvL55", ChainId: NULSChainId, AccType: NormalAccountType, EcKey: eckey.EcKey{}, Prefix: NULSPrefix},
			{Address: "tNULSeBaMvEtDfvZuukDf2mVyfGo3DdiN8KLRG", ChainId: TNULSChainId, AccType: NormalAccountType, EcKey: eckey.EcKey{}, Prefix: TNULSPrefix}}
		for index, address := range cases {
			account, err := ParseAccount(address)
			if err != nil {
				t.Fatalf("Parse account failed 0.")
			}
			if account.ChainId != values[index].ChainId || account.AccType != values[index].AccType || account.Prefix != values[index].Prefix || account.Address != values[index].Address {
				t.Fatalf("parse account failed")
			}
			if GetStringAddress(account.AddressBytes, account.Prefix) != account.Address {
				t.Fatalf("parse account failed 2")
			}
		}
	})
	t.Run("test import account", func(t *testing.T) {
		prikey := "16527f8847017225e39465cf98036fcf1c9b708bae24a38d139913d87e95e805"
		account, err := GetAccountFromPrkey(prikey, NULSChainId, NULSPrefix)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if account.Address != "NULSd6Hgh2fZhgPTbQ1UTpSoXBxNotChyoYgE" {
			t.Fatalf("import account failed.")
		}

	})
}

func TestGetRealAddress(t *testing.T) {
	cases := []string{
		"NULSd6HgW1XmWzGsAtPaFbFKrfmV7gXmbLc55",
		"NULSd6HgYs5xFJnfwDyfbjiMBbWsJEswJsj55",
		"NULSd6HghYNHMHLGcXb4MucuUGmAxiomCnE55",
		"SHITd6HgdJ4MNKrSJmtT8XoWMLnUw7g6TFB55",
		"NULSd6HggTFtM7ZMy1D1YBqmbgNeDB8RCit55",
		"tNULSeBaMvEtDfvZuukDf2mVyfGo3DdiN8KLRG",
		"tNULSeBaMnrs6JKrCy6TQdzYJZkMZJDng7QAsD",
	}

	values := []string{
		"6HgW1XmWzGsAtPaFbFKrfmV7gXmbLc55",
		"6HgYs5xFJnfwDyfbjiMBbWsJEswJsj55",
		"6HghYNHMHLGcXb4MucuUGmAxiomCnE55",
		"6HgdJ4MNKrSJmtT8XoWMLnUw7g6TFB55",
		"6HggTFtM7ZMy1D1YBqmbgNeDB8RCit55",
		"BaMvEtDfvZuukDf2mVyfGo3DdiN8KLRG",
		"BaMnrs6JKrCy6TQdzYJZkMZJDng7QAsD",
	}

	for index, address := range cases {
		_, got := getRealAddress(address)
		if got != values[index] {
			t.Fatalf("Get real address string failed.")
		}
	}

}

func BenchmarkNewNormalAccount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewNormalAccount(1, NULSPrefix)
	}
}

//批量创建NULS主网账户的示例
func ExampleNewNULSAccount() {
	//一次生成10个账户
	accounts, err := NewNULSAccount(10)
	if err != nil {
		log.Fatal("Create NULS mainnet account failed.")
		return
	}
	for _, account := range accounts {
		//打印账户地址和对应的私钥
		fmt.Println(account.Address + "==" + account.GetPriKeyHex())
	}
	//NULSd6Hgh2fZhgPTbQ1UTpSoXBxNotChyoYgE==16527f8847017225e39465cf98036fcf1c9b708bae24a38d139913d87e95e805
	//NULSd6HghwPEJHBhTqToddfJhKwLZfKpsw1ZL==fb5aec7082b676f855f7f352e2e3dcbc3d6a3d58f35e22a09d5e9c86666bc53a
	//NULSd6HgagjzmJzB74ueonkx29tYtJDDmUwGK==df379b4910d17940078c7a5da05312001a3b519171f424dfdb7473739a3e1a8e
	//NULSd6HgfkdyohcsN8jSwMyDz3JecnTpYSxdx==58a90cdfc5e6befde7ae10d440cf7442d2ad16e741b793776869e7828ba262be
	//NULSd6HgWEcod6fdK3Tx8ybQy3Mb2xGSUcRPa==f1abf7bd12e7f1631db7ccc6eed70b03443d7efa707c8f6c4fbf69d91cb7f97a
	//NULSd6Hgi2rL9ukUDcdAMKjtyK2jo49pGo6wo==8e987fa3bd7c34cfa54fe67993df13a6bbdde00e07f1fdd9dc9f1a1b0766f14d
	//NULSd6Hgbyk7nLhMtZ2YuhjWsaWhTaPWmbrFz==03109009826b68bfc9cdc14ee3e004f763a1accdab00d3321bf781005d1370de
	//NULSd6HgZi8p8pSjhKgwnfqRd2YDGfpFVjr2W==0edcc117020bf9a492aabdece99fa77884a10997322edfb7d54344e952913cfb
	//NULSd6HgUbjmcdbwVETFPM72EPQLa93ALdmD7==5e63094848e19360b00d27b4ed79501abd885553c9383fac36f7a402d4e675d0
	//NULSd6HgWEaMdkA7LsHXqbdA6os4EnyEZdJg4==4ffbe58760fdc4db10fe9bc930cb4ec27cabd7d1bb3ea8a2db7d34572e03e61c
}

//批量创建NULS测试网账户的示例
func ExampleNewTNULSAccount() {
	//一次生成10个账户
	accounts, err := NewTNULSAccount(10)
	if err != nil {
		log.Fatal("Create NULS mainnet account failed.")
		return
	}
	for _, account := range accounts {
		//打印账户地址和对应的私钥
		fmt.Println(account.Address + "==" + account.GetPriKeyHex())
	}
	//tNULSeBaMtnMutUtYdG7aY8ez1repGkSP6aGxU==f238c843ec712e9b855284522607ee100dfce6172e5cc6440a55504d373e5020
	//tNULSeBaMvT85ZfUStMUR3czgyqH3xX7sTYDmh==52cddc098953a9219de357c44effee2573e2f47f85d883a758b2908a1ae9ffb0
	//tNULSeBaMuKGxsYTDbC6owHJcj5pLSeUJ5hCEs==6ad702593cf281da212f8e2f0f233c4b2b368de9e82f63bde89383cf65a7cbf3
	//tNULSeBaMfsTDMm6fvzVgo7BhZwTr2KcddY2WN==2e180689e939d2f74aa30c1cc0dcec3d322b66bc7ea6031e2a43f1cfe08f08e4
	//tNULSeBaMo2noa2a1yPATgB1j48D7iRW1ip1Wm==1ffdf039d93773df333c0102995f45096db20b20ee722ee907615f2147656232
	//tNULSeBaMvfXtbzQ6BDzJZxJWaxFqDsafEjDHE==cdb2ea0dc44d0b75f21d5efd198e1f418aaf27e21930c8e3b5e0de5b0cd91ada
	//tNULSeBaMnBPNkwNzHkoLujxeyUCmZp7cDbPie==d8dd6dd2918090f9839766d08d67c16130857c866be642caef765e084904a4b6
	//tNULSeBaMfPtJojgo4326pkGDoL3urfDeUr7LY==a34fddfc29fd8f9ed926a7bfd3fc6d95f3ada562428037d2d025757656f2679d
	//tNULSeBaMhtBv61fnznegqTmMvYgcs9Yqnzk7g==9db69ef368b3e3176918de21ba5f1927e9beb61d0b02cd74fc796279aa85233c
	//tNULSeBaMgzeDmZQetKWfVb5AjsTg8SuzfRLiS==7127ff18134ee74f248850ae496f3617b1ef69bdc59e818539e5aaeea43c7969
}

//根据私钥创建账户的示例
func ExampleGetAccountFromPrkey() {
	//私钥
	prikey := "16527f8847017225e39465cf98036fcf1c9b708bae24a38d139913d87e95e805"
	//调用方法，生成NULS主网地址
	account, err := GetAccountFromPrkey(prikey, NULSChainId, NULSPrefix)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	//跟预先生成的地址做对比
	if account.Address != "NULSd6Hgh2fZhgPTbQ1UTpSoXBxNotChyoYgE" {
		log.Fatalf("import account failed.")
	}
	fmt.Println(account)
}

//地址验证是非常常用的功能，任何业务逻辑都应该先验证涉及的地址是否正确
func ExampleValid() {
	address := "tNULSeBaMgzeDmZQetKWfVb5AjsTg8SuzfRLiS"
	result := Valid(address)

	if !result {
		log.Println("%s is not a right address", address)
	}
}
