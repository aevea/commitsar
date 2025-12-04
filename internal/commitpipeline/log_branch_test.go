package commitpipeline

import (
	"testing"

	history "github.com/aevea/git/v4"
	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
	"github.com/stretchr/testify/assert"
)

func TestLogBranch(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	runner, err := New(nil)

	assert.NoError(t, err)

	gitRepo, err := history.OpenGit("../../testdata/commits-on-master")

	assert.NoError(t, err)

	err = runner.logBranch(gitRepo)

	assert.NoError(t, err)
	assert.Equal(t, "Starting analysis of commits on branch refs/heads/master", handler.Entries[0].Message)
}
