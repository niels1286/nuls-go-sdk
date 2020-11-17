// @Title
// @Description
// @Author  Niels  2020/11/17
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"math/big"
	"time"
)

func main() {
	sdk := nuls.NewNulsSdk("http://beta.api.nuls.io/jsonrpc/", "http://beta.wallet.nuls.io/api/", 2)
	acc, _ := account.ParseAccount("tNULSeBaMoG1oaW1JZnh6Ly65Ttp6raeTFBfCG")
	//balanceInfo, _ := sdk.GetNRC20Balance(acc.Address, "tNULSeBaN8nYn6u6GCNcLQcozy7mS94swraJzh")
	accountInfo, _ := sdk.GetBalance(acc.Address, 2, 1)

	tx := txprotocal.Transaction{
		TxType:   txprotocal.TX_TYPE_CALL_CONTRACT,
		Time:     uint32(time.Now().Unix()),
		Remark:   []byte("golang test"),
		Extend:   nil,
		CoinData: nil,
		SignData: nil,
	}

	ccData := txdata.CallContract{
		Sender:          acc.AddressBytes,
		ContractAddress: account.AddressStrToBytes("tNULSeBaN8nYn6u6GCNcLQcozy7mS94swraJzh"),
		Value:           big.NewInt(0),
		GasLimit:        50000,
		Price:           25,
		MethodName:      "transfer",
		MethodDesc:      "",
		ArgsCount:       2,
		Args:            [][]string{[]string{"tNULSeBaMnrs6JKrCy6TQdzYJZkMZJDng7QAsD"}, []string{"100000000"}},
	}
	tx.Extend, _ = ccData.Serialize()
	coinData := txprotocal.CoinData{
		Froms: []txprotocal.CoinFrom{{
			Coin: txprotocal.Coin{
				Address:       acc.AddressBytes,
				AssetsChainId: 2,
				AssetsId:      1,
				Amount:        big.NewInt(1350000),
			},
			Nonce:  accountInfo.Nonce,
			Locked: 0,
		}},
		Tos: nil,
	}
	tx.CoinData, _ = coinData.Serialize()

	realAcc, _ := account.GetAccountFromPrkey("", 2, "tNULS")

	hash, _ := tx.GetHash().Serialize()
	signValue, _ := realAcc.Sign(hash)

	txSign := txprotocal.CommonSignData{
		Signatures: []txprotocal.P2PHKSignature{{
			SignValue: signValue,
			PublicKey: realAcc.GetPubKeyBytes(true),
		}},
	}
	tx.SignData, _ = txSign.Serialize()
	resultBytes, _ := tx.Serialize()
	bcdResult, err := sdk.BroadcastTx(resultBytes)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(bcdResult)
	}

	txHex := hex.EncodeToString(resultBytes)
	fmt.Println(txHex)
}
