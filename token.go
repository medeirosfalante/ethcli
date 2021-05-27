package ethcli

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

func (t *TokenErc20) Transfer(address string, amount float64) (string, error) {

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
	nonce, err := t.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	log.Printf("fromAddress \n%s\n", fromAddress)

	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	chainID, err := t.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress(t.TokenAddress)
	instance, err := abi.NewToken(tokenAddress, t.client)
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
