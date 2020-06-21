package dispatcher

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDispatcher(t *testing.T) {
	dispatcher := New(nil)
	dispatcher.maxWorkers = 1

	pipelines := []Pipeliner{TestPipeline{TestName: "pipeline1", TestFn: successPipeline}}

	results := dispatcher.RunPipelines(pipelines)

	assert.Equal(t, "pipeline1", results.SuccessfulPipelines[0].PipelineName)
	assert.Equal(t, "It succeeded", results.SuccessfulPipelines[0].Message)

	pipelines2 := []Pipeliner{TestPipeline{TestName: "pipeline1", TestFn: successPipeline}, TestPipeline{TestName: "pipeline2", TestFn: failPipeline}, TestPipeline{TestName: "pipeline3", TestFn: failPipeline}, TestPipeline{TestName: "pipeline4", TestFn: successPipeline}}

	results2 := dispatcher.RunPipelines(pipelines2)

	assert.Equal(t, 2, len(results2.SuccessfulPipelines))
	assert.Equal(t, 2, len(results2.Errors))
}

type TestPipeline struct {
	TestName string
	TestFn   func(chan PipelineError) *PipelineSuccess
}

func (p TestPipeline) Name() string {
	return p.TestName
}

func (p TestPipeline) Run(errChannel chan PipelineError) *PipelineSuccess {
	return p.TestFn(errChannel)
}

func successPipeline(chan PipelineError) *PipelineSuccess {
	return &PipelineSuccess{
		PipelineName: "pipeline1",
		Message:      "It succeeded",
	}
}

func failPipeline(errChan chan PipelineError) *PipelineSuccess {
	errChan <- PipelineError{PipelineName: "pipeline2", Data: []FailureData{{Name: "test", Value: "test2"}}, Error: errors.New("some error")}

	return nil
}
