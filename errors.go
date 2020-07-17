package Faucet

import "errors"

type ServerError error
type BadRequestError error
type NotAuthenticated error
type NotAuthorized error

var (
	AddressMissing         BadRequestError  = errors.New("address is missing")
	AddressInvalid         BadRequestError  = errors.New("address is invalid")
	AddressRegistered      BadRequestError  = errors.New("Test-XPX can only be sent once every 24 hours")
	IpAddressRegistered    BadRequestError  = errors.New("Test-XPX can only be sent once every 24 hours")
	RecordAlready          BadRequestError  = errors.New("record already exists")
	MaximumQuantity        BadRequestError  = errors.New("The account has the maximum amount of XPX")
	TryAgainLater          BadRequestError  = errors.New("Try again later")
	Unauthenticated        NotAuthenticated = errors.New("not authenticated")
	Unauthorized           NotAuthorized    = errors.New("not authorized")
	DbError                ServerError      = errors.New("database error")
	BlockchainApiError     ServerError      = errors.New("blockchain call failed")
	CreateMosaicError      ServerError      = errors.New("mosaic creation failed")
	CreateTransferTxnError ServerError      = errors.New("transfer txn creation failed")
	UnexpectedError        ServerError      = errors.New("unexpected error occurred")
	WebsocketError         ServerError      = errors.New("websocket error")
	MosaicInvalid          BadRequestError  = errors.New("mosaic is invalid")
	MosaicMissing          BadRequestError  = errors.New("mosaicId is missing")
)
