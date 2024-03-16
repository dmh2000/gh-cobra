/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
)


func fetchReposGithub(owner string) ([]string, []int, error) {
	client := github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))

	var opt = &github.RepositoryListByUserOptions{
		Type: "public",
	}
	// list public repositories for owner
	repos, _, err := client.Repositories.ListByUser(context.Background(), owner, opt)

	if err != nil {
		fmt.Println(err)
		return  nil,nil,err
	}

	var names []string
	var stars []int
	for _, repo := range repos {
		names = append(names, *repo.Name)
		stars = append(stars, *repo.StargazersCount)
	}
	return names, stars, nil
}

var gogithubCmd = &cobra.Command{
	Use:   "gogithub [owner]",
	Short: "gogithub [owner] print list of repositories using google/go-github package",
	Long: `
	print list of repositories from github graphql search api using google/go-github package
	Format : gh-cobra gogithub [owner]
	Example: gh-cobra gogithub octocat
	This command requires a GitHub authentication token in a GITHUB_TOKEN environment variable
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the owner name")
			return
		}

		owner := args[0]
		fmt.Println(owner)

		names, stars, err := fetchReposGithub(owner)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Check if the GITHUB_TOKEN is set")
			return
		}

		for i, name := range names {
			if i < 9 {
				fmt.Printf("%d. %*s - %d\n", i+1, -33, name, stars[i])
			} else {
				fmt.Printf("%d. %*s - %d\n", i+1, -32, name, stars[i])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(gogithubCmd)

}
