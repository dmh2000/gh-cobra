/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
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
					Name string `graphql:"name"`
				} `graphql:"... on Repository"`
			}
		}
	} `graphql:"search(query: $query, type:REPOSITORY, first: $first)"`
}

var variables = map[string]any{}

func fetchRepos(owner string) []string {
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
		log.Fatal(err.Error())
	}

	var repoNames []string
	for _, edge := range repos.Search.Edges {
		repoNames = append(repoNames, edge.Node.SearchedRepository.Name)
	}
	return repoNames
}

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client [owner]",
	Short: "client [owner] print list of repositories from github graphql search",
	Long: `
	print list of repositories from github graphql search api using shurcool/graphql package
	Format : gh-cobra client [owner]
	Example: gh-cobra client octocat
	This command requires a GitHub authentication token in a GITHUB_TOKEN environment variable
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the owner name")
			return
		}

		owner := args[0]
		fmt.Println(owner)

		repoNames := fetchRepos(owner)
		for i, name := range repoNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

}
