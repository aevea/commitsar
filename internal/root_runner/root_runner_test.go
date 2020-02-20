package root_runner

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitsOnMaster(t *testing.T) {
	var testString bytes.Buffer

	testLogger := log.Logger{}
	testLogger.SetOutput(&testString)

	runner := Runner{
		DebugLogger: log.New(ioutil.Discard, "", 0),
		Logger:      &testLogger,
		Strict:      false,
		Debug:       false,
	}

	err := runner.Run("../../testdata/commits-on-master", "master" )

	assert.NoError(t, err)
	assert.Equal(t, "Starting analysis of commits on branch\n\n0 commits filtered out\n\nFound 1 commit to check\n\x1b[32mAll 1 commits are conventional commit compliant\n\x1b[0m\n", testString.String())
}

func TestCommitsOnBranch(t *testing.T) {
	var testString bytes.Buffer

	testLogger := log.Logger{}
	testLogger.SetOutput(&testString)

	runner := Runner{
		DebugLogger: log.New(ioutil.Discard, "", 0),
		Logger:      &testLogger,
		Strict:      false,
		Debug:       false,
	}

	err := runner.Run("../../testdata/commits-on-different-branches", "master")

	assert.Error(t, err)
}

func TestFromToCommits(t *testing.T) {
	var testString bytes.Buffer

	testLogger := log.Logger{}
	testLogger.SetOutput(&testString)

	runner := Runner{
		DebugLogger: log.New(ioutil.Discard, "", 0),
		Logger:      &testLogger,
		Strict:      false,
		Debug:       false,
	}

	err := runner.Run("../../testdata/commits-on-different-branches", "master", "7dbf3e7db93ae2e02902cae9d2f1de1b1e5c8c92...d0240d3ed34685d0a5329b185e120d3e8c205be4")

	assert.NoError(t, err)
}
