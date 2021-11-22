package tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"math/big"
)

type TransferOpts struct{
	From string
	PrivKey *ecdsa.PrivateKey
	To string
	Amount *big.Int
	TokenID *string
	FeeLimit *big.Int
}

type Client struct {
	url string
	conn *client.GrpcClient
}

func NewTronClient(url string, opts ...grpc.DialOption) (*Client, error) {
	conn := client.NewGrpcClient(url)
	err := conn.Start(opts...)
	if err != nil{
		return nil, err
	}
	return &Client{url: url, conn: conn}, nil
}

func (t Client) SetApiKey(apikey string) error{
	return t.conn.SetAPIKey(apikey)
}

func (t Client) Account(address string) (*core.Account, error){
	acc, err := t.conn.GetAccount(address)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (t Client) Balance(address string) (int64, error){
	acc, err := t.conn.GetAccount(address)
	if err != nil {
		return 0, err
	}
	return acc.GetBalance(), nil
}

func (t Client) IsContract(address string) (bool, error) {
	result, err := t.conn.GetAccount(address)
	if err != nil {
		return false, err
	}

	if result.GetType().String() == "Contract" {
		return true, nil
	}

	return false, nil
}

func (t Client) IsTRC10(tokenID string) bool{
	result, err := t.conn.GetAssetIssueByID(tokenID)

	if err != nil {
		log.Fatalf(err.Error())
		return false
	}

	if result.Id == ""{
		return false
	}

	return true
}

func (t *Client) Transfer(opts *TransferOpts) (string, error){
	var tx *api.TransactionExtention
	var err error

	if opts.TokenID != nil{
		if t.IsTRC10(*opts.TokenID) {
			tx, err = t.conn.TransferAsset(opts.From, opts.To, *opts.TokenID, opts.Amount.Int64())

			if err != nil {
				return "", err
			}
		} else {
			var feeLimit *big.Int
			if opts.FeeLimit == nil{
				feeLimit = big.NewInt(4e7)
			}
			tx, err = t.conn.TRC20Send(opts.From, opts.To, *opts.TokenID, opts.Amount, feeLimit.Int64())

			if err != nil {
				return "", err
			}
		}
	} else {
		tx, err = t.conn.Transfer(opts.From, opts.To, opts.Amount.Int64())
		if err != nil {
			return "", err
		}
	}

	signTx, err := t.signTx(opts.PrivKey, tx.GetTransaction())
	if err != nil {
		return "", err
	}

	err = t.sendTransaction(signTx)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(tx.GetTxid()), nil
}

func (t *Client) signTx(key *ecdsa.PrivateKey, tx *core.Transaction) (*core.Transaction, error)  {

	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	txHashBytes := h256h.Sum(nil)

	signature, err := crypto.Sign(txHashBytes, key)
	if err != nil {
		return nil, err
	}

	tx.Signature = append(tx.Signature, signature)

	return tx, err
}

func (t *Client) sendTransaction(tx *core.Transaction) error {

	result, err := t.conn.Broadcast(tx)

	if err != nil {
		return fmt.Errorf("err: %s", err.Error())
	}

	if result.Code != 0 {
		return fmt.Errorf("bad transaction: %v", string(result.GetMessage()))
	}

	return nil
}


