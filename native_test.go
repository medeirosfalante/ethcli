package ethcli_test

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethcli "github.com/medeirosfalante/ethcli"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/joho/godotenv"
)

func TestSendNative(t *testing.T) {
	godotenv.Load()

	t.Run("Polygon Native", func(t *testing.T) {
		client, err := ethclient.Dial("https://matic-mumbai.chainstacklabs.com/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}


		native := ethcli.NewNative(client)

		config := &ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"),
			Path: "0",
			Address: "0x9A034fbc67b2851e9E28F4bb45FD6655E9F9dAeE",
			Amount: 0.005}


		tx, err := native.Transfer(config)
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}

		if tx == "" {
			t.Errorf("tx is empty")
		}
	})


	t.Run("BSC Native", func(t *testing.T) {
		client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}

		native := ethcli.NewNative(client)


		config := &ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"),
			Path: "0",
			Address: "0x9A034fbc67b2851e9E28F4bb45FD6655E9F9dAeE",
			Amount: 0.1}


		tx, err := native.Transfer(config)
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}

		if tx == "" {
			t.Errorf("tx is empty")
		}
	})

}

func TestBalanceNative(t *testing.T) {
	godotenv.Load()
	client, err := ethclient.Dial("https://matic-mumbai.chainstacklabs.com/")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	native := ethcli.NewNative(client)
	balance, err := native.BalanceOf("0xf48CEE37b394B3dc04B4E16C1Ed9B12DCfac531E")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}


	if balance == nil {
		t.Error("balance is nil")
		return
	}


}

func TestGetAddress(t *testing.T) {
	godotenv.Load()
	wallet, err := hdwallet.NewFromMnemonic(os.Getenv("MNEMONIC"))
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	path := hdwallet.MustParseDerivationPath("1")
	account, err := wallet.Derive(path, true)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	priv, _ := wallet.PrivateKeyHex(account)
	t.Errorf("err : %s\n", priv)
	t.Errorf("err : %s\n", account.Address)
	if account.Address.Hex() != "" {
		t.Errorf("err : %s", err)
		return
	}
}

func TestNative_ChainID(t *testing.T) {

	t.Run("Polygon Network", func(t *testing.T) {
		client, err := ethclient.Dial("https://matic-mumbai.chainstacklabs.com/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			t.Errorf("chainID %s", err.Error())
		}

		t.Logf(chainID.String())
	})

	t.Run("BSC Network", func(t *testing.T) {
		client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			t.Errorf("chainID %s", err.Error())
		}

		t.Logf(chainID.String())
	})

}