package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const GanacheURL = "127.0.0.1:7545"

func main() {
	client, err := ethclient.Dial("http://" + GanacheURL)
	if err != nil {
		log.Fatal(err)
	}
	//Here give the private key from the accounts in ganache
	privateKey, err := crypto.HexToECDSA("6d4371b61608bc2fc68c46a238e3906e0cca4c8f7958129d57ba3cb90923ac06")
	if err != nil {
		log.Fatal(err)
	}

	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}

	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	auth := bind.NewKeyedTransactor(privateKey)
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice

	//input := "1.0"
	_, _, SDInstance, err := DeployStoringDetails(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	storingClientCaller := SDInstance.StoringDetailsCaller
	storingClientTransactor := SDInstance.StoringDetailsTransactor
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(contractAddressHex, err)
	_, err = storingClientTransactor.SetJWT(auth, []byte("you got this one"))
	data, err := storingClientCaller.GetJWT(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data from getJwt : ", string(data))
}
