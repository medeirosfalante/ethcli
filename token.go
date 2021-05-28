package ethcli

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/medeirosfalante/ethcli/abi"
)

//sudo apt-get install libhidapi-dev

type TokenErc20 struct {
	TokenAddress string
	client       *ethclient.Client
}

func NewTokenErc20(TokenAddress string, client *ethclient.Client) *TokenErc20 {
	return &TokenErc20{
		TokenAddress: TokenAddress,
		client:       client,
	}
}

func etherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func (t *TokenErc20) Transfer(mnemonic string, address string, amount float64) (string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return "", err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		return "", fmt.Errorf("account %s", err.Error())
	}

	fromAddress := account.Address
	nonce, err := t.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("nonce %s", err.Error())
	}

	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("gasPrice %s", err.Error())
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
	total := big.NewFloat(amount)
	value := etherToWei(total)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	addressRef := common.HexToAddress(address)
	tx, err := instance.Transfer(auth, addressRef, value)
	if err != nil {
		return "", fmt.Errorf("tx %s", err.Error())
	}
	return tx.Hash().Hex(), nil

}
