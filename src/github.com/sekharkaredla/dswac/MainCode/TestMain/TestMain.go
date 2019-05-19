package main

import (
	"encoding/hex"
	"fmt"

	crypto "github.com/ethereum/go-ethereum/crypto"
	keystorage "github.com/sekharkaredla/dswac/MainCode/Keystorage"
)

func main() {
	keyStore := keystorage.CreateNewKeyStore("./tmpstore")
	fmt.Println(keyStore.GenerateNewAccount("pass"))
	key, _ := crypto.GenerateKey()
	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Println(keyStore.CreateAccountFromPrivateKey(privateKey, "pass"))
	fmt.Println(keyStore.Store.Accounts())
}
