package ethcli_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethcli "github.com/medeirosfalante/ethcli"

	"github.com/joho/godotenv"
)

func TestSendNative(t *testing.T) {
	godotenv.Load()
	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	native := ethcli.NewNative(client)
	tx, err := native.Transfer(&ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"), Path: "1", Address: "0xB8A688D5A29a35B01CC00d0e2144E01d3c96bFC3", Amount: 350.50})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if tx == "" {
		t.Errorf("tx is empty")
	}

}

func TestBalanceNative(t *testing.T) {
	godotenv.Load()
	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	native := ethcli.NewNative(client)
	balance, err := native.BalanceOf("0xB8A688D5A29a35B01CC00d0e2144E01d3c96bFC3")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if balance == nil {
		t.Error("balance is nil")
		return
	}

}
