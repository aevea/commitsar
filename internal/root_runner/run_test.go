package root_runner

import (
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
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

	err := runner.Run(options, "7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92...d0240d3ed34685d0a5329b185e120d3e8c205be4")

	assert.NoError(t, err)
}
