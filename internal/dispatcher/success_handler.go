package dispatcher

import "sync"

func (dispatch *Dispatcher) handleSuccess(
	wg *sync.WaitGroup,
	channel <-chan PipelineSuccess,
	results *Results,
) {
	defer wg.Done()

	for message := range channel {
		dispatch.debugLogger.Printf("[%s] %s", message.PipelineName, message.Message)

		results.SuccessfulPipelines = append(results.SuccessfulPipelines, message)
	}
}
