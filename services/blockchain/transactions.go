package blockchain

import (
	"context"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"github.com/proximax-storage/xpx-catapult-faucet"
	"github.com/proximax-storage/xpx-catapult-faucet/db"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
	"strings"
	"time"
)

func TransferXpx(Address, ip string, id *sdk.NamespaceId) error {
	Address = strings.Replace(Address, "-", "", -1)

	for _, x := range Faucet.Config.WhiteList.Addresses {
		if Address == x {
			return createTransfer(Address, id)
		}
	}

	if Faucet.Config.BlackList.ByIp {
		if db.DbStorage.GetIp(ip) {
			return Faucet.IpAddressRegistered
		}
	}

	if Faucet.Config.BlackList.ByAddress {
		if db.DbStorage.GetAddress(Address) {
			return Faucet.AddressRegistered
		}
	}

	if err := createTransfer(Address, id); err != nil {
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
		return Faucet.WebsocketError
	}
	defer ws.Close()

	// open websocket to wait for status to be validated, either it end up as unconfirmed or failed
	// listen to either websocket
	unconfirmed, err := ws.Subscribe.UnconfirmedAdded(address)
	if err != nil {
		utils.Logger(3, "Failed to open websocket for unconfirmed txn: %v", err)
		return Faucet.WebsocketError
	}
	defer unconfirmed.Unsubscribe()

	confirmed, err := ws.Subscribe.ConfirmedAdded(address)
	if err != nil {
		utils.Logger(3, "Failed to open websocket for confirmed txn: %v", err)
		return Faucet.WebsocketError
	}
	defer confirmed.Unsubscribe()

	status, err := ws.Subscribe.Status(address)
	if err != nil {
		utils.Logger(3, "Failed to open websocket for status: %v", err)
		return Faucet.WebsocketError
	}
	defer status.Unsubscribe()

	// announce transaction
	utils.Logger(1, "Connecting to the node: %v", Faucet.Config.Blockchain.ApiUrl)
	_, err = Faucet.BlockchainClient.Transaction.Announce(context.Background(), signedTxn)
	if err != nil {
		utils.Logger(3, "Failed to announce status: %v", err)
		return Faucet.BlockchainApiError
	}

	for {
		select {
		case data := <-unconfirmed.Ch:
			if data.GetAbstractTransaction().TransactionHash.String() == signedTxn.Hash.String() {
				utils.Logger(0, "Unconfirmed transaction hash -> %v", data.GetAbstractTransaction().TransactionHash)
				return nil
			}

		case data := <-confirmed.Ch:
			if data.GetAbstractTransaction().TransactionHash.String() == signedTxn.Hash.String() {
				utils.Logger(0, "Confirmed transaction hash -> %v", data.GetAbstractTransaction().TransactionHash)
				return nil
			}

		case data := <-status.Ch:
			if strings.ToUpper(data.Hash) == strings.ToUpper(signedTxn.Hash.String()) {
				utils.Logger(2, "%v", data.Status)
				return fmt.Errorf("%v", strings.Replace(strings.Split(data.Status, "Failure_Core_")[1], "_", " ", 1))
			}
		}
	}
}

func createTransfer(Address string, MosaicId *sdk.NamespaceId) error {
	var mosaicId *Faucet.MosaicInfo
	for _, mosaic := range Faucet.Config.App.Mosaics {
		nsId, err := sdk.NewNamespaceIdFromName(mosaic.Name)
		if err != nil {
			utils.Logger(3, "Failed NamespaceIdFromName: %v", err)
			return err
		}
		if nsId.Id() == MosaicId.Id() {
			mosaicId = &mosaic
			break
		}
	}
	if mosaicId == nil {
		utils.Logger(3, "Mosaic invalid: %v", MosaicId.String())
		return Faucet.MosaicInvalid
	}

	mosaicInfo, err := Faucet.BlockchainClient.Resolve.GetMosaicInfoByAssetId(context.Background(), MosaicId)

	add := sdk.NewAddress(Address, Faucet.Config.NetworkType())

	var balance sdk.Amount

	restTx, err := Faucet.BlockchainClient.Account.GetAccountInfo(context.Background(), add)
	if err != nil {
		balance = 0
	} else {

		publicAccount, err := sdk.NewAccountFromPublicKey(restTx.PublicKey, Faucet.Config.NetworkType())
		if err != nil {
			return err
		}

		unconfirmedTx, err := Faucet.BlockchainClient.Account.UnconfirmedTransactions(context.Background(), publicAccount, nil)
		if err != nil {
			return err
		}

		if len(unconfirmedTx) != 0 {
			return Faucet.TryAgainLater
		}

		for _, m := range restTx.Mosaics {
			id := m.AssetId.String()

			if strings.ToUpper(id) == strings.ToUpper(mosaicInfo.MosaicId.String()) {
				balance = m.Amount
			}
		}

		if balance >= mosaicId.MaxQuantity {
			return Faucet.MaximumQuantity
		}
	}

	Mosaic, err := sdk.NewMosaic(MosaicId, mosaicId.MaxQuantity-balance)

	ttx, err := sdk.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		add,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{Mosaic},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Sirius faucet"),
		Faucet.Config.NetworkType(),
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
