package utils

import (
	"github.com/json-iterator/go"
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var AddressNet = map[uint8]sdk.NetworkType{
	'M': sdk.Mijin,
	'S': sdk.MijinTest,
	'X': sdk.Public,
	'V': sdk.PublicTest,
	'Z': sdk.Private,
	'W': sdk.PrivateTest,
}
