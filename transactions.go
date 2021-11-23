package ethcli

import (
	"context"
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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
	Hash            string  `json:"hash"`
	Value           string  `json:"value"`
	Gas             uint64  `json:"gas"`
	GasPrice        uint64  `json:"gasPrice"`
	Nonce           uint64  `json:"nonce"`
	To              string  `json:"to"`
	From            string  `json:"from"`
	Pending         bool    `json:"pending"`
	Input           bool    `json:"input"`
	ContractAddress string  `json:"contract_address"`
	Type            string  `json:"type"`
	Symbol          string  `json:"symbol"`
	Confirmation    int64   `json:"confirmation"`
	ValueFormated   float64 `json:"value_formated"`
	Blockhash       string  `json:"blockhash"`
	BlockIndex      int64   `json:"blockindex"`
	Blocktime       int64   `json:"blocktime"`
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
	client     *ethclient.Client
	NativeName string
}

func NewTransactions(client *ethclient.Client, NativeName string) *Transactions {
	return &Transactions{
		client:     client,
		NativeName: NativeName,
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

	if tx.To() == nil {
		return nil, errors.New("to is null")
	}
	txRaw := &Transaction{
		Hash:         txHash.String(),
		Value:        tx.Value().String(),
		To:           tx.To().String(),
		Pending:      isPending,
		Nonce:        tx.Nonce(),
		Confirmation: 1,
		Blockhash:    log.BlockHash.Hex(),
		BlockIndex:   int64(log.BlockNumber),
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
	if err != nil {
		return nil, err
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	ref, err := instance.Decimals(nil)
	if err != nil {
		return nil, err
	}

	amount := weiToEther(tokenTransfer.Value, math.Pow10(int(ref.Int64())))
	txRaw.Value = amount.String()
	txRaw.Symbol = symbol
	txRaw.Type = "token"

	return txRaw, nil
}

func (t *Transactions) GetBlockNative(from int64, to int64) ([]*Transaction, error) {
	var transactions []*Transaction
	for i := from; i < to; i++ {
		block, err := t.client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			continue
		}
		transactionsIn, _ := t.ProcessTransations(block.Transactions())
		transactions = append(transactions, transactionsIn...)
	}

	return transactions, nil
}

func (t *Transactions) GetBlockNativeByBlock(blockNumber int64) ([]*Transaction, error) {
	var transactions []*Transaction
	block, err := t.client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}
	transactions, _ = t.ProcessTransations(block.Transactions())
	return transactions, nil
}

func (t *Transactions) ProcessTransations(txs types.Transactions) ([]*Transaction, error) {
	var transactions []*Transaction
	for _, tx := range txs {
		if tx.Value().Int64() > 0 {
			var to string
			if tx.To() != nil {
				to = tx.To().String()
			}

			ValueFormated, _ := weiToEther(tx.Value(), params.Ether).Float64()
			txRaw := &Transaction{
				Hash:          tx.Hash().String(),
				Value:         tx.Value().String(),
				Gas:           tx.Gas(),
				GasPrice:      tx.GasPrice().Uint64(),
				To:            to,
				Nonce:         tx.Nonce(),
				Confirmation:  1,
				ValueFormated: ValueFormated,
				Symbol:        t.NativeName,
			}
			if index := t.FindTransactionByID(txRaw.Hash, transactions); index > 0 {
				transactions[index] = txRaw
			} else {
				transactions = append(transactions, txRaw)
			}

		}

	}

	return transactions, nil
}

func (t *Transactions) GetTrasactionByHex(hash string) (*Transaction, error) {
	header, err := t.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	tx, _, err := t.client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, err
	}
	receipt, err := t.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, err
	}

	confirmations := header.Number.Int64() - receipt.BlockNumber.Int64()

	ValueFormated, _ := weiToEther(tx.Value(), params.Ether).Float64()

	var to string
	if tx.To() != nil {
		to = tx.To().String()
	}

	chainID, err := t.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
	if err != nil {
		return nil, err
	}

	txRaw := &Transaction{
		Hash:          tx.Hash().String(),
		Value:         tx.Value().String(),
		Gas:           tx.Gas(),
		GasPrice:      tx.GasPrice().Uint64(),
		To:            to,
		Nonce:         tx.Nonce(),
		Confirmation:  int64(confirmations),
		ValueFormated: ValueFormated,
		Symbol:        t.NativeName,
		From:          msg.From().Hex(),
		Blockhash:     receipt.BlockHash.Hex(),
		BlockIndex:    int64(receipt.BlockNumber.Uint64()),
	}

	return txRaw, nil
}

func (t *Transactions) FindTransactionByID(hash string, txs []*Transaction) int {
	for index, item := range txs {
		if item.Hash == hash {
			return index
		}
	}

	return -1

}
