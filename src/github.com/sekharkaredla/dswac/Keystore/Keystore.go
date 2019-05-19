package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//	create new keystore
	keyStore := keystore.NewKeyStore("./tmpstore", keystore.StandardScryptN, keystore.StandardScryptP)
	//	create new key (2 ways)
	key1, err := keyStore.NewAccount("password")
	fmt.Println(keyStore, key1, err)
	//	generating a pub pri key
	key2, err := crypto.GenerateKey()
	fmt.Println(key2, err, keyStore)
	// lets see the adding to keystore part also, this below thing can be used if someone knows their private key
	account, err := keyStore.ImportECDSA(key2, "password")
	fmt.Println(account, err)
}
