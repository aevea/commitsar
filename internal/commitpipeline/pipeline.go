package commitpipeline

import (
	"io/ioutil"
	"log"
	"os"
)

type Pipeline struct {
	Logger      *log.Logger
	DebugLogger *log.Logger
	args        []string
	options     Options
}

type Options struct {
	// Path to repository
	Path string
	// UpstreamBranch is the branch against which to check
	UpstreamBranch string
	// Limit will limit how far back to check on upstream branch.
	Limit int
	// AllCommits will check all the commits on the upstream branch. Regardless of Limit setting.
	AllCommits bool
	Strict     bool
}

func New(logger, debugLogger *log.Logger, options *Options, args ...string) (*Pipeline, error) {
	if logger == nil {
		logger = log.New(os.Stdout, "", 0)
	}
	if debugLogger == nil {
		debugLogger = log.New(ioutil.Discard, "[DEBUG] ", 0)
	}

	if options == nil {
		options = &Options{
			Path:           ".",
			UpstreamBranch: "master",
			Limit:          0,
			AllCommits:     false,
			Strict:         true,
		}
	}

	return &Pipeline{
		Logger:      logger,
		DebugLogger: debugLogger,
		options:     *options,
		args:        args,
	}, nil
}

func (Pipeline) Name() string {
	return "commit-pipeline"
}
