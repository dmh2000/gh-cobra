/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// formatting the graphql json is tricky
var query = []string{
	"{\"query\":",
	`"query {search(query: \"is:public archived:false org:`,
	// owner name
	`\", type: REPOSITORY, first: 100) { repositoryCount edges { node { ... on Repository { name stargazerCount}}}}}"`,
	"}",
}

// JSON result format (figured out manually by printing the result types)
// probably should use a client package
type repoData struct {
	Data struct {
		Search struct {
			Edges []struct {
				Node struct {
					Name           string `json:"name"`
					StargazerCount int    `json:"stargazerCount"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"search"`
	} `json:"data"`
}

func gqlRepos(owner string) ([]string, []int, error) {
	token := os.Getenv("GITHUB_TOKEN")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// construct query
	q := fmt.Sprintf("%s%s%s%s", query[0], query[1], owner, query[2])

	// fmt.Println(q)

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", nil)
	if err != nil {
		return nil, nil, err
	}

	reqBody := io.NopCloser(strings.NewReader(q))

	req.Header.Add("Authorization", "bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	req.Body = reqBody
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var result repoData
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, nil, err
	}

	names := []string{}
	stars := []int{}

	for _, v := range result.Data.Search.Edges {
		names = append(names, v.Node.Name)
		stars = append(stars, v.Node.StargazerCount)
	}

	return names, stars, nil
}

var graphqlCmd = &cobra.Command{
	Use:   "graphql [owner]",
	Short: "$gh-cobra [owner]  Get list of repos using github GraphQl api",
	Long: `For the specified [owner], get the list of public repos using the github GraphQl api.
	Format : gh-cobra graphql [owner]
	Example: gh-cobra graphql octocat
	This command requires a GitHub authentication token in a GITHUB_TOKEN environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the owner name")
			return
		}

		owner := args[0]
		fmt.Println(owner)

		names, stars, err := gqlRepos(owner)
		if err != nil {
			fmt.Println(err)
			fmt.Println("This command requires a valid GitHub authentication token in a GITHUB_TOKEN environment variable.")
		}

		if len(names) == 0 {
			fmt.Println("No repos found")
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
	rootCmd.AddCommand(graphqlCmd)

}
