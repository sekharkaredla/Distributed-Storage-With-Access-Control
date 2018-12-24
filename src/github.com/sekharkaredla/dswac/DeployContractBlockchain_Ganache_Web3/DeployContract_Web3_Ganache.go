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
	bytecode := "608060405234801561001057600080fd5b506040516020806101e28339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f4d1e838204d8f51c35086fe2ca00346f802334e096d3bf4b4a2d8cc3633278a13382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15060e7806100fb6000396000f3fe6080604052600436106039576000357c010000000000000000000000000000000000000000000000000000000090048063893d20e814603e575b600080fd5b348015604957600080fd5b5060506092565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509056fea165627a7a7230582010eae4d48ad3354bbf646e149a8ccb0707cb6e34b707ca4762fe7e7a00fec80a0029"
	abi := `[{"constant":false,"inputs":[],"name":"getOwner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_JWT","type":"bytes32"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_sender","type":"address"},{"indexed":false,"name":"_JWT","type":"bytes32"}],"name":"StoringJWTEvent","type":"event"}]`
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
	transaction.Gas = big.NewInt(6500000)
	hash, err := contract.Deploy(transaction, bytecode, nil)
	if err != nil {
		fmt.Println("error deploying the contract : ", err)
		os.Exit(-1)
	}
	fmt.Println(hash)
}
