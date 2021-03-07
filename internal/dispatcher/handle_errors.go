package dispatcher

import (
	"sync"

	"github.com/apex/log"
)

func (dispatch *Dispatcher) handleErrors(
	wg *sync.WaitGroup,
	channel <-chan PipelineError,
	results *Results,
) {
	defer wg.Done()

	for message := range channel {
		log.Debugf("[%s] %s", message.PipelineName, message.Error)

		results.Errors = append(results.Errors, message)
	}
}
