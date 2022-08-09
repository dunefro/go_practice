package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v45/github"
)

func main() {
	ctx := context.Background()

	client := github.NewClient(nil)
	opt := &github.IssueListByRepoOptions{
		Since: time.Now().AddDate(-1, 0, 0),
	}
	issues, _, _ := client.Issues.ListByRepo(ctx, "google", "go-github", opt)
	for _, issue := range issues {
		fmt.Printf("%s\t%d\t%s\n", issue.GetRepositoryURL(), issue.GetNumber(), issue.GetURL())
	}
}
