package commitpipeline

import (
	"testing"

	history "github.com/aevea/git/v4"
	"github.com/stretchr/testify/assert"
)

func TestAllCommits(t *testing.T) {
	gitRepo, err := history.OpenGit("../../testdata/long-history")

	assert.NoError(t, err)

	options := Options{
		Path:           "",
		UpstreamBranch: "master",
		Limit:          0,
		AllCommits:     true,
	}

	pipeline, err := New(&options)

	assert.NoError(t, err)

	commits, err := pipeline.commitsBetweenBranches(gitRepo)

	assert.NoError(t, err)
	assert.Len(t, commits, 102)

	lastCommit, err := gitRepo.Commit(commits[0])

	assert.NoError(t, err)
	assert.Equal(t, "chore: add 100 file\n", lastCommit.Message)

	firstCommit, err := gitRepo.Commit(commits[101])

	assert.NoError(t, err)
	assert.Equal(t, "Initial commit\n", firstCommit.Message)
}

func TestLimitCommits(t *testing.T) {
	gitRepo, err := history.OpenGit("../../testdata/long-history")

	assert.NoError(t, err)

	options := Options{
		Path:           "",
		UpstreamBranch: "master",
		Limit:          50,
		AllCommits:     false,
	}

	pipeline, err := New(&options)

	assert.NoError(t, err)

	commits, err := pipeline.commitsBetweenBranches(gitRepo)

	assert.NoError(t, err)
	assert.Len(t, commits, 50)

	lastCommit, err := gitRepo.Commit(commits[0])

	assert.NoError(t, err)
	assert.Equal(t, "chore: add 100 file\n", lastCommit.Message)

	firstCommit, err := gitRepo.Commit(commits[49])

	assert.NoError(t, err)
	assert.Equal(t, "chore: add 51 file\n", firstCommit.Message)
}
