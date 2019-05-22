package ganacheBlockchain

import (
	big "math/big"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/complex/types"
	dto "github.com/regcostajr/go-web3/dto"
	providers "github.com/regcostajr/go-web3/providers"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
	storingContract "github.com/sekharkaredla/dswac/MainCode/StoringContract"
)

const GanacheURL = "http://127.0.0.1:7545"

func DeployContract() (string, error) {
	abi := storingContract.StoringDetailsABI
	log.Info.Println("abi being used : ", abi)
	binary := storingContract.StoringDetailsBin
	log.Info.Println("binary of contract : ", binary)
	log.Info.Println("connecting to ganache with URL : ", GanacheURL)
	var connection = web3.NewWeb3(providers.NewHTTPProvider(GanacheURL, 10, false))
	contract, err := connection.Eth.NewContract(abi)
	if err != nil {
		log.Error.Panicln("error creating contract instance, err : ", err)
	}
	transaction := new(dto.TransactionParameters)
	_, err = connection.Eth.ListAccounts()
	if err != nil {
		log.Error.Panicln("error fetching the ethereum accounts, err : ", err)
	}
	log.Info.Println("successfully fetched the accounts")
	coinbase, err := connection.Eth.GetCoinbase()
	transaction.From = coinbase
	transaction.Data = types.ComplexString(binary)
	transaction.Gas = big.NewInt(900000)
	log.Info.Println("using the transaction parameters : ", transaction)
	hash, err := contract.Deploy(transaction, binary, nil)
	if err != nil {
		log.Error.Panicln("error deploying the smart contract, err:", err)
	}
	transactionReceipt, err := connection.Eth.GetTransactionReceipt(hash)
	log.Info.Println("transaction receipt : ", transactionReceipt)
	log.Info.Println("contact address : ", transactionReceipt.ContractAddress)
	return transactionReceipt.ContractAddress, err
}
