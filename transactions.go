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

func NewTransactions(client *ethclient.Client) *Transactions {
	return &Transactions{
		client: client,
	}
}

func (t *Transactions) GetBlock(number int64) (*Block, error) {

	blockNumber := big.NewInt(number)
	block, err := t.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return t.MountBlockData(block)
}

func (t *Transactions) MountBlockData(block *types.Block) (*Block, error) {

	// Build the response to our model
	_block := &Block{
		BlockNumber:       block.Number().Int64(),
		Timestamp:         block.Time(),
		Difficulty:        block.Difficulty().Uint64(),
		Hash:              block.Hash().String(),
		TransactionsCount: len(block.Transactions()),
		Transactions:      []*Transaction{},
	}

	for _, tx := range block.Transactions() {
		txDetail, err := t.ContractCheckDetail(tx, false)
		body, _ := json.Marshal(txDetail)
		log.Printf("body %s", body)
		if err != nil {
			continue
		}
		_block.Transactions = append(_block.Transactions, txDetail)
	}

	return _block, nil

}

func (t *Transactions) GetTxByHash(hash common.Hash) (*Transaction, error) {

	tx, pending, err := t.client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, err
	}

	return t.ContractCheckDetail(tx, pending)
}

func (t *Transactions) ContractCheckDetail(tx *types.Transaction, pending bool) (*Transaction, error) {

	receipt, _ := t.client.TransactionReceipt(context.Background(), tx.Hash())

	txRaw := &Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		To:       tx.To().String(),
		Pending:  pending,
		Nonce:    tx.Nonce(),
	}

	if receipt != nil {
		if len(receipt.Logs) > 0 {
			txRaw.ContractAddress = receipt.Logs[0].Address.String()
			tokenAddress := common.HexToAddress(txRaw.ContractAddress)

			instance, err := abi.NewToken(tokenAddress, t.client)
			if err != nil {
				return nil, err
			}

			tokenTransfer, err := instance.ParseTransfer(*receipt.Logs[0])
			if err != nil {
				return nil, err
			}

			amount := weiToEther(tokenTransfer.Value)

			symbol, err := instance.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			txRaw.Value = amount.String()
			txRaw.Symbol = symbol
			txRaw.Type = "token"
		}
	}

	return txRaw, nil
}
