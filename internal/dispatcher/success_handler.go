package dispatcher

import (
	"sync"

	"github.com/apex/log"
)

func (dispatch *Dispatcher) handleSuccess(
	wg *sync.WaitGroup,
	channel <-chan PipelineSuccess,
	results *Results,
) {
	defer wg.Done()

	for message := range channel {
		log.Debugf("[%s] %s", message.PipelineName, message.Message)

		results.SuccessfulPipelines = append(results.SuccessfulPipelines, message)
	}
}
