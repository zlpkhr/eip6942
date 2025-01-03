package eip6942

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var messagePrefix = "\x19Ethereum Signed Message:\n"

func hashMessage(message string) common.Hash {
	return crypto.Keccak256Hash([]byte(messagePrefix + strconv.Itoa(len(message)) + message))
}
