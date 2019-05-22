package Keystorage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
)

type HTTPProvider struct {
	address string
	timeout int32
	secure  bool
	client  *http.Client
}

type JSONRPCObject struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

func (jrpc *JSONRPCObject) AsJsonString() string {
	resultBytes, err := json.Marshal(jrpc)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(resultBytes)
}

var KeyStore *keystore.KeyStore

func init() {
	CreateNewKeyStore("./tmpstore")
}

func CreateNewKeyStore(pathToKeystore string) error {
	KeyStore = keystore.NewKeyStore(pathToKeystore, keystore.StandardScryptN, keystore.StandardScryptP)
	log.Info.Println("Successfully created keystore")
	return nil
}

func GenerateNewAccount(password string) (string, error) {
	account, err := KeyStore.NewAccount(password)
	if err != nil {
		log.Error.Panicln("Error while creating account")
	}
	log.Info.Println("Succesfully created account with password")
	keyJSON, err := KeyStore.Export(account, password, password)
	return string(keyJSON), err
}

func CreateAccountFromPrivateKey(privateKey string, password string) (accounts.Account, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error.Panicln("Error converting the private key")
	}
	account, err := KeyStore.ImportECDSA(privateKeyECDSA, password)
	if err != nil {
		log.Error.Panicln("Error while adding to keystore")
	}
	log.Info.Println("Succesfully created account with private key and password")
	return account, err
}

func GetKeyStore() *keystore.KeyStore {
	return KeyStore
}

func SendRequest(ganacheURL string, v interface{}, method string, params interface{}) error {

	bodyString := util.JSONRPCObject{Version: "2.0", Method: method, Params: params, ID: rand.Intn(100)}

	body := strings.NewReader(bodyString.AsJsonString())
	req, err := http.NewRequest("POST", ganacheURL, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := provider.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var bodyBytes []byte

	if resp.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bodyBytes, v)

}
