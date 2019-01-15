package CreateNewUser

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
)

type User struct {
	UserName   string `json:"username"`
	PublicKey  string `json:"publickey,omitempty"`
	PrivateKey string `json:"privatekey,omitempty"`
}

func CreateNewUser(username string) (User, error) {
	log.Info.Println("Creating user : ", username)
	key, err := crypto.GenerateKey()
	var user User
	user.UserName = username
	if err != nil {
		log.Error.Println("error generating key, err : ", err)
		return user, err
	}
	publicKey := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())
	user.PublicKey = publicKey
	user.PrivateKey = privateKey
	return user, nil
}
