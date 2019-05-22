package cmd

import (
	"fmt"

	ganache "github.com/sekharkaredla/dswac/MainCode/Ganache"
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
			params := make([]string, 1)
			params[0] = password
			data, err := keystore.GenerateNewAccount(ganache.GanacheURL, "hello")
			if err != nil {
				log.Error.Fatalln("unable to create user", err)
			}
			log.Info.Println(data)
		} else {
			fmt.Println("enter private key : ")
			var privateKey string
			fmt.Scanln(&privateKey)
			fmt.Println("enter passphrase : ")
			var passphrase string
			fmt.Scanln(&passphrase)
			err := keystore.CreateAccountFromPrivateKey(ganache.GanacheURL, privateKey, passphrase, false)
			if err != nil {
				log.Error.Fatalln("unable to add user")
			}
			log.Info.Println(err)
		}
	},
}
