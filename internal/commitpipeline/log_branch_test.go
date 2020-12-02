package commitpipeline

import (
	"bytes"
	"log"
	"testing"

	history "github.com/aevea/git/v2"
	"github.com/stretchr/testify/assert"
)

func TestLogBranch(t *testing.T) {
	var testString bytes.Buffer

	testLogger := log.Logger{}
	testLogger.SetOutput(&testString)

	runner, err := New(&testLogger, nil, nil)

	assert.NoError(t, err)

	gitRepo, err := history.OpenGit("../../testdata/commits-on-master", nil)

	assert.NoError(t, err)

	err = runner.logBranch(gitRepo)

	assert.NoError(t, err)
	assert.Equal(t, "Starting analysis of commits on branch refs/heads/master\n", testString.String())
}
