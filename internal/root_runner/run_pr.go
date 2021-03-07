package root_runner

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/oauth2"

	"github.com/aevea/commitsar/pkg/jira"
	history "github.com/aevea/git/v3"
	"github.com/apex/log"
	"github.com/google/go-github/v32/github"
	"github.com/spf13/viper"
)

// RunPullRequest starts the runner for the PullRequest pipeline
func (runner *Runner) RunPullRequest(jiraKeys []string) ([]string, error) {
	ghToken := viper.GetString("GITHUB_TOKEN")

	if !viper.IsSet("GITHUB_REPOSITORY") {
		return nil, errors.New("missing GITHUB_REPOSITORY env variable. Please provide one in owner/repository format")
	}

	split := strings.Split(viper.GetString("GITHUB_REPOSITORY"), "/")

	gitRepo, err := history.OpenGit(".")

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

	prs, response, err := client.PullRequests.ListPullRequestsWithCommit(ctx, split[0], split[1], currentCommit.Hash.String(), nil)

	if err != nil {
		return nil, err
	}

	if len(prs) == 0 {
		log.Debugf("current commit %s", currentCommit.Hash.String())
		log.Debugf("%s", response)

		return nil, errors.New("No linked PullRequests found")
	}

	title := prs[0].Title

	references, err := jira.FindReferences(jiraKeys, *title)

	if err != nil {
		return nil, err
	}

	return references, nil
}
