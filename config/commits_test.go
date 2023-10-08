package config

import (
	"os"
	"testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCommitConfig(t *testing.T) {
	os.Clearenv()

	defaultConfig := CommitConfig()

	assert.Equal(t, true, defaultConfig.Strict, "expect strict to be true by default")
	assert.Equal(t, 0, defaultConfig.Limit, "expect the limit to be 0 by default")
	assert.Equal(t, false, defaultConfig.AllCommits, "expect AllCommits to be true by default")
	assert.Equal(t, "origin/master", defaultConfig.UpstreamBranch, "expect UpstreamBranch to be origin/master by default")
	assert.Equal(t, []string{}, defaultConfig.RequiredScopes, "expect required scopes to be empty slice by default")

	err := os.Setenv(CommitsarConfigPath, "./testdata")
	assert.NoError(t, err)

	err = LoadConfig()
	assert.NoError(t, err)

	commitConfig := CommitConfig()

	assert.Equal(t, false, commitConfig.Strict, "expect strict to be false as opposed to the default of true")
	assert.Equal(t, 100, commitConfig.Limit, "expect limit to be 100 as opposed to the default of 0")
	assert.Equal(t, true, commitConfig.AllCommits, "expect strict to be false as opposed to the default of false")
	assert.Equal(t, "origin/main", commitConfig.UpstreamBranch, "expect UpstreamBranch to be origin/main as opposed to the default of origin/master")
	assert.Equal(t, []string{}, commitConfig.RequiredScopes, "expect required scopes to be empty slice same as default for backward compatibility")

	os.Clearenv()

}

func TestDefaultConfig(t *testing.T) {
	viper.Reset()

	err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, false, viper.GetBool("verbose"))

	os.Clearenv()
}

func TestLoadConfigCustomPathFromEnv(t *testing.T) {
	viper.Reset()
	err := os.Setenv(CommitsarConfigPath, "./testdata")
	assert.NoError(t, err)

	err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"))

	os.Clearenv()
}

func TestLoadConfigCustomPathFromParams(t *testing.T){
	viper.Reset()
	viper.Set("config-path", "./testdata")

	err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"))

	os.Clearenv()
}

func TestLoadConfigCustomPathParamOverridesEnv(t *testing.T){
	viper.Reset()
	err := os.Setenv(CommitsarConfigPath, "./wrong-path")
	assert.NoError(t, err)
	viper.Set("config-path", "./testdata")

	err = LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, true, viper.GetBool("verbose"))

	os.Clearenv()
}
