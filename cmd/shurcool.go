/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/graphql"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

/*
from github graphql api explorer:

	search(
		after: String
		before: String
		first: Int
		last: Int
		query: String!
		type: SearchType!
	): SearchResultItemConnection!
*/
var repos struct {
	Search struct {
		Edges []struct {
			Node struct {
				SearchedRepository struct {
					Name       string `graphql:"name"`
					StarGazers struct {
						TotalCount int `graphql:"totalCount"`
					} `graphql:"stargazers"`
				} `graphql:"... on Repository"`
			}
		}
	} `graphql:"search(query: $query, type:REPOSITORY, first: $first)"`
}

var variables = map[string]any{}

func fetchRepos(owner string) ([]string, []int, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("https://api.github.com/graphql", httpClient)

	// update the variables
	variables["query"] = graphql.String(fmt.Sprintf("is:public archived:false org:%s", owner))
	variables["first"] = graphql.Int(100)

	// execute the query
	err := client.Query(context.Background(), &repos, variables)
	if err != nil {
		return nil, nil, err
	}

	var names []string
	var stars []int
	for _, edge := range repos.Search.Edges {
		names = append(names, edge.Node.SearchedRepository.Name)
		stars = append(stars, edge.Node.SearchedRepository.StarGazers.TotalCount)
	}
	return names, stars, nil
}

var clientCmd = &cobra.Command{
	Use:   "shurcool [owner]",
	Short: "shurcool [owner] print list of repositories from github graphql search",
	Long: `
	print list of repositories from github graphql search api using shurcooL/graphql package
	Format : gh-cobra shurcool [owner]
	Example: gh-cobra shurcool octocat
	This command requires a GitHub authentication token in a GITHUB_TOKEN environment variable
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the owner name")
			return
		}

		owner := args[0]
		fmt.Println(owner)

		names, stars, err := fetchRepos(owner)
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
	rootCmd.AddCommand(clientCmd)

}
