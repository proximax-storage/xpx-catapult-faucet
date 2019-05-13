package utils

import (
	e "errors"
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk"
	"regexp"
	"strings"
)

// Check if a address is valid
// param address - A address string
// param nt - A NetworkType
// return - error or nil
func IsAddressValid(address string, networkType sdk.NetworkType) error {
	address = strings.Replace(address, "-", "", -1)
	address = strings.ToUpper(address)

	if len(address) != 40 {
		return e.New("address length must be 40 characters")
	} else if AddressNet[address[0]] != networkType {
		return e.New("unsupported network")
	} else {
		return nil
	}
}

// Test if a string is hexadecimal
// param str - A string to test
// return True if correct, false otherwise
func IsHexadecimal(str string) bool {
	exp := regexp.MustCompile("^[0-9a-fA-F]+$")
	if exp.MatchString(str) == true {
		return true
	}
	return false
}
func Logger(level Level, format string, args ...interface{}) {
	logger(level, 0, format, args...)
}
