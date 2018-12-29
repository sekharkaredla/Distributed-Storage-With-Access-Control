package main

import (
	"fmt"
	"log"
	big "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	//		"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	//	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/complex/types"
	dto "github.com/regcostajr/go-web3/dto"
	providers "github.com/regcostajr/go-web3/providers"
)

const GanacheURL = "127.0.0.1:7545"

func DeployContract() (string, error) {
	abi := StoringDetailsABI
	binary := StoringDetailsBin
	var connection = web3.NewWeb3(providers.NewHTTPProvider(GanacheURL, 10, false))
	contract, err := connection.Eth.NewContract(abi)
	if err != nil {
		fmt.Errorf("error generating contract with ABI")
	}
	transaction := new(dto.TransactionParameters)
	Accounts, err := connection.Eth.ListAccounts()
	if err != nil {
		fmt.Errorf("error getting the accounts")
	}
	coinbase, err := connection.Eth.GetCoinbase()
	// fmt.Println("Accounts : ", Accounts)
	transaction.From = coinbase
	// transaction.To = coinbase
	transaction.From = Accounts[0]
	transaction.Data = types.ComplexString(binary)
	transaction.Gas = big.NewInt(900000)
	// fmt.Println(transaction)
	hash, err := contract.Deploy(transaction, binary, nil)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(hash)
	transactionReceipt, err := connection.Eth.GetTransactionReceipt(hash)
	fmt.Println("contact address : ", transactionReceipt.ContractAddress)
	return transactionReceipt.ContractAddress, err
}

func main() {

	client, err := ethclient.Dial("http://" + GanacheURL)
	//	key, _ := crypto.GenerateKey()
	//	auth := bind.NewKeyedTransactor(key)
	if err != nil {
		log.Fatal(err)
	}
	priv := "6d4371b61608bc2fc68c46a238e3906e0cca4c8f7958129d57ba3cb90923ac06"
	key, err := crypto.HexToECDSA(priv)
	contractAddress, err := DeployContract()
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(key)
	// var nonce int64 = 0
	// auth.Nonce = big.NewInt(nonce)
	// fmt.Println(contractAddress)
	contractAddressHex := common.HexToAddress(contractAddress[2:])
	storingClient, err := NewStoringDetails(contractAddressHex, client)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(contractAddressHex, err)
	_, err = storingClient.SetJWT(auth, []byte("test"))
	data, err := storingClient.GetJWT(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data from getJwt : ", string(data))
}
