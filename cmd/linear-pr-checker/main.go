package main

import (
	"context"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	repo := os.Getenv("GITHUB_REPOSITORY")
	prid, _ := strconv.Atoi(parsePullRequestId(os.Getenv("GITHUB_REF")))
	justrepo := repo[strings.LastIndex(repo, "/")+1:]
	token := os.Getenv("INPUT_GITHUB_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	pullrequest, _, err := client.PullRequests.Get(ctx, owner, justrepo, prid)
	if err != nil {
		println(err.Error())
	}
	title := pullrequest.Title

	println(*title)
	issueId, err := praseIssueFromBranch(*title)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	print(issueId)
}

func parsePullRequestId(ref string) string {
	var re = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return re.FindString(ref)
}

func praseIssueFromBranch(input string) (string, error) {
	re := regexp.MustCompile("^(\\w*)(?:\\((.*)\\))?\\:\\s(.*)$")
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		for i, s := range match {
			if i == 2 && s != "" {
				return s, nil
			}
		}
	}
	return "", errors.New("missing linear ticket")
}
