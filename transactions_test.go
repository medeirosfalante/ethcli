package ethcli_test

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethcli "github.com/medeirosfalante/ethcli"
)

func TestGetByBlock(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	txs := ethcli.NewTransactions(client, "BNB")
	tx, err := txs.GetBlock(9197771, 9197775, []string{"0xf35d75E2Ce765fD4aB1Da7b331eB03C56D4859c4"})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(tx) <= 0 {
		t.Errorf("invalid counter  need more than 0")
		return
	}

}

func TestGetByBlockNative(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	txs := ethcli.NewTransactions(client, "BNB")
	tx, err := txs.GetBlockNative(9337807, 9337809)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(tx) <= 0 {
		t.Errorf("invalid counter  need more than 0")
		return
	}

}

func TestGetByBlockNativeByBlock(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	txs := ethcli.NewTransactions(client, "BNB")

	tx, err := txs.GetBlockNativeByBlock(9337809)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(tx) <= 0 {
		t.Errorf("invalid counter  need more than 0")
		return
	}

}

func TestGetTransactionByHash(t *testing.T) {

	client, err := ethclient.Dial("https://data-seed-prebsc-2-s2.binance.org:8545")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	txs := ethcli.NewTransactions(client, "BNB")

	tx, err := txs.GetTrasactionByHex("0x955a08ca3e25d5a28b40e286417df9bd18e7f59e7ef30f4bc2c24f49c05f5c3f")
	t.Errorf("invalid counter  need more than 0")
	data, _ := json.MarshalIndent(tx, "", "\t")

	t.Errorf("err : \n%s\n", data)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

}
