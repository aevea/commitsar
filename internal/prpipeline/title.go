package prpipeline

import (
	"context"
	"errors"
	"strings"

	history "github.com/aevea/git/v4"
	"github.com/apex/log"
	"github.com/google/go-github/v32/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func getPRTitle(path string) (*string, error) {
	ghToken := viper.GetString("GITHUB_TOKEN")

	if !viper.IsSet("GITHUB_REPOSITORY") {
		return nil, errors.New("missing GITHUB_REPOSITORY env variable. Please provide one in owner/repository format")
	}

	split := strings.Split(viper.GetString("GITHUB_REPOSITORY"), "/")

	gitRepo, err := history.OpenGit(path)

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

		return nil, errors.New("no linked PullRequests found")
	}

	return prs[0].Title, nil
}
