package main

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"

	"github.com/bradleyfalzon/ghinstallation"
)

var githubAppID = 1
var githubInstallID = 1
var githubSecret = ""
var gitPEM = `-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----`

func main() {
	// POC connecting to github via google package
	// manpiulate data.. commit, branch, create a repo
	//  look at data comparisons
	ctx := context.Background()

	itr, err := ghinstallation.New(http.DefaultTransport, githubAppID, githubInstallID, []byte(gitPEM))
	if err != nil {
		panic(err)
	}

	client := github.NewClient(&http.Client{Transport: itr})

	// create a new private repository named "foo"
	repo := &github.Repository{
		Name:    github.String("foo"),
		Private: github.Bool(false),
	}
	_, _, err = client.Repositories.Create(ctx, "bhsi-underwriting", repo)
	if err != nil {
		panic(err)
	}
}
