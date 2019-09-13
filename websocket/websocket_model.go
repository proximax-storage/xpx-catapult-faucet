// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/binary"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"math/big"
)

type Uint64Dto [2]uint32

func (dto Uint64Dto) ToBigInt() *big.Int {
	if dto[0] == 0 && dto[1] == 0 {
		return &big.Int{}
	}
	var int big.Int
	b := make([]byte, len(dto)*4)
	binary.BigEndian.PutUint32(b[:len(dto)*2], dto[1])
	binary.BigEndian.PutUint32(b[len(dto)*2:], dto[0])

	int.SetBytes(b)
	return &int
}

type blockInfoDTO struct {
	BlockMeta struct {
		Hash            string    `json:"hash"`
		GenerationHash  string    `json:"generationHash"`
		TotalFee        Uint64Dto `json:"totalFee"`
		NumTransactions uint64    `json:"numTransactions"`
		// MerkleTree      uint64DTO `json:"merkleTree"` is needed?
	} `json:"meta"`
	Block struct {
		Signature             string    `json:"signature"`
		Signer                string    `json:"signer"`
		Version               uint64    `json:"version"`
		Type                  uint64    `json:"type"`
		Height                Uint64Dto `json:"height"`
		Timestamp             Uint64Dto `json:"timestamp"`
		Difficulty            Uint64Dto `json:"difficulty"`
		PreviousBlockHash     string    `json:"previousBlockHash"`
		BlockTransactionsHash string    `json:"blockTransactionsHash"`
	} `json:"block"`
}

type StatusInfo struct {
	Status string `json:"status"`
	Hash   string `json:"hash"`
}

type SignerInfo struct {
	Signer     string `json:"signer"`
	Signature  string `json:"signature"`
	ParentHash string `json:"parentHash"`
}

type ErrorInfo struct {
	Error error
}

// structure for Subscribe status
type HashInfo struct {
	Hash string `json:"hash"`
}

// structure for Subscribe PartialRemoved
type PartialRemovedInfo struct {
	Meta SubscribeHash `json:"meta"`
}

func (dto *blockInfoDTO) toStruct() (*sdk.BlockInfo, error) {
	nt := sdk.ExtractNetworkType(int64(dto.Block.Version))

	pa, err := sdk.NewAccountFromPublicKey(dto.Block.Signer, nt)
	if err != nil {
		return nil, err
	}

	v := sdk.ExtractVersion(int64(dto.Block.Version))

	generationHash, _ := sdk.StringToHash(dto.BlockMeta.GenerationHash)

	signature, _ := sdk.StringToSignature(dto.Block.Signature)

	previousBlockHash, _ := sdk.StringToHash(dto.Block.PreviousBlockHash)

	blockTransactionsHash, _ := sdk.StringToHash(dto.Block.BlockTransactionsHash)

	return &sdk.BlockInfo{
		NetworkType:           nt,
		GenerationHash:        generationHash,
		TotalFee:              sdk.Amount(dto.BlockMeta.TotalFee.ToBigInt().Int64()),
		NumTransactions:       dto.BlockMeta.NumTransactions,
		Signature:             signature,
		Signer:                pa,
		Version:               v,
		Type:                  dto.Block.Type,
		Height:                sdk.Height(dto.Block.Height.ToBigInt().Int64()),
		Timestamp:             sdk.NewTimestamp(dto.Block.Timestamp.ToBigInt().Int64()),
		Difficulty:            sdk.Difficulty(dto.Block.Difficulty.ToBigInt().Int64()),
		PreviousBlockHash:     previousBlockHash,
		BlockTransactionsHash: blockTransactionsHash,
	}, nil
}
