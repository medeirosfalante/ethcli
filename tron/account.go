package tron

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

type TronAccount struct {
	Mnemonic   string
	Path       string
	privKey *ecdsa.PrivateKey
	pubKey string
}

func NewTronAccount(mnemonic string, path string) (TronAccount, error) {

	key, err := PrivateKey(mnemonic, path)
	if err != nil {
		return TronAccount{}, err
	}

	pkey, err := PublicKey(mnemonic, path)

	if err != nil {
		return TronAccount{}, err
	}

	return TronAccount{
		Mnemonic: mnemonic,
		Path:     path,
		privKey:  key,
		pubKey:   pkey,
	}, nil
}

func (a TronAccount) PubKeyFromPath(path string) (string, error){

	wallet, err := hdwallet.NewFromMnemonic(a.Mnemonic)
	if err != nil {
		return "", fmt.Errorf("mnemonic %s", err.Error())
	}

	p := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(p, false)
	if err != nil {
		return "", fmt.Errorf("account %s", err.Error())
	}

	publicKeyECDSA, err := wallet.PublicKey(account)
	if err != nil {
		return "", fmt.Errorf("pkey %s", err.Error())
	}

	addr := address.PubkeyToAddress(*publicKeyECDSA)

	a.pubKey = addr.String()

	return addr.String(), nil

}

func (a TronAccount) PublicKey() string{
	return a.pubKey
}

func (a TronAccount) PrivateKey() *ecdsa.PrivateKey {
	return a.privKey
}



func PrivateKey(mnemonic, path string) (*ecdsa.PrivateKey, error){
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	p := hdwallet.MustParseDerivationPath(fmt.Sprintf(path))
	account, err := wallet.Derive(p, false)
	if err != nil {
		return nil, err
	}

	privateKeyECDSA, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, err
	}

	return privateKeyECDSA, nil
}

func PublicKey(mnemonic, path string) (string, error){

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return "", fmt.Errorf("mnemonic %s", err.Error())
	}

	p := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(p, false)
	if err != nil {
		return "", fmt.Errorf("account %s", err.Error())
	}

	publicKeyECDSA, err := wallet.PublicKey(account)
	if err != nil {
		return "", fmt.Errorf("pkey %s", err.Error())
	}

	addr := address.PubkeyToAddress(*publicKeyECDSA)

	log.Println(addr.String())

	return addr.String(), nil
}