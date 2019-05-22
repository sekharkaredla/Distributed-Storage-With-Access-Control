package Keystorage

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	//	"github.com/ethereum/go-ethereum/accounts"
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

func GenerateNewAccount(ganacheURL string, password string) (string, error) {
	account, err := KeyStore.NewAccount(password)
	if err != nil {
		log.Error.Panicln("Error while creating account")
	}
	log.Info.Println("Succesfully created account with password")
	keyJSON, err := KeyStore.Export(account, password, password)
	key, err := keystore.DecryptKey(keyJSON, password)
	log.Info.Println(key)
	privateKeyConverted := hex.EncodeToString(key.PrivateKey.D.Bytes())
	stringConverted := string(privateKeyConverted)
	CreateAccountFromPrivateKey(ganacheURL, stringConverted, password, true)
	return string(keyJSON), err

}

func CreateAccountFromPrivateKey(ganacheURL string, privateKey string, password string, added bool) error {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if !(added) {
		if err != nil {
			log.Error.Panicln("Error converting the private key")
		}
		_, err := KeyStore.ImportECDSA(privateKeyECDSA, password)
		if err != nil {
			log.Error.Panicln("Error while adding to keystore")
		}
	}
	log.Info.Println(privateKeyECDSA)
	params := make([]string, 2)
	params[0] = "0x" + privateKey
	params[1] = password
	log.Info.Println(privateKey)
	log.Info.Println("Succesfully created account with private key and password")
	SendRequestForCreateAccount(ganacheURL, "personal_importRawKey", params)
	return nil
}

func GetKeyStore() *keystore.KeyStore {
	return KeyStore
}

func SendRequestForCreateAccount(ganacheURL string, method string, params interface{}) error {

	bodyString := JSONRPCObject{Version: "2.0", Method: method, Params: params, ID: rand.Intn(100)}

	body := strings.NewReader(bodyString.AsJsonString())
	log.Info.Println(body)
	req, err := http.NewRequest("POST", ganacheURL, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	var client http.Client
	resp, err := client.Do(req)

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
	log.Info.Println(bodyBytes, string(bodyBytes))
	return nil
}
