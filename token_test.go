package ethcli_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethcli "github.com/medeirosfalante/ethcli"

	"github.com/joho/godotenv"
)


func TestSendTokenErc20(t *testing.T) {
	godotenv.Load()

	t.Run("Polygon Network", func(t *testing.T) {
		client, err := ethclient.Dial("https://matic-mumbai.chainstacklabs.com/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}


		token := ethcli.NewTokenErc20("0x2d7882bedcbfddce29ba99965dd3cdf7fcb10a1e", client)
		config := &ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"),
			Path: "0",
			Address: "0x9A034fbc67b2851e9E28F4bb45FD6655E9F9dAeE",
			Amount: 0.05}

		tx, err := token.Transfer(config)
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}
		if tx == "" {
			t.Errorf("tx is empty")
		}
	})

	t.Run("BSC network", func(t *testing.T) {
		client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}

		token := ethcli.NewTokenErc20("0xed24fc36d5ee211ea25a80239fb8c4cfd80f12ee", client)
		config := &ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"),
			Path: "0",
			Address: "0x9A034fbc67b2851e9E28F4bb45FD6655E9F9dAeE",
			Amount: 1}

		tx, err := token.Transfer(config)
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}
		if tx == "" {
			t.Errorf("tx is empty")
		}
	})

}

func TestBuyTokenErc20(t *testing.T) {
	godotenv.Load()
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	token := ethcli.NewDefiToken("0x9e691fd624410d631c082202b050694233031cb7", client)

	tx, err := token.BuyToken(&ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"), Path: "1", Amount: 0.001})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	t.Errorf("err : %s", tx)
	if tx == "" {
		t.Errorf("tx is empty")
	}

}

func TestBalanceTokenErc20(t *testing.T) {
	godotenv.Load()

	t.Run("Polygon Network", func(t *testing.T) {
		client, err := ethclient.Dial("https://matic-mumbai.chainstacklabs.com/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}
		token := ethcli.NewTokenErc20("0x2d7882bedcbfddce29ba99965dd3cdf7fcb10a1e", client)
		balance, err := token.BalanceOf("0x2e0dD29872aee6dA37fe5c9eDf84C0b263AE04DF")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}
		if balance == nil {
			t.Error("balance is nil")
			return
		}

		t.Log(balance.String())
	})

	t.Run("BSC Network", func(t *testing.T) {
		client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}
		token := ethcli.NewTokenErc20("0xed24fc36d5ee211ea25a80239fb8c4cfd80f12ee", client)
		balance, err := token.BalanceOf("0x9A034fbc67b2851e9E28F4bb45FD6655E9F9dAeE")
		if err != nil {
			t.Errorf("err : %s", err)
			return
		}
		if balance == nil {
			t.Error("balance is nil")
			return
		}
		t.Log(balance.String())
	})

}

