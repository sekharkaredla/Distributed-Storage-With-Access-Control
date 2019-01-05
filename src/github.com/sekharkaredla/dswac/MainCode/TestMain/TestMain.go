package main

import (
	"fmt"

	ganache "github.com/sekharkaredla/dswac/MainCode/Ganache"
)

func main() {
	out, _ := ganache.DeployContract()
	fmt.Println(out)
}
