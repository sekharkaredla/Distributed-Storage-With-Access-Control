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

	address, _, contract, err := DeployStoringDetails(
		auth,
		blockchain,
		//		[]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaWxlSGFzaCI6ImZpbGVfaGFzaCIsImZpbGVOYW1lIjoiZmlsZW5hbWUiLCJvd25lckhhc2giOiJ0ZXN0X293bmVyX2hhc2gifQ.qbNU91JBZy2dXOK5vjsWVDxARILQ1CVyONJLxB5Vl7I"),
		[]byte("old bytes"),
	)
	if err != nil {
		fmt.Println("Failed to deploy the Inbox contract: %v", err)
	}
	if len(address.Bytes()) == 0 {
		fmt.Println("Expected a valid deployment address. Received empty address byte array instead")
	}
	out := fmt.Sprintf("address : %s , contract : %s", address, contract)
	fmt.Println(out)
	blockchain.Commit()
	if got, err := contract.GetJWT(nil); err == nil {
		fmt.Println("current JWT : ", string(got))
	}
	_, err = contract.SetJWT(auth, []byte("new bytes"))
	blockchain.Commit()
	if got, err := contract.GetJWT(nil); err == nil {
		fmt.Println("new JWT : ", string(got))
	}
}
