package root_runner

import (
	"io/ioutil"
	"log"
	"os"
)

type Runner struct {
	DebugLogger *log.Logger
	Logger      *log.Logger
	Strict      bool
	// Debug is a deprecated flag. Will be replaced as all repos accept the debugLogger
	Debug bool
}

type RunnerOptions struct {
	// Path to repository
	Path string
	// UpstreamBranch is the branch against which to check
	UpstreamBranch string
	// Limit will limit how far back to check on upstream branch.
	Limit int
	// AllCommits will check all the commits on the upstream branch. Regardless of Limit setting.
	AllCommits bool
}

// New returns a new instance of a RootRunner with fallbacks for logging
func New(logger *log.Logger, debugLogger *log.Logger, strict bool) *Runner{
	if logger == nil {
		logger = log.New(os.Stdout, "", 0)
	}
	if debugLogger == nil {
		debugLogger = log.New(ioutil.Discard, "[DEBUG] ", 0)
	}

	return &Runner{
		Logger:logger,
		DebugLogger:debugLogger,
		Strict:strict,
	}
}
