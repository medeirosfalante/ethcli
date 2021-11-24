package ethcli_test

import (
	"github.com/medeirosfalante/ethcli/tron"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"math/big"
	"testing"
)

const (
	TestGrpcAddress = "47.252.19.181:50051"
	TestMnemonic    = "crunch goat helmet buddy metal crowd fish rack entire glare job milk call wish protect"
)

func TestTronClient_IsTRC10(t *testing.T) {

	client, err := tron.NewTronClient(TestGrpcAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	t.Run("must not return an error", func(t *testing.T) {
		ok := client.IsTRC10("1000016")
		assert.Equal(t, true, ok)
	})

	t.Run("must return an error", func(t *testing.T) {
		ok := client.IsTRC10("TRYCsQN7mfyCVYvXQuc2LTAUC4EDxDMoMJ")
		assert.Equal(t, false, ok)
	})
}

func TestTronClient_Transfer(t *testing.T) {

	client, err := tron.NewTronClient(TestGrpcAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	tronAccount, err := tron.NewTronAccount(TestMnemonic, "0")

	assert.NoError(t, err)
	assert.NotEmpty(t, tronAccount)

	t.Run("Must not return an error on send Native coin", func(t *testing.T) {

		tx, err := client.Transfer(&tron.TransferOpts{
			From:    tronAccount.PublicKey(),
			PrivKey: tronAccount.PrivateKey(),
			To:      "TVvxthSkGf4FGd8B5C2JQdQYPprSQgSvJX",
			Amount:  big.NewInt(1e2),
		})

		assert.NoError(t, err)
		assert.NotEqual(t, "", tx)
	})

	t.Run("Must not return an error on send TRC10", func(t *testing.T) {
		tokenID := "1000016"

		tx, err := client.Transfer(&tron.TransferOpts{
			From:    tronAccount.PublicKey(),
			PrivKey: tronAccount.PrivateKey(),
			To:      "TVvxthSkGf4FGd8B5C2JQdQYPprSQgSvJX",
			Amount:  big.NewInt(1e2),
			TokenID: &tokenID,
		})

		assert.NoError(t, err)
		assert.NotEqual(t, "", tx)
	})

	t.Run("Must not return an error on send TRC20", func(t *testing.T) {
		tokenID := "TU2T8vpHZhCNY8fXGVaHyeZrKm8s6HEXWe"

		tx, err := client.Transfer(&tron.TransferOpts{
			From:    tronAccount.PublicKey(),
			PrivKey: tronAccount.PrivateKey(),
			To:      "TVvxthSkGf4FGd8B5C2JQdQYPprSQgSvJX",
			Amount:  big.NewInt(1e1),
			TokenID: &tokenID,
		})

		assert.NoError(t, err)
		assert.NotEqual(t, "", tx)
	})

	t.Run("Must return an error because token id is invalid", func(t *testing.T) {

		tokenID := "10000164654654"

		tx, err := client.Transfer(&tron.TransferOpts{
			From:    tronAccount.PublicKey(),
			PrivKey: tronAccount.PrivateKey(),
			To:      "TVvxthSkGf4FGd8B5C2JQdQYPprSQgSvJX",
			Amount:  big.NewInt(1e1),
			TokenID: &tokenID,
		})

		assert.Error(t, err)
		assert.Empty(t, tx)
	})

	t.Run("Must return an error because token contract address is invalid", func(t *testing.T) {

		tokenID := "10000164654654"

		tx, err := client.Transfer(&tron.TransferOpts{
			From:    tronAccount.PublicKey(),
			PrivKey: tronAccount.PrivateKey(),
			To:      "TVvxthSkGf4FGd8B5C2JQdQYPprSQgSvJX",
			Amount:  big.NewInt(1e1),
			TokenID: &tokenID,
		})

		assert.Error(t, err)
		assert.Empty(t, tx)
	})

}

func TestTronClient_Balance(t *testing.T) {

	client, err := tron.NewTronClient(TestGrpcAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	balance, err := client.Balance("TVvxthSkGf4FGd8B5C2JQdQYPprSQgSvJX")

	assert.NoError(t, err)
	assert.NotNil(t, balance)
}

func TestTronClient_TokenBalance(t *testing.T) {

	client, err := tron.NewTronClient(TestGrpcAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	balance, err := client.TokenBalance("TYyxg8cPoR1uWh7aJUqShikruPvnJfv1e9",
		"1002413")

	assert.NoError(t, err)
	assert.NotNil(t, balance)
}

func TestTronClient_AccountBalance(t *testing.T) {

	client, err := tron.NewTronClient(TestGrpcAddress, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	balance, err := client.AccountBalance("TYyxg8cPoR1uWh7aJUqShikruPvnJfv1e9")

	assert.NoError(t, err)
	assert.NotNil(t, balance)
}