package prpipeline

type Options struct {
	// Path to the git repository
	Path string
	// Style is the style of PR title to enforce
	Style PRStyle
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
