package dispatcher

// Pipeliner interface describes the requirements for pipelines the dispatcher can run
type Pipeliner interface {
	Name() string
	Run() (*PipelineSuccess, error)
}

// FailureData is used to build the table output
// Example: "hash": "somehash". It is built as a struct to enforce order
type FailureData struct {
	Name  string
	Value string
}

// PipelineSuccess is returned by a pipeline if it encounters no errors
type PipelineSuccess struct {
	PipelineName string
	Message      string
}

// PipelineError is used to group errors based on PipelineName
type PipelineError struct {
	PipelineName string
	Data         []FailureData
	Error        error
}
