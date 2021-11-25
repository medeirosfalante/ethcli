package tron

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/gogo/protobuf/proto"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"
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
	http *http.Client
}

func NewTronClient(url string, opts ...grpc.DialOption) (*Client, error) {
	conn := client.NewGrpcClient(url)
	err := conn.Start(opts...)
	if err != nil{
		return nil, err
	}
	return &Client{url: url, conn: conn, http: &http.Client{Timeout: time.Second * 10}}, nil
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

//	TokenBalance only work with tron mainnet
//	TokenID can also be contract address in case of TRC20
func (t *Client) TokenBalance(address, tokenID string) (*big.Float, error) {

	body, err := json.Marshal(&struct {
		Address string `json:"address,omitempty"`
	}{
		Address: address,
	})

	if err != nil {
		return nil, fmt.Errorf("address is a invalid type")
	}

	req, err := http.NewRequest(http.MethodPost,
		"https://apilist.tronscan.org/api/account", bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("cannot request: %s", err.Error())
	}

	resp, err := t.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot request: %s", err.Error())
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %s", err.Error())
	}

	var account Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, fmt.Errorf("cannot parse information: %s", err.Error())
	}

	for _, trc20 := range account.TRC20 {
		if trc20.TokenID == tokenID {
			amount, _ := decimal.NewFromString(trc20.Balance)
			return amount.BigFloat(), nil
		}
	}

	for _, trc10 := range account.TRC10 {
		if trc10.TokenID == tokenID {
			amount, _ := decimal.NewFromString(trc10.Balance)
			return amount.BigFloat(), nil
		}
	}

	for _, balance := range account.Balances {
		if balance.TokenID == tokenID {
			amount, _ := decimal.NewFromString(balance.Balance)
			return amount.BigFloat(), nil
		}
	}

	return new(big.Float).SetFloat64(0), nil
}

func (t *Client) AccountBalance(address string) (*Account, error){

	body, err := json.Marshal(&struct {
		Address string `json:"address,omitempty"`
	}{
		Address: address,
	})

	if err != nil {
		return nil, fmt.Errorf("address is a invalid type")
	}

	req, err := http.NewRequest(http.MethodPost,
		"https://apilist.tronscan.org/api/account", bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("cannot request: %s", err.Error())
	}

	resp, err := t.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot request: %s", err.Error())
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %s", err.Error())
	}

	var account Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, fmt.Errorf("cannot parse information: %s", err.Error())
	}

	return &account, nil

}
