package blockchain

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	Faucet "github.com/proximax-storage/xpx-catapult-faucet"
	"github.com/proximax-storage/xpx-catapult-faucet/db"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
)

func TransferXpx(Address, ip string) error {
	Address = strings.Replace(Address, "-", "", -1)

	for _, x := range Faucet.Config.WhiteList.Addresses {
		if Address == x {
			return createTransfer(Address)
		}
	}

	if Faucet.Config.BlackList.ByIp {
		if db.DbStorage.GetIp(ip) {
			return Faucet.ErrIpAddressRegistered
		}
	}

	if Faucet.Config.BlackList.ByAddress {
		if db.DbStorage.GetAddress(Address) {
			return Faucet.ErrAddressRegistered
		}
	}

	if err := createTransfer(Address); err != nil {
		return err
	}

	if Faucet.Config.BlackList.ByIp {
		err := db.DbStorage.AddIp(ip)
		if err != nil {
			return err
		}
	}

	if Faucet.Config.BlackList.ByAddress {
		err := db.DbStorage.AddAddress(Address)
		if err != nil {
			return err
		}
	}
	return nil
}

func announceTxn(signedTxn *sdk.SignedTransaction) error {

	address := Faucet.Config.FaucetAccount().Address

	ws, err := Faucet.NewWebsocket()
	if err != nil {
		utils.Logger(3, "Failed to create websocket: %v", err)
		return Faucet.ErrWebsocket
	}
	defer ws.Close()

	// open websocket to wait for status to be validated, either it end up as unconfirmed or failed
	// listen to either websocket
	unconfirmed, err := ws.Subscribe.UnconfirmedAdded(address)
	if err != nil {
		utils.Logger(3, "Failed to open websocket for unconfirmed txn: %v", err)
		return Faucet.ErrWebsocket
	}
	defer unconfirmed.Unsubscribe()

	confirmed, err := ws.Subscribe.ConfirmedAdded(address)
	if err != nil {
		utils.Logger(3, "Failed to open websocket for confirmed txn: %v", err)
		return Faucet.ErrWebsocket
	}
	defer confirmed.Unsubscribe()

	status, err := ws.Subscribe.Status(address)
	if err != nil {
		utils.Logger(3, "Failed to open websocket for status: %v", err)
		return Faucet.ErrWebsocket
	}
	defer status.Unsubscribe()

	// announce transaction
	utils.Logger(1, "Connecting to the node: %v", Faucet.Config.Blockchain.ApiUrl)
	_, err = Faucet.BlockchainClient.Transaction.Announce(context.Background(), signedTxn)
	if err != nil {
		utils.Logger(3, "Failed to announce status: %v", err)
		return Faucet.ErrBlockchainApi
	}

	for {
		select {
		case data := <-unconfirmed.Ch:
			if data.GetAbstractTransaction().TransactionHash.String() == signedTxn.Hash.String() {
				utils.Logger(0, "Transaction hash -> %v", data.GetAbstractTransaction().TransactionHash)
				return nil
			}

		case data := <-confirmed.Ch:
			if data.GetAbstractTransaction().TransactionHash.String() == signedTxn.Hash.String() {
				utils.Logger(0, "Transaction hash -> %v", data.GetAbstractTransaction().TransactionHash)
				return nil
			}

		case data := <-status.Ch:
			if data.Hash == signedTxn.Hash.String() {
				utils.Logger(2, "%v", data.Status)
				return fmt.Errorf("%v", strings.Replace(strings.Split(data.Status, "Failure_Core_")[1], "_", " ", 1))
			}
		}
	}
}

func createTransfer(Address string) error {

	add := sdk.NewAddress(Address, Faucet.Config.NetworkType())

	var balance sdk.Amount

	restTx, err := Faucet.BlockchainClient.Account.GetAccountInfo(context.Background(), add)
	if err != nil {
		balance = 0
	} else {
		for _, m := range restTx.Mosaics {
			id := m.AssetId.String()

			if strings.EqualFold(strings.ToUpper(id), strings.ToUpper(Faucet.Config.App.MosaicId)) {
				balance = m.Amount
			}
		}
		if balance >= Faucet.Config.App.MaxXpx {
			return Faucet.ErrMaximumQuantity
		}
	}

	ttx, _ := Faucet.BlockchainClient.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		add,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(uint64(Faucet.Config.App.MaxXpx - balance))},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Sirius faucet"),
	)

	// Sign transaction
	signedTxn, err := Faucet.Config.FaucetAccount().Sign(ttx)
	if err != nil {
		return fmt.Errorf("TransaferTransaction signing returned error: %s", err)
	}

	err = announceTxn(signedTxn)
	if err != nil {
		return err
	}

	return nil
}

//// analog JAVA Uint64.bigIntegerToHex
//func bigIntegerToHex(id *big.Int) string {
//	u := fromBigInt(id)
//	return strings.ToUpper(intToHex(u[1]) + intToHex(u[0]))
//}
//
//func intToHex(u uint32) string {
//	return fmt.Sprintf("%08x", u)
//}
//
//func fromBigInt(int *big.Int) []uint32 {
//	if int == nil {
//		return []uint32{0, 0}
//	}
//
//	var u64 = uint64(int.Int64())
//	l := uint32(u64 & 0xFFFFFFFF)
//	r := uint32(u64 >> 32)
//	return []uint32{l, r}
//}
