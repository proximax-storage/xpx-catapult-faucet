package websocket

import (
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"golang.org/x/net/websocket"
	"sync"
)

type concurrentMaps struct {
	*confirmedAdded
	*connectsWs
	*statusInfo
	*partialRemovedInfo
	*signerInfo
	*unconfirmedRemoved
	*partialAdded
	*unconfirmedAdded
}

type connectsWs struct {
	sync.Mutex
	connectsWs map[string]*websocket.Conn
}

type statusInfo struct {
	sync.Mutex
	statusInfoChannels map[string]chan *StatusInfo
}

type signerInfo struct {
	sync.Mutex
	signerInfoChannels map[string]chan *SignerInfo
}

type partialAdded struct {
	sync.Mutex
	partialAddedChannels map[string]chan sdk.Transaction
}

type confirmedAdded struct {
	sync.Mutex
	confirmedAddedChannels map[string]chan sdk.Transaction
}

type unconfirmedAdded struct {
	sync.Mutex
	unconfirmedAddedChannels map[string]chan sdk.Transaction
}

type unconfirmedRemoved struct {
	sync.Mutex
	unconfirmedRemovedChannels map[string]chan *HashInfo
}

type partialRemovedInfo struct {
	sync.Mutex
	partialRemovedInfoChannels map[string]chan *PartialRemovedInfo
}

func (u *connectsWs) Get(uid string) *websocket.Conn {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.connectsWs[uid]; !ok {
		return nil
	}
	return u.connectsWs[uid]
}

func (u *connectsWs) conn(str string, ch *websocket.Conn) bool {
	u.Lock()
	defer u.Unlock()

	if len(u.connectsWs) != 0 {
		u.connectsWs[str] = ch
		return true
	}

	return false
}

func (u *statusInfo) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.statusInfoChannels[uid]; ok {
		return false
	}
	u.statusInfoChannels[uid] = make(chan *StatusInfo)
	return true
}

func (u *statusInfo) Get(uid string) chan *StatusInfo {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.statusInfoChannels[uid]; !ok {
		return nil
	}
	return u.statusInfoChannels[uid]
}

func (u *statusInfo) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.statusInfoChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.statusInfoChannels, uid)

	return true
}

func (u *signerInfo) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.signerInfoChannels[uid]; ok {
		return false
	}
	u.signerInfoChannels[uid] = make(chan *SignerInfo)
	return true
}

func (u *signerInfo) Get(uid string) chan *SignerInfo {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.signerInfoChannels[uid]; !ok {
		return nil
	}
	return u.signerInfoChannels[uid]
}

func (u *signerInfo) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.signerInfoChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.signerInfoChannels, uid)

	return true
}

func (u *partialAdded) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.partialAddedChannels[uid]; ok {
		return false
	}
	u.partialAddedChannels[uid] = make(chan sdk.Transaction)
	return true
}

func (u *partialAdded) Get(uid string) chan sdk.Transaction {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.partialAddedChannels[uid]; !ok {
		return nil
	}
	return u.partialAddedChannels[uid]
}

func (u *partialAdded) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.partialAddedChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.partialAddedChannels, uid)

	return true
}

func (u *confirmedAdded) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.confirmedAddedChannels[uid]; ok {
		return false
	}
	u.confirmedAddedChannels[uid] = make(chan sdk.Transaction)
	return true
}

func (u *confirmedAdded) Get(uid string) chan sdk.Transaction {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.confirmedAddedChannels[uid]; !ok {
		return nil
	}
	return u.confirmedAddedChannels[uid]
}

func (u *confirmedAdded) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.confirmedAddedChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.confirmedAddedChannels, uid)

	return true
}

func (u *unconfirmedAdded) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.unconfirmedAddedChannels[uid]; ok {
		return false
	}
	u.unconfirmedAddedChannels[uid] = make(chan sdk.Transaction)
	return true
}

func (u *unconfirmedAdded) Get(uid string) chan sdk.Transaction {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.unconfirmedAddedChannels[uid]; !ok {
		return nil
	}
	return u.unconfirmedAddedChannels[uid]
}

func (u *unconfirmedAdded) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.unconfirmedAddedChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.unconfirmedAddedChannels, uid)

	return true
}

func (u *partialRemovedInfo) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.partialRemovedInfoChannels[uid]; ok {
		return false
	}
	u.partialRemovedInfoChannels[uid] = make(chan *PartialRemovedInfo)
	return true
}

func (u *partialRemovedInfo) Get(uid string) chan *PartialRemovedInfo {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.partialRemovedInfoChannels[uid]; !ok {
		return nil
	}
	return u.partialRemovedInfoChannels[uid]
}

func (u *partialRemovedInfo) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.partialRemovedInfoChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.partialRemovedInfoChannels, uid)

	return true
}

func (u *unconfirmedRemoved) Add(uid string) bool {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.unconfirmedRemovedChannels[uid]; ok {
		return false
	}
	u.unconfirmedRemovedChannels[uid] = make(chan *HashInfo)
	return true
}

func (u *unconfirmedRemoved) Get(uid string) chan *HashInfo {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.unconfirmedRemovedChannels[uid]; !ok {
		return nil
	}
	return u.unconfirmedRemovedChannels[uid]
}

func (u *unconfirmedRemoved) Del(uid string) bool {
	u.Lock()
	defer u.Unlock()

	if val, ok := u.unconfirmedRemovedChannels[uid]; ok {
		close(val)
	} else {
		return false
	}

	delete(u.unconfirmedRemovedChannels, uid)

	return true
}

func NewConcurrentMaps() *concurrentMaps {
	return &concurrentMaps{
		&confirmedAdded{
			confirmedAddedChannels: make(map[string]chan sdk.Transaction),
		},
		&connectsWs{
			connectsWs: make(map[string]*websocket.Conn),
		},
		&statusInfo{
			statusInfoChannels: make(map[string]chan *StatusInfo),
		},
		&partialRemovedInfo{
			partialRemovedInfoChannels: make(map[string]chan *PartialRemovedInfo),
		},
		&signerInfo{
			signerInfoChannels: make(map[string]chan *SignerInfo),
		},
		&unconfirmedRemoved{
			unconfirmedRemovedChannels: make(map[string]chan *HashInfo),
		},
		&partialAdded{
			partialAddedChannels: make(map[string]chan sdk.Transaction),
		},
		&unconfirmedAdded{
			unconfirmedAddedChannels: make(map[string]chan sdk.Transaction),
		},
	}
}
