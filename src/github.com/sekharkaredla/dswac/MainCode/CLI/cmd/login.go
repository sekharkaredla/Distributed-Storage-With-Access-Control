package cmd

import (
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"

	"github.com/spf13/cobra"
)

var CmdLogin = &cobra.Command{
	Use:   "login",
	Short: "login with options",
	Long:  `Login using the private keys generated on this system, or user login using the private and public key also by giving private and public key in format pub:private`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//add code for logging in user
		if len(args) == 0 {
			log.Info.Println("List of users available for login")
		} else {
			log.Info.Println("pub : private key = ", args[0])
		}
	},
}
