package main

import (
	"github.com/sekharkaredla/dswac/MainCode/CLI/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "dswac"}
	rootCmd.AddCommand(cmd.CmdAdd, cmd.CmdUser, cmd.CmdLogin, cmd.CmdUpload, cmd.CmdDetails)
	rootCmd.Execute()
}
