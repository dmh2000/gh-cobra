/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

// get list of repos for the specified owner with name and star count
func apiRepos(ctx context.Context, owner string) ([]string, []int, error) {

	client := github.NewClient(nil)

	repos, _, err := client.Repositories.List(ctx, owner, nil)
	if err != nil {
		return nil, nil, err
	}

	var names []string
	var stars []int
	for _, repo := range repos {
		names = append(names, *repo.Name)
		stars = append(stars, *repo.StargazersCount)
	}

	return names, stars, nil
}

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "$gh-cobra <owner> Get list repos using github REST api",
	Long: `For the specified <owner>, get the list of public repos using the github REST api.
	Format : gh-cobra api <owner>
	Example: gh-cobra api octocat
	This command does not require authentication`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the owner name")
			return
		}

		owner := args[0]
		fmt.Println(owner)

		ctx := context.Background()
		repos, stars, err := apiRepos(ctx, owner)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, repo := range repos {
			if i < 9 {
				fmt.Printf("%d. %*s - %d\n", i+1, -33, repo, stars[i])
			} else {
				fmt.Printf("%d. %*s - %d\n", i+1, -32, repo, stars[i])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
