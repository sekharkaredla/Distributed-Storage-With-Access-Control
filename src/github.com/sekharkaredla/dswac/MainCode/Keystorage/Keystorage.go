package Keystorage

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
)

type Keystore struct {
	Store *keystore.KeyStore `json:"store"`
}

func CreateNewKeyStore(pathToKeystore string) Keystore {
	var keyStore Keystore
	keyStore.Store = keystore.NewKeyStore(pathToKeystore, keystore.StandardScryptN, keystore.StandardScryptP)
	return keyStore
}

func (keyStore Keystore) GenerateNewAccount(password string) (accounts.Account, error) {
	account, err := keyStore.Store.NewAccount(password)
	if err != nil {
		log.Error.Panicln("Error while creating account")
	}
	return account, nil
}

func (keyStore Keystore) CreateAccountFromPrivateKey(privateKey string, password string) (accounts.Account, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error.Panicln("Error converting the private key")
	}
	account, err := keyStore.Store.ImportECDSA(privateKeyECDSA, password)
	if err != nil {
		log.Error.Panicln("Error while adding to keystore")
	}
	return account, err
}
