package Faucet

import (
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
)

var (
	BlockchainClient *sdk.Client
)

func InitClient() {
	utils.Logger(0, "Initializing rest clients")

	conf, err := sdk.NewConfig([]string{Config.Blockchain.ApiUrl}, Config.NetworkType(), 0)
	if err != nil {
		panic(err)
	}

	BlockchainClient = sdk.NewClient(nil, conf)

	utils.Logger(0, "Initializing rest clients - completed")
}
