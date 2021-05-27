package ethcli

import (
	"context"
	"encoding/json"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/medeirosfalante/ethcli/abi"
)

type Block struct {
	BlockNumber       int64          `json:"blockNumber"`
	Timestamp         uint64         `json:"timestamp"`
	Difficulty        uint64         `json:"difficulty"`
	Hash              string         `json:"hash"`
	TransactionsCount int            `json:"transactionsCount"`
	Transactions      []*Transaction `json:"transactions"`
}

// Transaction data structure
type Transaction struct {
	Hash            string `json:"hash"`
	Value           string `json:"value"`
	Gas             uint64 `json:"gas"`
	GasPrice        uint64 `json:"gasPrice"`
	Nonce           uint64 `json:"nonce"`
	To              string `json:"to"`
	Pending         bool   `json:"pending"`
	Input           bool   `json:"input"`
	ContractAddress string `json:"contract_address"`
	Type            string `json:"type"`
	Symbol          string `json:"symbol"`
}

// TransferEthRequest data structure
type TransferEthRequest struct {
	PrivKey string `json:"privKey"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

// HashResponse data structure
type HashResponse struct {
	Hash string `json:"hash"`
}

// BalanceResponse data structure
type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Symbol  string `json:"symbol"`
	Units   string `json:"units"`
}

// Error data structure
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}

type Transactions struct {
	client *ethclient.Client
}
