package dispatcher

import (
	"sync"

	"github.com/apex/log"
)

func (dispatch *Dispatcher) work(
	wg *sync.WaitGroup,
	pipelineChannel chan Pipeliner,
	successChan chan PipelineSuccess,
	errorChannel chan PipelineError,
) {
	defer wg.Done()
	defer close(successChan)
	defer close(errorChannel)

	for {
		pipeline, more := <-pipelineChannel

		if more {
			log.Infof("Starting pipeline: %s", pipeline.Name())
			success, err := pipeline.Run()

			if err != nil {
				errorChannel <- PipelineError{
					Error:        err,
					PipelineName: pipeline.Name(),
				}
			}

			if success != nil {
				successChan <- *success
			}
		} else {
			log.Debug("All pipelines complete")
			return
		}
	}

}
