package commitpipeline

import (
	"io/ioutil"
	"log"
)

// CommitPipeline is used to run checks in commits
type CommitPipeline struct {
	debugLogger *log.Logger
}

// New returns a setup instance of the CommitPipeline
func New(debugLogger *log.Logger) *CommitPipeline {
	if debugLogger == nil {
		debugLogger = log.New(ioutil.Discard, "", 0)
	}

	return &CommitPipeline{debugLogger: debugLogger}
}
