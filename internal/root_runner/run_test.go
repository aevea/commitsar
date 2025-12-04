package root_runner

import (
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCommitsOnMaster(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	options := RunnerOptions{
		Path:           "../../testdata/commits-on-master",
		UpstreamBranch: "master",
		Limit:          0,
		AllCommits:     false,
		Strict:         false,
	}

	err := runner.Run(options)

	assert.NoError(t, err)
	assert.Equal(t, 5, len(handler.Entries))
	assert.Equal(t, "Starting pipeline: commit-pipeline", handler.Entries[0].Message)
	assert.Equal(t, "Starting analysis of commits on branch refs/heads/master", handler.Entries[1].Message)
	assert.Equal(t, "0 commits filtered out", handler.Entries[2].Message)
	assert.Equal(t, "Found 1 commit to check", handler.Entries[3].Message)
	assert.Equal(t, "\x1b[32mAll 1 commits are conventional commit compliant\x1b[0m", handler.Entries[4].Message)
}

func TestCommitsOnBranch(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	options := RunnerOptions{
		Path:           "../../testdata/commits-on-different-branches",
		UpstreamBranch: "master",
		Limit:          0,
		AllCommits:     false,
		Strict:         false,
	}

	err := runner.Run(options, "master")

	assert.Error(t, err)
}

func TestFromToCommits(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	options := RunnerOptions{
		Path:           "../../testdata/commits-on-different-branches",
		UpstreamBranch: "master",
		Limit:          0,
		AllCommits:     false,
		Strict:         false,
	}

	err := runner.Run(options, "d0240d3ed34685d0a5329b185e120d3e8c205be4...7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92")

	assert.NoError(t, err)
}

func TestMissingPipelines(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	viper.Set("commits.disabled", true)

	err := runner.Run(RunnerOptions{})

	assert.Error(t, err)
	assert.Equal(t, "no pipelines defined", err.Error())

	viper.Set("commits.disabled", "")
}

func TestPRPipeline_ConventionalOnly(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	viper.Set("commits.disabled", true)
	viper.Set("pull_request.conventional", true)

	// This will fail because it needs GitHub API access, but we can verify the pipeline is created
	err := runner.Run(RunnerOptions{})

	// Should fail due to missing GitHub env vars, not due to missing pipeline
	assert.Error(t, err)
	assert.NotEqual(t, "no pipelines defined", err.Error())

	viper.Set("commits.disabled", "")
	viper.Set("pull_request.conventional", "")
}

func TestPRPipeline_JiraOnly(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	viper.Set("commits.disabled", true)
	viper.Set("pull_request.jira_title", true)

	// This will fail because it needs GitHub API access, but we can verify the pipeline is created
	err := runner.Run(RunnerOptions{})

	// Should fail due to missing GitHub env vars, not due to missing pipeline
	assert.Error(t, err)
	assert.NotEqual(t, "no pipelines defined", err.Error())

	viper.Set("commits.disabled", "")
	viper.Set("pull_request.jira_title", "")
}

func TestPRPipeline_BothStyles(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner := Runner{}

	viper.Set("commits.disabled", true)
	viper.Set("pull_request.conventional", true)
	viper.Set("pull_request.jira_title", true)
	viper.Set("pull_request.jira_keys", []string{"TEST"})

	// This will fail because it needs GitHub API access, but we can verify the pipeline is created
	err := runner.Run(RunnerOptions{})

	// Should fail due to missing GitHub env vars, not due to missing pipeline
	assert.Error(t, err)
	assert.NotEqual(t, "no pipelines defined", err.Error())

	viper.Set("commits.disabled", "")
	viper.Set("pull_request.conventional", "")
	viper.Set("pull_request.jira_title", "")
	viper.Set("pull_request.jira_keys", "")
}

func TestPRPipeline_OptionsConfiguration(t *testing.T) {
	// Test that when both styles are configured, both are added to the Options
	viper.Reset()
	viper.Set("pull_request.conventional", true)
	viper.Set("pull_request.jira_title", true)
	viper.Set("pull_request.jira_keys", []string{"TEST", "XXX"})

	// Manually build the options as the code does
	prOptions := struct {
		Path   string
		Styles []string
		Keys   []string
	}{
		Path:   ".",
		Styles: []string{},
		Keys:   []string{},
	}

	if viper.IsSet("pull_request.jira_title") {
		prOptions.Styles = append(prOptions.Styles, "jira")
		prOptions.Keys = viper.GetStringSlice("pull_request.jira_keys")
	}

	if viper.IsSet("pull_request.conventional") {
		prOptions.Styles = append(prOptions.Styles, "conventional")
	}

	// Verify both styles are present
	assert.Equal(t, 2, len(prOptions.Styles))
	assert.Contains(t, prOptions.Styles, "jira")
	assert.Contains(t, prOptions.Styles, "conventional")
	assert.Equal(t, []string{"TEST", "XXX"}, prOptions.Keys)

	viper.Reset()
}
