package cmd

import (
	"strings"

	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"

	"github.com/spf13/cobra"
)

var CmdUpload = &cobra.Command{
	Use:   "upload",
	Short: "upload file details into blockchain",
	Long:  `Upload the filehash,filename,userhash as a JWT to blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		//add code for uploading JWT into blockchain
		if len(args) == 0 {
			log.Info.Println("List of files added to IPFS from this system")
		} else {
			log.Info.Println("file hash to be added : ", strings.Join(args, " "))
		}
	},
}
