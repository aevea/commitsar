package dispatcher

import "sync"

func (dispatch *Dispatcher) handleErrors(
	wg *sync.WaitGroup,
	channel <-chan PipelineError,
	results *Results,
) {
	defer wg.Done()

	for message := range channel {
		dispatch.debugLogger.Printf("[%s] %s", message.PipelineName, message.Error)

		results.Errors = append(results.Errors, message)
	}
}
