package version_runner

type Logger interface {
	Printf(string, ...interface{})
}

// VersionInfo houses all info regarding to the version. It is here to avoid global variables in this package.
type VersionInfo struct {
	Version string
	Date string
}

// Run executes the command to get version and prints it into stdout
func Run(versionInfo VersionInfo, logger Logger) error {
	logger.Printf("Commitsar version: %s\t Built on: %s", versionInfo.Version, versionInfo.Date)
	return nil
}