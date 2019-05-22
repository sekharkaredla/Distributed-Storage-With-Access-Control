package main

import (
	"fmt"

	web3 "github.com/regcostajr/go-web3"
	providers "github.com/regcostajr/go-web3/providers"
)

const GanacheURL = "127.0.0.1:7545"

func CreateNewUserInBlockchain() {
	var connection = web3.NewWeb3(providers.NewHTTPProvider(GanacheURL, 10, false))
	//	a, b := connection.Personal.NewAccount("abcd")
	//	fmt.Println(a, b)
	strings, _ := connection.Eth.ListAccounts()
	fmt.Println(len(strings), strings)
}

func main() {
	CreateNewUserInBlockchain()
}
