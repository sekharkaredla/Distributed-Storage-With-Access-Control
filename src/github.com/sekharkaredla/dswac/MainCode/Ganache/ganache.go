package ganacheBlockchain

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ipfs "github.com/sekharkaredla/dswac/MainCode/IPFS"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
	storingContract "github.com/sekharkaredla/dswac/MainCode/StoringContract"
)

const GanacheURL = "http://127.0.0.1:7545"

var SDInstance *storingContract.StoringDetails

func DeployContract() error {
	client, err := ethclient.Dial(GanacheURL)
	if err != nil {
		log.Error.Fatalln(err)
	}
	//Here give the private key from the accounts in ganache
	privateKey, err := crypto.HexToECDSA("7cc50abd61846f182eb5fa2737751a065da74ee5323c94f265b09a477fb45c49")
	if err != nil {
		log.Error.Fatalln(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	_, _, sdInstance, err := storingContract.DeployStoringDetails(auth, client)
	if err != nil {
		log.Error.Fatalln(err)
	}
	SDInstance = sdInstance
	return err
}
func AddFileToGanache(filepath string) (string, error) {
	cid, err := ipfs.AddFileToIPFS(filepath)
	if err != nil {
		log.Error.Panicln(err)
	}
	log.Info.Println("successfully added to IPFS ", cid)
	userPublicKey := "0x4e15871aFA224d13311A30822613Bb7488166DB2"
	userPrivateKey := "7cc50abd61846f182eb5fa2737751a065da74ee5323c94f265b09a477fb45c49"
	var mySiginingKey = []byte(userPrivateKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ownerHash"] = userPublicKey
	claims["fileName"] = filepath
	claims["fileHash"] = cid
	tokenString, err := token.SignedString(mySiginingKey)
	if err != nil {
		log.Error.Panicln(err)
	}
	privateKeyECDSA, _ := crypto.HexToECDSA(userPrivateKey)
	auth := bind.NewKeyedTransactor(privateKeyECDSA)
	transaction, err := SDInstance.StoringDetailsTransactor.SetJWT(auth, []byte(tokenString))
	if err != nil {
		log.Error.Panicln(err)
	}
	log.Info.Println("successfully added to blockchain ", transaction)
	return tokenString, err
}
