package ethcli_test

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethcli "github.com/medeirosfalante/ethcli"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

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
