package ethecli

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/medeirosfalante/ethcli/abi"
)

//sudo apt-get install libhidapi-dev

type TokenErc20 struct {
	TokenAddress string
	url          string
}

func NewTokenErc20(TokenAddress string, urlClient string) *TokenErc20 {
	return &TokenErc20{
		TokenAddress: TokenAddress,
		url:          urlClient,
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

func (t *TokenErc20) Transfer(address string, amount float64) (string, error) {

	client, err := ethclient.Dial(t.url)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVKEY"))
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress(t.TokenAddress)
	instance, err := abi.NewToken(tokenAddress, client)
	if err != nil {
		return "", err
	}
	total := big.NewFloat(amount)
	value := etherToWei(total)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	addressRef := common.HexToAddress(address)
	tx, err := instance.Transfer(auth, addressRef, value)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil

}
