package cmd

import (
	"strings"

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
	},
}
