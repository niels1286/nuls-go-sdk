# nuls-go-sdk 

	    sdk := utils.GetOfficalSdk()

			hash, err := sdk.BroadcastTx(resultBytes)
			if err != nil {
				fmt.Println("Failed!\nNewTxHex:", resultHex)
				fmt.Println(err.Error())
				return
			}
