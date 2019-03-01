package cmd

import (
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"

	"github.com/spf13/cobra"
)

var CmdUser = &cobra.Command{
	Use:   "create",
	Short: "create a new user",
	Long:  `Creating a new user with own public and private key`,
	Run: func(cmd *cobra.Command, args []string) {
		//add code for creating a new user
		log.Info.Println("Creating a new user")
	},
}
