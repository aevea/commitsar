package version_runner

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVersionRun(t *testing.T) {
	var testString bytes.Buffer

	testLogger := log.Logger{}
	testLogger.SetOutput(&testString)
	
	err := Run(VersionInfo{
			Version: "development",
			Date:    "2012-1-1",
		},
		&testLogger,
	)

	assert.NoError(t, err)
	assert.Equal(t, "Commitsar version: development\t Built on: 2012-1-1\n", testString.String())
}