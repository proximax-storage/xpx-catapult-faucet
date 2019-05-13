package blockchain

import (
	"context"
	"fmt"
	"github.com/proximax-storage/faucet-backend"
	"github.com/proximax-storage/faucet-backend/db"
	"github.com/proximax-storage/faucet-backend/utils"
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk"
	"math/big"
	"strings"
	"time"
)

func TransferXpx(Address string) error {

	if Faucet.Config.BlackList.ByIp {
		err := db.StoreClient(Address, "byIp")
		if err != nil {
			return Faucet.IpAddressRegistered
		}
	}
	if Faucet.Config.BlackList.ByAddress {
		err := db.StoreClient(Address, "byAddress")
		if err != nil {
			return Faucet.AddressRegistered
		}
	}

	return createTransfer(Address)
}

func AnnounceTxn(signedTxn *sdk.SignedTransaction) error {
	// announce transaction
	utils.Logger(1, "Connecting to the node: %v", Faucet.Config.Blockchain.ApiUrl)
	_, err := Faucet.BlockchainClient.Transaction.Announce(context.Background(), signedTxn)
	if err != nil {
		utils.Logger(3, "Failed to announce status: %v", err)
		return Faucet.BlockchainApiError
	}
	return nil
}

func createTransfer(Address string) error {

	add := sdk.NewAddress(Address, Faucet.Config.NetworkType())

	var balance int64

	restTx, err := Faucet.BlockchainClient.Account.GetAccountInfo(context.Background(), add)
	if err != nil {
		balance = 0
	} else {
		for _, m := range restTx.Mosaics {
			id := sdk.BigIntegerToHex((*big.Int)(m.MosaicId))
			if strings.ToUpper(id) == strings.ToUpper(Faucet.Config.App.MosaicId) {
				balance = m.Amount.Int64()
			}
		}
		if balance >= Faucet.Config.App.MaxXpx {
			return Faucet.MaximumQuantity
		}
	}

	ttx, err := sdk.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		add,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(Faucet.Config.App.MaxXpx - balance)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Sirius faucet"),
		Faucet.Config.NetworkType(),
	)

	// Sign transaction
	signedTxn, err := Faucet.Config.FaucetAccount().Sign(ttx)
	if err != nil {
		return fmt.Errorf("TransaferTransaction signing returned error: %s", err)
	}

	err = AnnounceTxn(signedTxn)
	if err != nil {
		return err
	}

	return nil
}
