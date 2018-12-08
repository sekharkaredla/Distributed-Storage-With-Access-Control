package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	fmt.Print("Enter File Name : ")
	var fileName string
	fmt.Scanf("%s", &fileName)
	fmt.Println("Adding to IPFS : ", fileName)
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(string(bytes)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println("Successfully added, IPFS Hash : %s", cid)
}
