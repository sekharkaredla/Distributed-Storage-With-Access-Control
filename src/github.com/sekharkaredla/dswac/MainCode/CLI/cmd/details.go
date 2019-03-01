package cmd

import (
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"

	"github.com/spf13/cobra"
)

var CmdDetails = &cobra.Command{
	Use:   "details",
	Short: "details of all the files pushed by the user",
	Long:  `Details of files pushed by the user into the blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		//add code for creating a new user
		log.Info.Println("Details : ")
	},
}
