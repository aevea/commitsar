package root_runner

import (
	"context"
	"errors"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"

	"github.com/aevea/commitsar/pkg/jira"
	history "github.com/aevea/git/v2"
	"github.com/google/go-github/v32/github"
	"github.com/spf13/viper"
)

// RunPullRequest starts the runner for the PullRequest pipeline
func (runner *Runner) RunPullRequest(jiraKeys []string) ([]string, error) {
	ghToken := viper.GetString("GITHUB_TOKEN")

	gitRepo, err := history.OpenGit(".", log.New(ioutil.Discard, "", 0))

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	currentCommit, err := gitRepo.CurrentCommit()

	if err != nil {
		return nil, err
	}

	prs, response, err := client.PullRequests.ListPullRequestsWithCommit(ctx, "aevea", "commitsar", currentCommit.Hash.String(), nil)

	if err != nil {
		return nil, err
	}

	if len(prs) == 0 {
		runner.DebugLogger.Printf("current commit %s", currentCommit.Hash.String())
		runner.DebugLogger.Print(response)

		return nil, errors.New("No linked PullRequests found")
	}

	title := prs[0].Title

	references, err := jira.FindReferences(jiraKeys, *title)

	if err != nil {
		return nil, err
	}

	return references, nil
}
