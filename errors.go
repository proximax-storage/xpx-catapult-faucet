package Faucet

import "errors"

type ServerError error
type BadRequestError error
type NotAuthenticated error
type NotAuthorized error

var (
	AddressMissing         BadRequestError  = errors.New("address is missing")
	AddressInvalid         BadRequestError  = errors.New("address is invalid")
	AddressRegistered      BadRequestError  = errors.New("XPX can only be send once every 24 hours")
	IpAddressRegistered    BadRequestError  = errors.New("XPX can only be send once every 24 hours")
	RecordAlready          BadRequestError  = errors.New("record already exists")
	MaximumQuantity        BadRequestError  = errors.New("The account has the maximum amount of XPX")
	Unauthenticated        NotAuthenticated = errors.New("not authenticated")
	Unauthorized           NotAuthorized    = errors.New("not authorized")
	DbError                ServerError      = errors.New("database error")
	BlockchainApiError     ServerError      = errors.New("blockchain call failed")
	CreateMosaicError      ServerError      = errors.New("mosaic creation failed")
	CreateTransferTxnError ServerError      = errors.New("transfer txn creation failed")
	UnexpectedError        ServerError      = errors.New("unexpected error occurred")
	WebsocketError         ServerError      = errors.New("websocket error")
)
