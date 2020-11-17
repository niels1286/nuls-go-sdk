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

	SimpleCallContractTx(sdk, "", "tNULSeBaN8nYn6u6GCNcLQcozy7mS94swraJzh", "transfer", [][]string{[]string{"tNULSeBaMnrs6JKrCy6TQdzYJZkMZJDng7QAsD"}, []string{"200000000"}})

}

func SimpleCallContractTx(sdk *nuls.NulsSdk, prikeyHex string, contractAddress string, methodName string, args [][]string) {
	chainId := uint16(2)
	assetsId := uint16(1)
	realAcc, _ := account.GetAccountFromPrkey(prikeyHex, chainId, "tNULS")
	accountInfo, _ := sdk.GetBalance(realAcc.Address, chainId, assetsId)

	tx := txprotocal.Transaction{
		TxType:   txprotocal.TX_TYPE_CALL_CONTRACT,
		Time:     uint32(time.Now().Unix()),
		Remark:   nil,
		Extend:   nil,
		CoinData: nil,
		SignData: nil,
	}

	ccData := txdata.CallContract{
		Sender:          realAcc.AddressBytes,
		ContractAddress: account.AddressStrToBytes(contractAddress),
		Value:           big.NewInt(0),
		GasLimit:        100000,
		Price:           25,
		MethodName:      methodName,
		MethodDesc:      "",
		ArgsCount:       uint8(len(args)),
		Args:            args,
	}
	tx.Extend, _ = ccData.Serialize()
	coinData := txprotocal.CoinData{
		Froms: []txprotocal.CoinFrom{{
			Coin: txprotocal.Coin{
				Address:       realAcc.AddressBytes,
				AssetsChainId: chainId,
				AssetsId:      assetsId,
				Amount:        big.NewInt(2600000),
			},
			Nonce:  accountInfo.Nonce,
			Locked: 0,
		}},
		Tos: nil,
	}
	tx.CoinData, _ = coinData.Serialize()
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
