/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain [bash utility name,...]",
	Short: "$gh-cobra explain [bash utility name,...] : Use OpenAI to get bash help",
	Long: `
	This command uses OpenAI to get help for bash utilities,  using the Langchain API.
	It takes a list of bash utility names as arguments and returns the help for each utility.
	Format : gh-cobra explain [utility name,...]
	Example: gh-cobra explain ls,cat,echo
	This command requires an OpenAI API Key in an environment variable named OPENAI_API_KEY.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the utility name")
			return
		}
		for _, util := range args {
			if lookupUtil(util) {
				result, err := Explain(util)
				if err != nil {
					fmt.Println(err)
					fmt.Println("This command requires a valid OpenAI API Key in an environment variable named OPENAI_API_KEY.")
					break
				}
				fmt.Println(result)
			} else {
				fmt.Printf("Sorry I don't recognise [%s]. I only know basic bash utilities\n", util)
			}
			fmt.Println("------------------------------------------------------------------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

}
