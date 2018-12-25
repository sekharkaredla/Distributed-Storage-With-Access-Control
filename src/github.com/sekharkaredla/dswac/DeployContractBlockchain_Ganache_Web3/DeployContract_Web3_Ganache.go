package main

import (
	"fmt"
	big "math/big"
	"os"

	web3 "github.com/regcostajr/go-web3"
	dto "github.com/regcostajr/go-web3/dto"
	providers "github.com/regcostajr/go-web3/providers"
)

func main() {
	bytecode := "608060405234801561001057600080fd5b506040516020806101028339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000819055505060ab806100576000396000f3fe6080604052600436106039576000357c010000000000000000000000000000000000000000000000000000000090048063b32fead514603e575b600080fd5b348015604957600080fd5b50607360048036036020811015605e57600080fd5b81019080803590602001909291905050506075565b005b806000819055505056fea165627a7a72305820d84d31b418699570f617fe333ac27f96127f81b5ec888d5fa891bd924face88d0029"
	abi := `[{"constant":false,"inputs":[{"name":"newMessage","type":"bytes32"}],"name":"setMessage","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"initialMessage","type":"bytes32"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`
	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))
	contract, err := connection.Eth.NewContract(abi)
	if err != nil {
		fmt.Errorf("error generating contract with ABI")
	}
	transaction := new(dto.TransactionParameters)
	coinbase, err := connection.Eth.GetCoinbase()
	if err != nil {
		fmt.Errorf("error getting the coinbase")
	}
	fmt.Println("Coinbase : ", coinbase)
	transaction.From = coinbase
	transaction.Gas = big.NewInt(6721974)
	fmt.Println(transaction)
	hash, err := contract.Deploy(transaction, bytecode, nil)
	fmt.Println(hash,err)
	if err != nil {
		fmt.Println("error deploying the contract : ", err)
		os.Exit(-1)
	}
	fmt.Println(hash)
}
