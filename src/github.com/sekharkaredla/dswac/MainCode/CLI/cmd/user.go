package cmd

import (
	"fmt"

	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"
	user "github.com/sekharkaredla/dswac/MainCode/UserDetails"

	"github.com/spf13/cobra"
)

var CmdUser = &cobra.Command{
	Use:   "create",
	Short: "create a new user",
	Long:  `Creating a new user with own public and private key`,
	Run: func(cmd *cobra.Command, args []string) {
		//add code for creating a new user
		log.Info.Println("Creating a new user")
		var username string
		fmt.Print("enter username : ")
		fmt.Scanln(&username)
		userDetails, err := user.CreateNewUser(username)
		if err != nil {
			log.Error.Fatal("unable to create user")
		}
		log.Info.Println(userDetails)
	},
}
