package ethcli_test

import (
	"encoding/json"
	"log"
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
	txs := ethcli.NewTransactions(client)

	tx, err := txs.GetBlock(9197775)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	body, _ := json.Marshal(tx)
	log.Printf("body %s", body)

	t.Error("block ")

}
