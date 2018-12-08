package main

import (
	"fmt"

	ipfs_api "github.com/ipfs/go-ipfs-api"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(ipfs_api.EnvDir)
}
