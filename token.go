package ethcli

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	abiRef "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/medeirosfalante/ethcli/abi"
	"github.com/medeirosfalante/ethcli/util"
)

//sudo apt-get install libhidapi-dev

type TokenErc20 struct {
	wethTokenAddress common.Address
	TokenAddress     string
	client           *ethclient.Client
	instance         *abi.Token
}

type TransferOpts struct {
	Mnemonic string
	Path     string
	Address  string
	Amount   float64
}

const wethTokenAddress = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"

func NewDefiToken(TokenAddress string, client *ethclient.Client) *TokenErc20 {
	tokenAddress := common.HexToAddress(TokenAddress)
	instance, _ := abi.NewToken(tokenAddress, client)
	return &TokenErc20{
		TokenAddress:     TokenAddress,
		client:           client,
		instance:         instance,
		wethTokenAddress: common.HexToAddress(wethTokenAddress),
	}
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

func (t *TokenErc20) BuyToken(req *TransferOpts) (string, error) {
	tokenWantedAddress := common.HexToAddress(t.TokenAddress)
	wethTokenAddressRef := common.HexToAddress(wethTokenAddress)
	pathRouter := []common.Address{wethTokenAddressRef, tokenWantedAddress}

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
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, tx, nil)
		},
	}

	deadline := &big.Int{}
	deadline.SetInt64(time.Now().Add(10 * time.Minute).Unix())

	ref, err := t.instance.Decimals(nil)
	if err != nil {
		return "", err
	}

	amountOutMin := util.ToWei(0.01, int(ref.Int64()))
	auth.Nonce = big.NewInt(int64(nonce))

	file, err := os.OpenFile("./pancakeswap/router.json", os.O_RDWR, 0644)
	if err != nil {
		return "", fmt.Errorf("file %s", err.Error())
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	routerABI, err := abiRef.JSON(strings.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}

	data, err := routerABI.Pack("swapExactETHForTokens", amountOutMin, pathRouter, fromAddress, deadline)
	if err != nil {
		return "", fmt.Errorf("data %s", err.Error())
	}

	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("gasPrice %s", err.Error())
	}
	gasLimit := uint64(91000)

	tx := types.NewTransaction(nonce, fromAddress, big.NewInt(0), gasLimit, gasPrice, data)

	chainID, err := t.client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("chainID %s", err.Error())
	}
	privRef, _ := wallet.PrivateKey(account)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privRef)
	if err != nil {
		return "", fmt.Errorf("signedTx %s", err.Error())
	}
	err = t.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("SendTransaction %s", err.Error())
	}
	return signedTx.Hash().String(), nil

}

func (t *TokenErc20) BalanceOf(account string) (*big.Float, error) {
	tokenAddress := common.HexToAddress(account)
	amount, err := t.instance.BalanceOf(&bind.CallOpts{}, tokenAddress)
	if err != nil {
		return nil, err
	}
	instance, err := abi.NewToken(common.HexToAddress(t.TokenAddress), t.client)
	if err != nil {
		return nil, err
	}
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

	chainID, err := t.client.ChainID(context.Background())
	if err != nil {
		return "", fmt.Errorf("chainID %s", err.Error())
	}

	auth := &bind.TransactOpts{
		From: fromAddress,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != fromAddress {
				return nil, errors.New("not authorized to sign this account")
			}
			// if chainID.Int64() == 80001 {
			// 	return wallet.SignTx(account, tx, chainID)

			// }
			return wallet.SignTxEIP155(account, tx, chainID)
		},
	}

	gasPrice, err := t.GetGasPrice()
	if err != nil {
		return "", fmt.Errorf("GetGasPrice %s", err.Error())
	}
	auth.GasPrice = gasPrice
	tokenAddress := common.HexToAddress(t.TokenAddress)
	instance, err := abi.NewToken(tokenAddress, t.client)
	if err != nil {
		return "", fmt.Errorf("instance %s", err.Error())
	}
	ref, err := instance.Decimals(nil)
	if err != nil {
		return "", err
	}
	value := util.ToWei(req.Amount, int(ref.Int64()))
	auth.Nonce = big.NewInt(int64(nonce))
	addressRef := common.HexToAddress(req.Address)
	tx, err := instance.Transfer(auth, addressRef, value)
	if err != nil {
		return "", fmt.Errorf("tx %s", err.Error())
	}
	return tx.Hash().Hex(), nil

}

func (t *TokenErc20) ChainID() (*big.Int, error) {
	chainID, err := t.client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("chainID %s", err.Error())
	}
	return chainID, nil
}

func (t *TokenErc20) GetGasPrice() (*big.Int, error) {
	return t.client.SuggestGasPrice(context.Background())
}
