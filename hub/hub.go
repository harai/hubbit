package hub

import (
	"net/http"

	"github.com/github/hub/github"
	"github.com/octokit/go-octokit/octokit"
)

// Authenticate authenticates and returns GitHub API Client.
func Authenticate() *octokit.Client {
	config := github.CurrentConfig().Find("github.com")
	if config == nil {
		panic("No github.com config found in Hub.")
	}
	tokenAuth := octokit.TokenAuth{AccessToken: config.AccessToken}
	return octokit.NewClientWith(
		"https://api.github.com", "Hub", tokenAuth, &http.Client{})
}
