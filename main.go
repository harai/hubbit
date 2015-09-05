package main

// import (
// 	"fmt"
// 	"os"
//
// 	"github.com/harai/hubbit/git"
// 	"github.com/harai/hubbit/github"
// 	"github.com/harai/hubbit/hub"
// 	"github.com/octokit/go-octokit/octokit"
// )
//
// func main() {
// 	issue, err := git.CurrentIssueNo()
// 	if err != nil {
// 		fmt.Println("not in an issue branch")
// 		os.Exit(1)
// 	}
// 	fmt.Println(github.IssueAsHashtag(issue))
// 	client := hub.Authenticate()
// 	url, err := octokit.CurrentUserURL.Expand(octokit.M{})
// 	if err != nil {
// 		panic("Not happen")
// 	}
// 	fmt.Println(url)
// 	user, result := client.Users(url).One()
// 	if result.Err != nil {
// 		panic(result.Response.Status)
// 	}
// 	fmt.Println(user.ReposURL)
// }

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/harai/hubbit/git"
	"github.com/harai/hubbit/github"
)

func commit(c *cli.Context) {
	issue, err := git.CurrentIssueNo()
	if err != nil {
		log.Fatalln("not in an issue branch")
	}
	tmpl := github.IssueAsHashtag(issue)
	if err := git.CommitWithTemplate(tmpl); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "hubbit"
	app.Usage = "custom commands for Git/GitHub"
	app.Commands = []cli.Command{
		{
			Name:    "commit",
			Aliases: []string{"c"},
			Usage:   "record changes to the git repository with a hashtag",
			Action:  commit,
		},
	}

	app.Run(os.Args)
}
