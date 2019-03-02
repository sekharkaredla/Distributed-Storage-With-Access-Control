package cmd

import (
	"strings"

	ipfs "github.com/sekharkaredla/dswac/MainCode/IPFS"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"

	"github.com/spf13/cobra"
)

var CmdAdd = &cobra.Command{
	Use:   "add [filename to add]",
	Short: "Add a file to IPFS",
	Long:  `Add is for uploading a file to the IPFS system`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//add code for adding file to IPFS
		log.Info.Println("Adding Filename : " + strings.Join(args, " "))
		for _, each_file := range args {
			hash, err := ipfs.AddFileToIPFS(each_file)
			if err != nil {
				log.Error.Fatal("error adding to ipfs file : ", each_file)
			}
			log.Info.Println("filename : ", each_file, " , hash : ", hash)
		}
	},
}
