package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	if !isValidIssue(issueId) {
		println("Not a Valid Issue.")
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

func isValidIssue(issueId string) bool {
	query := fmt.Sprintf("query Issue {  issue(id: \"%s\") {    id    title   }}", issueId)
	url := "https://api.linear.app/graphql"
	method := "POST"
	values := map[string]interface{}{"query": query}
	json_data, err := json.Marshal(values)

	if err != nil {
		fmt.Println(err)
		return false
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", os.Getenv("LINEAR_TOKEN"))
	fmt.Println(os.Getenv("INPUT_LINEAR_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil || strings.Contains(string(body), "error") {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(body))
	return true
}
