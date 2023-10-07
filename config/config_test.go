package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigCustomPath(t *testing.T) {
	viper.Reset()
	err := os.Setenv(CommitsarConfigPath, "./testdata")
	assert.NoError(t, err)

	err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"), "expected verbose to be true, but got false")

	os.Clearenv()
}
