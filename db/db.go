// this acts as a simple in memory db to handle users and refresh tokens.
// replace these functions with actual calls to your db

package db

import (
	"github.com/proximax-storage/faucet-backend"
	"sync"
)

var (
	once          sync.Once
	listByIP      map[string]string
	listByAddress map[string]string
)

func Init() {
	once.Do(blackList)
}

func blackList() {
	if Faucet.Config.BlackList.ByIp {
		listByIP = make(map[string]string)
	}
	if Faucet.Config.BlackList.ByAddress {
		listByAddress = make(map[string]string)
	}
}

func StoreClient(address, t string) error {
	// check to make sure our jti is unique
	if CheckBlackList(address, t) {
		return Faucet.RecordAlready
	}

	if t == "byAddress" {
		listByAddress[address] = "valid"
	}
	if t == "byIp" {
		listByIP[address] = "valid"
	}
	return nil
}

func DeleteBlackList(value, t string) {
	if t == "byAddress" {
		delete(listByAddress, value)
	}
	if t == "byIp" {
		delete(listByIP, value)
	}
}

func CheckBlackList(value string, t string) bool {
	if t == "byAddress" {
		return listByAddress[value] != ""
	}
	if t == "byIp" {
		return listByIP[value] != ""
	}

	return false
}
