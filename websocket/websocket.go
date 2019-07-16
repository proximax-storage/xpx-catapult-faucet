// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	j "encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"golang.org/x/net/websocket"
	"io"
	"net/url"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var errChannels = make(map[string]chan *ErrorInfo)

type sendJson struct {
	Uid       string `json:"uid"`
	Subscribe string `json:"subscribe"`
}

type subscribeInfo struct {
	name, uid string
}

type serviceWs struct {
	client *ClientWebsocket
}

type subscribe struct {
	Uid       string `json:"uid"`
	Subscribe string `json:"subscribe"`
	conn      *websocket.Conn
	Ch        interface{}
}

// Catapult Websocket Client configuration
type ClientWebsocket struct {
	client    *websocket.Conn
	Uid       string
	duration  *time.Duration
	config    *sdk.Config
	common    serviceWs // Reuse a single struct instead of allocating one for each service on the heap.
	Subscribe *SubscribeService
}

type SubscribeBlock struct {
	*subscribe
	Ch chan *sdk.BlockInfo
}

type SubscribeTransaction struct {
	*subscribe
	Ch chan sdk.Transaction
}

type SubscribeHash struct {
	*subscribe
	Ch chan *HashInfo
}

type SubscribePartialRemoved struct {
	*subscribe
	Ch chan *PartialRemovedInfo
}

type SubscribeStatus struct {
	*subscribe
	Ch chan *StatusInfo
}

type SubscribeSigner struct {
	*subscribe
	Ch chan *SignerInfo
}

type SubscribeError struct {
	*subscribe
	Ch chan *ErrorInfo
}

func msgParser(msg []byte) (*subscribe, error) {
	var message subscribe
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func restParser(data []byte) (string, error) {
	var raw []j.RawMessage
	err := json.Unmarshal([]byte(fmt.Sprintf("[%v]", string(data))), &raw)
	if err != nil {
		return "", err
	}

	var subscribe string
	for _, r := range raw {
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return "", err
		}

		if _, ok := obj["block"]; ok {
			subscribe = "block"
		} else if _, ok := obj["status"]; ok {
			subscribe = "status"
		} else if _, ok := obj["signer"]; ok {
			subscribe = "signer"
		} else if v, ok := obj["meta"]; ok {
			channelName := v.(map[string]interface{})
			subscribe = fmt.Sprintf("%v", channelName["channelName"])
		} else {
			subscribe = "none"
		}
	}
	return subscribe, nil
}

func (s *subscribeInfo) buildType(t []byte) error {
	defer panicCtrl()
	switch s.name {
	case "block":
		var b blockInfoDTO
		err := json.Unmarshal(t, &b)
		if err != nil {
			return err
		}
		data, err := b.toStruct()
		if err != nil {
			return err
		}
		Block.Ch <- data
		return nil

	case "status":
		var data StatusInfo
		err := json.Unmarshal(t, &data)
		if err != nil {
			return err
		}
		ch := serviceCh.statusInfo.Get(s.uid)
		ch <- &data
		return nil

	case "signer":
		var data SignerInfo
		err := json.Unmarshal(t, data)
		if err != nil {
			return err
		}
		ch := serviceCh.signerInfo.Get(s.uid)
		ch <- &data
		return nil

	case "unconfirmedRemoved":
		var data HashInfo
		err := json.Unmarshal(t, data)
		if err != nil {
			return err
		}
		ch := serviceCh.unconfirmedRemoved.Get(s.uid)
		ch <- &data
		return nil

	case "partialRemoved":
		var data PartialRemovedInfo
		err := json.Unmarshal(t, &data)
		if err != nil {
			return err
		}
		ch := serviceCh.partialRemovedInfo.Get(s.uid)
		ch <- &data
		return nil

	case "partialAdded":
		data, err := sdk.MapTransaction(bytes.NewBuffer([]byte(t)))
		if err != nil {
			return err
		}
		ch := serviceCh.partialAdded.Get(s.uid)
		ch <- data
		return nil

	case "unconfirmedAdded":
		data, err := sdk.MapTransaction(bytes.NewBuffer([]byte(t)))
		if err != nil {
			return err
		}
		ch := serviceCh.unconfirmedAdded.Get(s.uid)
		ch <- data
		return nil

	default:
		data, err := sdk.MapTransaction(bytes.NewBuffer([]byte(t)))
		if err != nil {
			return err
		}
		ch := serviceCh.confirmedAdded.Get(s.uid)
		ch <- data

		return nil
	}
}

// Get subscribe name from subscribe struct
func (s *subscribe) getSubscribe() string {
	return strings.Split(s.Subscribe, "/")[0]
}

func (c *ClientWebsocket) changeURLPort() {
	c.config.BaseURLs[0].Scheme = "ws"
	c.config.BaseURLs[0].Path = "/ws"
	split := strings.Split(c.config.BaseURLs[0].Host, ":")
	host, port := split[0], "3000"
	c.config.BaseURLs[0].Host = strings.Join([]string{host, port}, ":")
}

func (c *ClientWebsocket) buildSubscribe(destination string) *subscribe {
	b := new(subscribe)
	b.Uid = c.Uid
	b.Subscribe = destination
	return b
}

func (c *ClientWebsocket) wsConnect() error {
	c.changeURLPort()
	var timeout <-chan time.Time
	if *c.duration != time.Duration(0) {
		timeout = time.After(*c.duration * time.Millisecond)
	}

	ticker := time.NewTicker(time.Millisecond)

	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timed out")

		case <-ticker.C:
			conn, err := websocket.Dial(c.config.BaseURLs[0].String(), "", "http://localhost")

			if err != nil {
				return err
			}
			c.client = conn

			var msg []byte
			if err = websocket.Message.Receive(c.client, &msg); err != nil {
				return err
			}

			imsg, err := msgParser(msg)
			if err != nil {
				return err
			}
			c.Uid = imsg.Uid
			return nil
		}
	}
}

func (c *ClientWebsocket) reconnectWs(s *subscribe) error {
	fmt.Println("Reconnecting Websocket....")

	if err := c.wsConnect(); err != nil {
		return err
	}

	s.Uid = c.Uid

	if err := websocket.JSON.Send(c.client, sendJson{
		Uid:       s.Uid,
		Subscribe: s.Subscribe,
	}); err != nil {
		return err
	}

	fmt.Println("New Websocket negotiated uid:", s.Uid)

	return nil
}

func (c *ClientWebsocket) subsChannel(s *subscribe) error {
	if err := websocket.JSON.Send(c.client, sendJson{
		Uid:       s.Uid,
		Subscribe: s.Subscribe,
	}); err != nil {
		return err
	}

	go func() {
		var resp []byte

		uid := "block"
		if s.Subscribe != "block" {
			uid = c.Uid
			serviceCh.connectsWs.conn(c.Uid, c.client)
		}
		errCh := errChannels[uid]

		for {
			if err := websocket.Message.Receive(c.client, &resp); err == io.EOF {
				if err != nil {
					errCh <- &ErrorInfo{
						Error: err,
					}
					return
				}

			} else if err != nil {
				if err != nil {
					errCh <- &ErrorInfo{
						Error: err,
					}
					break
				}

			} else {
				subName, err := restParser(resp)
				if err != nil {
					errCh <- &ErrorInfo{
						Error: err,
					}
					break
				}

				b := subscribeInfo{
					name: subName,
					uid:  s.Uid,
				}

				if err := b.buildType(resp); err != nil {
					errCh <- &ErrorInfo{
						Error: err,
					}
				}
			}

			if *c.duration != time.Duration(0) {
				tout := time.Now().Add(*c.duration * time.Millisecond)
				c.client.SetDeadline(tout)
			}
		}
	}()
	return nil
}

func (t *ClientWebsocket) Close() {
	t.client.Close()
}
func NewConnectWs(host string, timeout time.Duration) (*ClientWebsocket, error) {
	var v []*url.URL
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	} else {
		v = append(v, u)
	}

	newconf := &sdk.Config{BaseURLs: v}
	c := &ClientWebsocket{config: newconf}
	c.common.client = c
	c.Subscribe = (*SubscribeService)(&c.common)
	c.duration = &timeout

	err = c.wsConnect()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *SubscribeBlock) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func (s *SubscribeTransaction) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func (s *SubscribeHash) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func (s *SubscribePartialRemoved) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func (s *SubscribeStatus) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func (s *SubscribeSigner) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func (s *SubscribeError) Unsubscribe() error {
	return s.subscribe.unsubscribe()
}

func panicCtrl() {
	if r := recover(); r != nil {
	}
}
