package Faucet

import (
	"github.com/proximax-storage/faucet-backend/utils"
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk"
)

var (
	BlockchainClient *sdk.Client
)

func InitClient() {
	utils.Logger(0, "Initializing rest clients")

	conf, err := sdk.NewConfig(Config.Blockchain.ApiUrl, Config.NetworkType())
	if err != nil {
		panic(err)
	}

	BlockchainClient = sdk.NewClient(nil, conf)

	utils.Logger(0, "Initializing rest clients - completed")
}
