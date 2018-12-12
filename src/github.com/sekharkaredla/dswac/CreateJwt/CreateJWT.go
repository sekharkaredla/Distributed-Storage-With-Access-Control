package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	var mySiginingKey = []byte("secret")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["ownerHash"] = "test_owner_hash"
	claims["filename"] = "filename"
	claims["fileHash"] = "file_hash"
	tokenString, _ := token.SignedString(mySiginingKey)
	fmt.Println(tokenString)
}
