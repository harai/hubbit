package main

import (
	"fmt"
	"os"

	"github.com/harai/hubbit/git"
	"github.com/harai/hubbit/github"
	"github.com/harai/hubbit/hub"
	"github.com/octokit/go-octokit/octokit"
)

func main() {
	issue, err := git.CurrentIssueNo()
	if err != nil {
		fmt.Println("not in a issue branch")
		os.Exit(1)
	}
	fmt.Println(github.IssueAsHashtag(issue))
	client := hub.Authenticate()
	url, err := octokit.CurrentUserURL.Expand(octokit.M{})
	if err != nil {
		panic("Not happen")
	}
	fmt.Println(url)
	user, result := client.Users(url).One()
	if result.Err != nil {
		panic(result.Response.Status)
	}
	fmt.Println(user.ReposURL)
}
