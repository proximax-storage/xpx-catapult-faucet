// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package websocket

import (
	"errors"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"golang.org/x/net/websocket"
)

var Block *SubscribeBlock

var serviceCh *concurrentMaps

func init() {
	serviceCh = NewConcurrentMaps()
}

type SubscribeService serviceWs

// const routers path for methods SubscribeService
const (
	pathBlock              = "block"
	pathConfirmedAdded     = "confirmedAdded"
	pathUnconfirmedAdded   = "unconfirmedAdded"
	pathUnconfirmedRemoved = "unconfirmedRemoved"
	pathStatus             = "status"
	pathPartialAdded       = "partialAdded"
	pathPartialRemoved     = "partialRemoved"
	pathCosignature        = "cosignature"
)

// Closes the subscription channel.
func (s *subscribe) closeChannel() error {
	switch s.Ch.(type) {

	case chan *sdk.BlockInfo:
		chType := s.Ch.(chan *sdk.BlockInfo)
		close(chType)

	case chan *ErrorInfo:
		chType := s.Ch.(chan *ErrorInfo)
		delete(errChannels, s.Uid)
		close(chType)

	case chan *StatusInfo:
		serviceCh.statusInfo.Del(s.Uid)

	case chan *HashInfo:
		serviceCh.unconfirmedRemoved.Del(s.Uid)

	case chan *PartialRemovedInfo:
		serviceCh.partialRemovedInfo.Del(s.Uid)

	case chan *SignerInfo:
		serviceCh.statusInfo.Del(s.Uid)

	case chan sdk.Transaction:
		if s.getSubscribe() == "partialAdded" {
			serviceCh.partialAdded.Del(s.Uid)
		} else if s.getSubscribe() == "unconfirmedAdded" {
			serviceCh.unconfirmedAdded.Del(s.Uid)
		} else {
			serviceCh.confirmedAdded.Del(s.Uid)
		}

	default:
		return errors.New("WRONG TYPE CHANNEL")
	}
	return nil
}

// Unsubscribe terminates the specified subscription.
// It does not have any specific param.
func (c *subscribe) unsubscribe() error {
	if err := websocket.JSON.Send(c.conn, sendJson{
		Uid:       c.Uid,
		Subscribe: c.Subscribe,
	}); err != nil {
		return err
	}

	if err := c.closeChannel(); err != nil {
		return err
	}

	return nil
}

// Generate a new channel and subscribe to the websocket.
// param route A subscription channel route.
// return A pointer Subscribe struct or an error.
func (c *SubscribeService) newSubscribe(route string, ch interface{}) (*subscribe, error) {
	subMsg := c.client.buildSubscribe(route)

	err := c.client.subsChannel(subMsg)
	if err != nil {
		return nil, err
	}

	subMsg.conn = c.client.client
	subMsg.Ch = ch
	return subMsg, nil
}

// Block notifies for every new block.
// The message contains the BlockInfo struct.
func (c *SubscribeService) Block() (*SubscribeBlock, error) {
	subBlock := new(SubscribeBlock)
	Block = subBlock
	subBlock.Ch = make(chan *sdk.BlockInfo)
	subscribe, err := c.newSubscribe(pathBlock, subBlock.Ch)
	if err != nil {
		return nil, err
	}
	subBlock.subscribe = subscribe
	return subBlock, nil
}

// ConfirmedAdded notifies when a transaction related to an
// address is included in a block.
// The message contains the transaction.
func (c *SubscribeService) ConfirmedAdded(add *sdk.Address) (*SubscribeTransaction, error) {
	subTransaction := new(SubscribeTransaction)
	if serviceCh.confirmedAdded.Add(c.client.Uid) {
		subTransaction.Ch = serviceCh.confirmedAdded.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathConfirmedAdded+"/"+add.Address, subTransaction.Ch)
	if err != nil {
		return nil, err
	}

	subTransaction.subscribe = subscribe
	return subTransaction, nil
}

// UnconfirmedAdded notifies when a transaction related to an
// address is in unconfirmed state and waiting to be included in a block.
// The message contains the transaction.
func (c *SubscribeService) UnconfirmedAdded(add *sdk.Address) (*SubscribeTransaction, error) {
	subTransaction := new(SubscribeTransaction)
	if serviceCh.unconfirmedAdded.Add(c.client.Uid) {
		subTransaction.Ch = serviceCh.unconfirmedAdded.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathUnconfirmedAdded+"/"+add.Address, subTransaction.Ch)
	if err != nil {
		return nil, err
	}

	subTransaction.subscribe = subscribe
	return subTransaction, nil
}

// UnconfirmedRemoved notifies when a transaction related to an
// address was in unconfirmed state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) UnconfirmedRemoved(add *sdk.Address) (*SubscribeHash, error) {
	subHash := new(SubscribeHash)
	if serviceCh.unconfirmedRemoved.Add(c.client.Uid) {
		subHash.Ch = serviceCh.unconfirmedRemoved.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathUnconfirmedRemoved+"/"+add.Address, subHash.Ch)
	if err != nil {
		return nil, err
	}

	subHash.subscribe = subscribe
	return subHash, nil
}

// Status notifies when a transaction related to an address rises an error.
// The message contains the error message and the transaction hash.
func (c *SubscribeService) Status(add *sdk.Address) (*SubscribeStatus, error) {
	subStatus := new(SubscribeStatus)
	if serviceCh.statusInfo.Add(c.client.Uid) {
		subStatus.Ch = serviceCh.statusInfo.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathStatus+"/"+add.Address, subStatus.Ch)
	if err != nil {
		return nil, err
	}

	subStatus.subscribe = subscribe
	return subStatus, nil
}

// PartialAdded notifies when an aggregate bonded transaction related to an
// address is in partial state and waiting to have all required cosigners.
// The message contains a transaction.
func (c *SubscribeService) PartialAdded(add *sdk.Address) (*SubscribeTransaction, error) {
	subTransaction := new(SubscribeTransaction)
	if serviceCh.partialAdded.Add(c.client.Uid) {
		subTransaction.Ch = serviceCh.partialAdded.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathPartialAdded+"/"+add.Address, subTransaction.Ch)
	if err != nil {
		return nil, err
	}

	subTransaction.subscribe = subscribe
	return subTransaction, nil
}

// PartialRemoved notifies when a transaction related to an
// address was in partial state but not anymore.
// The message contains the transaction hash.
func (c *SubscribeService) PartialRemoved(add *sdk.Address) (*SubscribePartialRemoved, error) {
	subPartialRemoved := new(SubscribePartialRemoved)
	if serviceCh.partialRemovedInfo.Add(c.client.Uid) {
		subPartialRemoved.Ch = serviceCh.partialRemovedInfo.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathPartialRemoved+"/"+add.Address, subPartialRemoved.Ch)
	if err != nil {
		return nil, err
	}

	subPartialRemoved.subscribe = subscribe
	return subPartialRemoved, nil
}

// Cosignature notifies when a cosignature signed transaction related to an
// address is added to an aggregate bonded transaction with partial state.
// The message contains the cosignature signed transaction.
func (c *SubscribeService) Cosignature(add *sdk.Address) (*SubscribeSigner, error) {
	subCosignature := new(SubscribeSigner)
	if serviceCh.partialRemovedInfo.Add(c.client.Uid) {
		subCosignature.Ch = serviceCh.signerInfo.Get(c.client.Uid)
	}

	subscribe, err := c.newSubscribe(pathCosignature+"/"+add.Address, subCosignature.Ch)
	if err != nil {
		return nil, err
	}

	subCosignature.subscribe = subscribe
	return subCosignature, nil
}

func (c *SubscribeService) Error(add *sdk.Address) (*SubscribeError, error) {
	address := "block"
	if add != nil {
		address = add.Address
	}

	subError := new(SubscribeError)
	subError.Ch = make(chan *ErrorInfo)
	errChannels[address] = subError.Ch
	subscribe := new(subscribe)
	subscribe.Subscribe = "error/" + address
	subError.subscribe = subscribe
	subscribe.Ch = errChannels[c.client.Uid]
	return subError, nil
}
