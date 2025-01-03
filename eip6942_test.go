package eip6942_test

import (
	"context"
	"eip6942"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type fixture struct {
	Address   string `json:"address"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
	Valid     bool   `json:"valid"`
}

func loadFixtures(path string) ([]fixture, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Join(errors.New("failed to read fixtures file"), err)
	}

	var fixtures []fixture
	if err := json.Unmarshal(data, &fixtures); err != nil {
		return nil, errors.Join(errors.New("failed to parse fixtures"), err)
	}

	return fixtures, nil
}

func TestValidateSignature(t *testing.T) {
	fixtures, err := loadFixtures("fixtures/eip6942.json")
	if err != nil {
		t.Fatalf("failed to load fixtures: %v", err)
	}

	client, err := rpc.Dial("https://polygon-rpc.com")

	if err != nil {
		t.Fatalf("failed to dial client: %v", err)
	}

	ec := &eip6942.Client{client}

	for _, fixture := range fixtures {
		err := ec.ValidateSignature(context.Background(), common.HexToAddress(fixture.Address), fixture.Message, hexutil.MustDecode(fixture.Signature))

		if fixture.Valid && err != nil {
			t.Fatalf("expected signature to be valid, but got error: %v", err)
		}

		if !fixture.Valid && err == nil {
			t.Fatalf("expected signature to be invalid, but got no error")
		}
	}
}
