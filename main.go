package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/github/hub/github"
	"github.com/octokit/go-octokit/octokit"
)

func main() {
	_, err := exec.Command("git", "version").Output()
	if err != nil {
		panic("Panicked1")
	}
	c, err := github.CurrentConfig().DefaultHost()
	if err != nil {
		panic("Panicked2")
	}
	tokenAuth := octokit.TokenAuth{AccessToken: c.AccessToken}
	host := c.Host
	client := octokit.NewClientWith("https://api.github.com", "Hub", tokenAuth, &http.Client{})
	url, err := octokit.UserURL.Expand(octokit.M{"user": "harai"})
	if err != nil {
		panic("Panicked2")
	}
	user, err := client.Users(url).One()
	fmt.Println(user.ReposURL)
	fmt.Println(tokenAuth)
	fmt.Println(c.AccessToken)
	fmt.Println(host)
}
