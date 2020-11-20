// @Title
// @Description
// @Author  Niels  2020/11/17
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/niels1286/nuls-go-sdk"
	"github.com/niels1286/nuls-go-sdk/account"
	txprotocal "github.com/niels1286/nuls-go-sdk/tx/protocal"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"github.com/niels1286/nuls-go-sdk/utils/seria"
	"math/big"
	"time"
)

type ResultVo struct {
	Sender          string
	ContractAddress string
	Value           string
	GasLimit        uint64
	Price           uint64
	MethodName      string
	MethodDesc      string
	ArgsCount       uint8
	Args            [][]string
}

func main() {
	//sdk := nuls.NewNulsSdk("http://beta.api.nuls.io/jsonrpc/", "http://beta.wallet.nuls.io/api/", 2)
	//SimpleCallContractTx(sdk, "", "tNULSeBaN8nYn6u6GCNcLQcozy7mS94swraJzh", "transfer", [][]string{[]string{"tNULSeBaMnrs6JKrCy6TQdzYJZkMZJDng7QAsD"}, []string{"200000000"}})
	//sdk := nuls.NewNulsSdk("https://api.nuls.io/jsonrpc/", "https://wallet.nuls.io/api/", 1)
	//
	//result ,err := sdk.SCMethodInvokeView(1,"NULSd6HgvBGqSQBr49QmB9BJia4RnzsAWpjtE","tokenList","",[][]string{[]string{""}})
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(result)
	txDataHex := "0100011d52afe277b0c575355e618a194ffa1ae1cb518f010002be36277487d7c45974d1fcc722d6729f47cddd8900000000000000000000000000000000000000000000000000000000000000001b540000000000001900000000000000087472616e73666572000201254e554c5364364867554763523351626f76593155467561753532374532464e354c70434e54010a33303030303030303030"

	txDataBytes, err := hex.DecodeString(txDataHex)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := &txdata.CallContract{}
	data.Parse(seria.NewByteBufReader(txDataBytes, 0))
	result := &ResultVo{
		Sender:          account.GetStringAddress(data.Sender, "NULS"),
		ContractAddress: account.GetStringAddress(data.ContractAddress, "NULS"),
		Value:           data.Value.String(),
		GasLimit:        data.GasLimit,
		Price:           data.Price,
		MethodName:      data.MethodName,
		MethodDesc:      data.MethodDesc,
		ArgsCount:       data.ArgsCount,
		Args:            data.Args,
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(bytes))
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
