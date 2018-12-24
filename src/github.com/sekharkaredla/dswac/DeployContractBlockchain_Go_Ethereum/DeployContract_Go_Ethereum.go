package main

import (
	"fmt"
	big "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 100000000)

	address, _, contract, err := DeploySampleContract(
		auth,
		blockchain,
		"Sample Deploy",
	)

	blockchain.Commit()

	if err != nil {
		fmt.Println("Failed to deploy the Inbox contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		fmt.Println("Expected a valid deployment address. Received empty address byte array instead")
	}
	out := fmt.Sprintf("address : %s , contract : %s", address, contract)
	fmt.Println(out)
	if got, err := contract.Message(nil); err == nil {
		fmt.Println("Expected message to be: Sample Deploy. Got : ", got)
	}
}
