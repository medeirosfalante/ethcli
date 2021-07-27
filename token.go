package ethcli

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/medeirosfalante/ethcli/abi"
)

//sudo apt-get install libhidapi-dev

type TokenErc20 struct {
	TokenAddress string
	client       *ethclient.Client
	instance     *abi.Token
}

type TransferOpts struct {
	Mnemonic string
	Path     string
	Address  string
	Amount   float64
}

func NewTokenErc20(TokenAddress string, client *ethclient.Client) *TokenErc20 {
	tokenAddress := common.HexToAddress(TokenAddress)
	instance, _ := abi.NewToken(tokenAddress, client)
	return &TokenErc20{
		TokenAddress: TokenAddress,
		client:       client,
		instance:     instance,
	}
}

func etherToWei(eth *big.Float, parm int) *big.Int {
	truncInt, _ := eth.Int(nil)
	potencial := strings.Repeat("0", parm)
	i, _ := strconv.ParseInt("1"+potencial, 10, 64)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(i))
	fracInt, _ := new(big.Int).SetString(potencial, parm)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func weiToEther(wei *big.Int, parm float64) *big.Float {
	f := new(big.Float)
	f.SetPrec(236)
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236)
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(parm))
}

func (t *TokenErc20) BalanceOf(account string) (*big.Float, error) {
	tokenAddress := common.HexToAddress(account)
	amount, err := t.instance.BalanceOf(&bind.CallOpts{}, tokenAddress)
	if err != nil {
		return nil, err
	}
	instance, err := abi.NewToken(tokenAddress, t.client)
	ref, err := instance.Decimals(nil)
	if err != nil {
		return nil, err
	}

	return weiToEther(amount, math.Pow10(int(ref.Int64()))), nil
}

func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

func (t *TokenErc20) Transfer(req *TransferOpts) (string, error) {
	wallet, err := hdwallet.NewFromMnemonic(req.Mnemonic)
	if err != nil {
		return "", err
	}

	path := hdwallet.MustParseDerivationPath(req.Path)
	account, err := wallet.Derive(path, true)
	if err != nil {
		return "", fmt.Errorf("account %s", err.Error())
	}

	fromAddress := account.Address
	nonce, err := t.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("nonce %s", err.Error())
	}

	auth := &bind.TransactOpts{
		From: fromAddress,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != fromAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			return wallet.SignTx(account, tx, nil)
		},
	}

	tokenAddress := common.HexToAddress(t.TokenAddress)
	instance, err := abi.NewToken(tokenAddress, t.client)
	if err != nil {
		return "", fmt.Errorf("instance %s", err.Error())
	}
	total := big.NewFloat(req.Amount)
	ref, err := instance.Decimals(nil)
	if err != nil {
		return "", err
	}
	value := etherToWei(total, int(ref.Int64()))
	auth.Nonce = big.NewInt(int64(nonce))
	addressRef := common.HexToAddress(req.Address)
	tx, err := instance.Transfer(auth, addressRef, value)
	if err != nil {
		return "", fmt.Errorf("tx %s", err.Error())
	}
	return tx.Hash().Hex(), nil

}
