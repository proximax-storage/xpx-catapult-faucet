package Faucet

import (
	"context"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
	ws "github.com/proximax-storage/xpx-catapult-faucet/websocket"

	"time"
)

var BlockchainClient = new(sdk.Client)

func InitClient() error {
	utils.Logger(0, "initializing sirius client")
	conf, err := sdk.NewConfig(context.Background(), []string{Config.Blockchain.ApiUrl})
	conf.FeeCalculationStrategy = Config.Blockchain.FeeCalculationStrategy
	if err != nil {
		return err
	}

	// Use the default http client
	BlockchainClient = sdk.NewClient(nil, conf)
	utils.Logger(0, "initialized sirius client - completed")

	return nil
}

func NewWebsocket() (*ws.ClientWebsocket, error) {
	return ws.NewConnectWs(Config.Blockchain.ApiUrl, time.Second*60)
}
