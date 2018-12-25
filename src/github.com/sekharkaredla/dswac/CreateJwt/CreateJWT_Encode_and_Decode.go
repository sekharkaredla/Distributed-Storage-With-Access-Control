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
	claims["fileName"] = "filename"
	claims["fileHash"] = "file_hash"
	tokenString, _ := token.SignedString(mySiginingKey)
	fmt.Println(tokenString)

	claims = jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println("Error is getting claims ", err)
	}
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}
