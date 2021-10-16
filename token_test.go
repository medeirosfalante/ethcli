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
	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	token := ethcli.NewTokenErc20("0xf35d75E2Ce765fD4aB1Da7b331eB03C56D4859c4", client)

	tx, err := token.Transfer(&ethcli.TransferOpts{Mnemonic: os.Getenv("MNEMONIC"), Path: "1", Address: "0xB8A688D5A29a35B01CC00d0e2144E01d3c96bFC3", Amount: 350.50})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if tx == "" {
		t.Errorf("tx is empty")
	}

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
	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	token := ethcli.NewTokenErc20("0xf35d75E2Ce765fD4aB1Da7b331eB03C56D4859c4", client)
	balance, err := token.BalanceOf("0xB8A688D5A29a35B01CC00d0e2144E01d3c96bFC3")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if balance == nil {
		t.Error("balance is nil")
		return
	}

}
