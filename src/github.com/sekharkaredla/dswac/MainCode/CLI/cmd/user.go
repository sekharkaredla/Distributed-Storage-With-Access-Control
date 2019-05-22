package cmd

import (
	"fmt"

	keystore "github.com/sekharkaredla/dswac/MainCode/Keystorage"
	log "github.com/sekharkaredla/dswac/MainCode/LogSetup"

	"github.com/spf13/cobra"
)

var CmdUser = &cobra.Command{
	Use:   "create",
	Short: "create a new user",
	Long:  `Creating a new user with own public and private key`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info.Println("Creating a new user")
		fmt.Println("enter 1 for creating new pub/pri key pair, enter 2 for using an existing one: ")
		var option int
		fmt.Scanln(&option)
		if option == 1 {
			var password string
			fmt.Print("enter password : ")
			fmt.Scanln(&password)
			userAccount, err := keystore.GenerateNewAccount(password)
			if err != nil {
				log.Error.Fatalln("unable to create user")
			}
			log.Info.Println(userAccount)
		} else {
			fmt.Println("enter private key : ")
			var privateKey string
			fmt.Scanln(&privateKey)
			fmt.Println("enter passphrase : ")
			var passphrase string
			fmt.Scanln(&passphrase)
			userAccount, err := keystore.CreateAccountFromPrivateKey(privateKey, passphrase)
			if err != nil {
				log.Error.Fatalln("unable to add user")
			}
			log.Info.Println(userAccount)
		}
	},
}
