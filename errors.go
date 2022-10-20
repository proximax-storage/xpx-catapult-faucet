package Faucet

import "errors"

type ServerError error
type BadRequestError error
type NotAuthenticated error
type NotAuthorized error

var (
	ErrAddressMissing         BadRequestError  = errors.New("address is missing")
	ErrAddressInvalid         BadRequestError  = errors.New("address is invalid")
	ErrAddressRegistered      BadRequestError  = errors.New("Test-XPX can only be sent once every 24 hours")
	ErrIpAddressRegistered    BadRequestError  = errors.New("Test-XPX can only be sent once every 24 hours")
	ErrRecordAlready          BadRequestError  = errors.New("record already exists")
	ErrMaximumQuantity        BadRequestError  = errors.New("account has the maximum amount of XPX")
	ErrUnauthenticated        NotAuthenticated = errors.New("not authenticated")
	ErrUnauthorized           NotAuthorized    = errors.New("not authorized")
	ErrDbError                ServerError      = errors.New("database error")
	ErrBlockchainApi     ServerError      = errors.New("blockchain call failed")
	ErrCreateMosaic      ServerError      = errors.New("mosaic creation failed")
	ErrCreateTransferTxn ServerError      = errors.New("transfer txn creation failed")
	ErrUnexpected        ServerError      = errors.New("unexpected error occurred")
	ErrWebsocket         ServerError      = errors.New("websocket error")
)
