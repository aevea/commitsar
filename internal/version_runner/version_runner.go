package version_runner

import "github.com/apex/log"

// VersionInfo houses all info regarding to the version. It is here to avoid global variables in this package.
type VersionInfo struct {
	Version string
	Date    string
}

// Run executes the command to get version and prints it into stdout
func Run(versionInfo VersionInfo) error {
	log.Infof("Commitsar version: %s\t Built on: %s", versionInfo.Version, versionInfo.Date)
	return nil
}
