package version_runner

import (
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/memory"
	"github.com/stretchr/testify/assert"
)

func TestVersionRun(t *testing.T) {
	handler := memory.New()

	log.SetHandler(handler)

	err := Run(VersionInfo{
		Version: "development",
		Date:    "2012-1-1",
	},
	)

	assert.NoError(t, err)
	assert.Equal(t, "Commitsar version: development\t Built on: 2012-1-1", handler.Entries[0].Message)
}
