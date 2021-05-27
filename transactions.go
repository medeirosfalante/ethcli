package ethcli

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	From            string `json:"from"`
	Pending         bool   `json:"pending"`
	Input           bool   `json:"input"`
	ContractAddress string `json:"contract_address"`
	Type            string `json:"type"`
	Symbol          string `json:"symbol"`
	Confirmation    int64  `json:"confirmation"`
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

func (t *Transactions) GetBlock(from int64, to int64, contractsAddresses []string) ([]*Transaction, error) {
	list := []common.Address{}
	for _, item := range contractsAddresses {
		contractAddress := common.HexToAddress(item)
		list = append(list, contractAddress)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(from),
		ToBlock:   big.NewInt(to),
		Addresses: list,
	}

	logs, err := t.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			txDetail, err := t.ContractCheckDetail(vLog, false)
			if err != nil {
				continue
			}
			transactions = append(transactions, txDetail)
		}
	}
	return transactions, nil
}

func (t *Transactions) ContractCheckDetail(log types.Log, pending bool) (*Transaction, error) {
	txHash := common.HexToHash(log.TxHash.Hex())

	tx, isPending, err := t.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, err
	}

	header, err := t.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	confirmations := header.Number.Uint64() - log.BlockNumber

	txRaw := &Transaction{
		Hash:         tx.Hash().String(),
		Value:        tx.Value().String(),
		Gas:          tx.Gas(),
		GasPrice:     tx.GasPrice().Uint64(),
		To:           tx.To().String(),
		Pending:      isPending,
		Nonce:        tx.Nonce(),
		Confirmation: int64(confirmations),
	}

	txRaw.From = common.HexToAddress(log.Topics[1].Hex()).String()
	txRaw.To = common.HexToAddress(log.Topics[2].Hex()).String()

	txRaw.ContractAddress = log.Address.String()
	tokenAddress := common.HexToAddress(txRaw.ContractAddress)

	instance, err := abi.NewToken(tokenAddress, t.client)
	if err != nil {
		return nil, err
	}

	tokenTransfer, err := instance.ParseTransfer(log)

	amount := weiToEther(tokenTransfer.Value)
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	txRaw.Value = amount.String()
	txRaw.Symbol = symbol
	txRaw.Type = "token"

	return txRaw, nil
}
