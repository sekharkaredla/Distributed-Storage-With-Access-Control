package main

import (
	"encoding/hex"
	"fmt"

	crypto "github.com/ethereum/go-ethereum/crypto"
	keystorage "github.com/sekharkaredla/dswac/MainCode/Keystorage"
)

func main() {
	fmt.Println(keystorage.GenerateNewAccount("pass"))
	key, _ := crypto.GenerateKey()
	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Println(keystorage.CreateAccountFromPrivateKey(privateKey, "pass"))
	fmt.Println(keystorage.GetKeyStore().Accounts())
}
