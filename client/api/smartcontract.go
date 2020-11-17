// @Title
// @Description
// @Author  Niels  2020/11/17
package api

import (
	"errors"
	"github.com/niels1286/nuls-go-sdk/client/jsonrpc"
	"github.com/niels1286/nuls-go-sdk/tx/txdata"
	"github.com/niels1286/nuls-go-sdk/utils/mathutils"
	"math/rand"
	"time"
)

func GetNRC20Balance(client *jsonrpc.NulsApiClient, chainId uint16, address, contractAddress string) (*TokenBalance, error) {
	if client == nil || address == "" {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "getTokenBalance", []interface{}{chainId, contractAddress, address})
	result, err := client.ApiRequest(param)
	if err != nil {
		return nil, err
	}
	if nil == result || nil == result.Result {
		return nil, errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	balance, err := mathutils.StringToBigInt(resultMap["amount"].(string))
	if err != nil {
		return nil, err
	}
	name := resultMap["name"].(string)
	symbol := resultMap["symbol"].(string)
	decimals := resultMap["decimals"].(float64)
	status := resultMap["status"].(float64)

	return &TokenBalance{
		ContractAddress: contractAddress,
		Name:            name,
		Symbol:          symbol,
		Amount:          balance,
		Decimals:        int64(decimals),
		Status:          int(status),
	}, nil
}

func ImputedContractCallGas(client *jsonrpc.NulsApiClient, chainId uint16, txData *txdata.CallContract) (float64, error) {
	if client == nil || txData == nil {
		return 0, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "imputedContractCallGas", []interface{}{chainId, txData.Sender, txData.Value, txData.ContractAddress, txData.MethodName, txData.MethodDesc, txData.Args})
	result, err := client.ApiRequest(param)
	if err != nil {
		return 0, err
	}
	if nil == result || nil == result.Result {
		return 0, errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	gasLimit := resultMap["gasLimit"].(float64)
	return gasLimit, nil
}

func SCMethodInvokeView(client *jsonrpc.NulsApiClient, chainId uint16, contractAddress, methodName, methodDesc string, args [][]string) (map[string]interface{}, error) {
	if client == nil || contractAddress == "" || methodName == "" || methodDesc == "" {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := jsonrpc.NewRequestParam(rand.Intn(10000), "imputedContractCallGas", []interface{}{chainId, contractAddress, methodName, methodDesc, args})
	result, err := client.ApiRequest(param)
	if err != nil {
		return nil, err
	}
	if nil == result || nil == result.Result {
		return nil, errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	return resultMap, nil
}
