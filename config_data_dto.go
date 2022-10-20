package Faucet

import (
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"strings"
	// "sync"
	"time"
)

// var (
// 	mtx sync.RWMutex
// )

type ConfigData struct {
	Blockchain Blockchain `json:"blockchain"`
	Server     Server     `json:"server"`
	Logging    Logging    `json:"logging"`
	BlackList  BlackList  `json:"blackList"`
	WhiteList  WhiteList  `json:"whiteList"`
	App        App        `json:"app"`
	DbStorage  DbStorage  `json:"db"`
}

type DbStorage struct {
	Dir          string        `json:"dir"`
	DiscardRatio float64       `json:"discardRatio"`
	GcInterval   time.Duration `json:"gcInterval"`
}

type BlackList struct {
	ByIp      bool `json:"byIp"`
	ByAddress bool `json:"byAddress"`
}

type WhiteList struct {
	Addresses []string `json:"addresses"`
}

type Blockchain struct {
	ApiUrl  string `json:"apiUrl"`
	Network string `json:"network"`
}

type Server struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Port             int    `json:"port"`
	Host             string `json:"host"`
	AllowCrossDomain bool   `json:"allowCrossDomain"`
}

type ServiceInfo struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type Logging struct {
	Console struct {
		Colorize         bool   `json:"colorize"`
		Level            string `json:"level"`
		HandleExceptions bool   `json:"handleExceptions"`
		PrettyPrint      bool   `json:"prettyPrint"`
		Timestamp        bool   `json:"timestamp"`
	} `json:"console"`
	ErrCtrl struct {
		MaxNumErr uint16 `json:"maxNumErr"`
	} `json:"errCtrl"`
}

type App struct {
	FaucetMasterAcctPrivateKey string     `json:"FaucetMasterAcctPrivateKey"`
	Namespace                  string     `json:"namespace"`
	MaxXpx                     sdk.Amount `json:"maxXpx"`
	MosaicId                   string     `json:"mosaicId"`
}

var FaucetAccount *sdk.Account

func (c *ConfigData) FaucetAccount() *sdk.Account {
	if FaucetAccount != nil {
		return FaucetAccount
	} else {
		FaucetAccount, err := BlockchainClient.NewAccountFromPrivateKey(strings.ToUpper(c.App.FaucetMasterAcctPrivateKey))
		if err != nil {
			panic(err)
		}
		return FaucetAccount
	}
}

func (c *ConfigData) FaucetAccountPublicKey() string {
	return strings.ToUpper(c.FaucetAccount().PublicAccount.PublicKey)
}

func (c *ConfigData) FaucetPublicAccount() *sdk.PublicAccount {
	return c.FaucetAccount().PublicAccount
}

func (c *ConfigData) FaucetAccountAddress() string {
	return strings.ToUpper(c.FaucetAccount().Address.Address)
}

var networkType *sdk.NetworkType

func (c *ConfigData) NetworkType() sdk.NetworkType {
	if networkType != nil {
		return *networkType
	} else {
		nt := sdk.NetworkTypeFromString(c.Blockchain.Network)
		networkType = &nt
		return *networkType
	}
}
