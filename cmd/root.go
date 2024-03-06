/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-cobra",
	Short: "A simple application of Github Cobra with subcommands",
	Long: `
	A simple application of Github Cobra with subcommands
	Commands:
	1. gh-cobra api <owner> : Get list repos using github REST api
	2. gh-cobra graphql <owner> : Get list of repos using github GraphQl api
	3. gh-cobra explain <question> : Use OpenAI to answer questions about Linux utilities
	4. gh-cobra help : Help about any command`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
