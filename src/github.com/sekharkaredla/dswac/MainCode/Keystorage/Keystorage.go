package Keystorage

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
)

var KeyStore *keystore.KeyStore

func init() {
	CreateNewKeyStore("./tmpstore")
}

func CreateNewKeyStore(pathToKeystore string) error {
	KeyStore = keystore.NewKeyStore(pathToKeystore, keystore.StandardScryptN, keystore.StandardScryptP)
	log.Info.Println("Successfully created keystore")
	return nil
}

func GenerateNewAccount(password string) (string, error) {
	account, err := KeyStore.NewAccount(password)
	if err != nil {
		log.Error.Panicln("Error while creating account")
	}
	log.Info.Println("Succesfully created account with password")
	keyJSON, err := KeyStore.Export(account, password, password)
	return string(keyJSON), err
}

func CreateAccountFromPrivateKey(privateKey string, password string) (accounts.Account, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error.Panicln("Error converting the private key")
	}
	account, err := KeyStore.ImportECDSA(privateKeyECDSA, password)
	if err != nil {
		log.Error.Panicln("Error while adding to keystore")
	}
	log.Info.Println("Succesfully created account with private key and password")
	return account, err
}

func GetKeyStore() *keystore.KeyStore {
	return KeyStore
}
