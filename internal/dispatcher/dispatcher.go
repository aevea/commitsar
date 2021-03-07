package dispatcher

import (
	"runtime"
	"sync"
)

// Dispatcher is the central place which runs Pipelines. maxWorkers is based by default on number of CPUs * 2 accounting for modern CPU architectures.
type Dispatcher struct {
	maxWorkers int
}

// Results contains the aggregated results of both the succesful and error pipelines.
type Results struct {
	SuccessfulPipelines []PipelineSuccess
	Errors              []PipelineError
}

// New returns a set up instance of Dispatcher
func New() *Dispatcher {

	return &Dispatcher{maxWorkers: runtime.NumCPU() * 2}
}

// RunPipelines will run asynchronously all pipelines passed to it. It is limited only by the maxWorkers field on Dispatcher.
func (dispatch *Dispatcher) RunPipelines(pipelines []Pipeliner) *Results {
	// pipelineChannel is limited to the amount of CPU to prevent overloading the machie
	pipelineChannel := make(chan Pipeliner, dispatch.maxWorkers)
	successChannel := make(chan PipelineSuccess, 10)
	errChannel := make(chan PipelineError, 10)
	results := &Results{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for _, pipeline := range pipelines {
			pipelineChannel <- pipeline
		}
		close(pipelineChannel)
		wg.Done()
	}()

	wg.Add(1)
	go dispatch.handleSuccess(&wg, successChannel, results)

	wg.Add(1)
	go dispatch.handleErrors(&wg, errChannel, results)

	wg.Add(1)
	go dispatch.work(&wg, pipelineChannel, successChannel, errChannel)

	wg.Wait()

	return results
}
