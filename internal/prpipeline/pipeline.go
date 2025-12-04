package prpipeline

type Options struct {
	// Path to the git repository
	Path string
	// Styles are the styles of PR title to enforce (can be multiple)
	Styles []PRStyle
	// Keys checks for required keys in the PR title
	Keys []string
}

// Pip
type Pipeline struct {
	options Options
}

func New(options Options) (*Pipeline, error) {
	if options.Path == "" {
		options.Path = "."
	}

	return &Pipeline{options: options}, nil
}

func (pipeline *Pipeline) Name() string {
	return "pr-pipeline"
}
