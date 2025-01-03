package eip6942

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrFailedToCreateType = errors.New("failed to create type")
	ErrFailedToPackArgs   = errors.New("failed to pack args")
)

func packCalldata(signer common.Address, messageHash common.Hash, signature []byte) ([]byte, error) {
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, errors.Join(ErrFailedToCreateType, err)
	}

	bytes32Type, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return nil, errors.Join(ErrFailedToCreateType, err)
	}

	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, errors.Join(ErrFailedToCreateType, err)
	}

	args := abi.Arguments{
		{Name: "signer", Type: addressType},
		{Name: "hash", Type: bytes32Type},
		{Name: "signature", Type: bytesType},
	}

	packed, err := args.Pack(signer, messageHash, signature)
	if err != nil {
		return nil, errors.Join(ErrFailedToPackArgs, err)
	}

	return append(validateSigOffchainBinBytes, packed...), nil
}
