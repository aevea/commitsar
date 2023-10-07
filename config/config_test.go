package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	viper.Reset()

	err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, false, viper.GetBool("verbose"), "expected verbose to be true, but got false")

	os.Clearenv()
}

func TestLoadConfigCustomPathFromEnv(t *testing.T) {
	viper.Reset()
	err := os.Setenv(CommitsarConfigPath, "./testdata")
	assert.NoError(t, err)

	err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"), "expected verbose to be true, but got false")

	os.Clearenv()
}

func TestLoadConfigCustomPathFromParams(t *testing.T){
	viper.Reset()
	viper.Set("commits.config-path", "./testdata")

	err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"), "expected verbose to be true, but got false")

	os.Clearenv()
}

func TestLoadConfigCustomPathParamOverridesEnv(t *testing.T){
	viper.Reset()
	err := os.Setenv(CommitsarConfigPath, "./wrong-path")
	assert.NoError(t, err)
	viper.Set("commits.config-path", "./testdata")

	err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"), "expected verbose to be true, but got false")

	os.Clearenv()
}
