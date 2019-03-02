package IPFS

import (
	"io/ioutil"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
)

const IPFSUrl = "localhost:5001"

func AddFileToIPFS(filename string) (string, error) {
	log.Info.Println("Adding to IPFS : ", filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error.Println("Error reading file, err : ", err)
		return "", err
	}
	sh := shell.NewShell(IPFSUrl)
	cid, err := sh.Add(strings.NewReader(string(bytes)))
	if err != nil {
		log.Error.Println("Error adding file to IPFS, err : ", err)
		return cid, err
	}
	return cid, nil
}
