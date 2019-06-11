// this acts as a simple in memory db to handle users and refresh tokens.
// replace these functions with actual calls to your db

package db

import (
	"github.com/proximax-storage/xpx-catapult-faucet"
	"sync"
	"time"
)

type reviews struct {
	date time.Time
}

var (
	once          sync.Once
	listByIP      map[string]reviews
	listByAddress map[string]reviews
)

func Init() {
	once.Do(blackList)
}

func blackList() {
	if Faucet.Config.BlackList.ByIp {
		listByIP = make(map[string]reviews)
	}
	if Faucet.Config.BlackList.ByAddress {
		listByAddress = make(map[string]reviews)
	}
}

func StoreClient(address, t string) error {
	// check to make sure our jti is unique
	if CheckBlackList(address, t) {
		return Faucet.RecordAlready
	}

	if t == "byAddress" {
		listByAddress[address] = reviews{date: time.Now().Add(24 * time.Hour)}
	}
	if t == "byIp" {
		listByIP[address] = reviews{date: time.Now().Add(24 * time.Hour)}
	}
	return nil
}

func CheckBlackList(value string, t string) bool {
	if t == "byAddress" {
		d, v := listByAddress[value]
		if v == true {
			if d.date.Unix() >= time.Now().Unix() {
				return true
			}
		}
	}

	if t == "byIp" {
		d, v := listByIP[value]
		if v == true {
			if d.date.Unix() >= time.Now().Unix() {
				return true
			}
		}
	}

	return false
}
