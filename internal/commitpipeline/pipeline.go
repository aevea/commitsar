package commitpipeline

type Pipeline struct {
	args    []string
	options Options
}

type Options struct {
	// Path to repository
	Path string
	// UpstreamBranch is the branch against which to check
	UpstreamBranch string
	// Limit will limit how far back to check on upstream branch.
	Limit int
	// AllCommits will check all the commits on the upstream branch. Regardless of Limit setting.
	AllCommits bool
	Strict     bool
	// RequiredScopes will check scope in commit message against list of required ones
	RequiredScopes []string
}

func New(options *Options, args ...string) (*Pipeline, error) {
	if options == nil {
		options = &Options{
			Path:           ".",
			UpstreamBranch: "master",
			Limit:          0,
			AllCommits:     false,
			Strict:         true,
		}
	}

	return &Pipeline{
		options: *options,
		args:    args,
	}, nil
}

func (Pipeline) Name() string {
	return "commit-pipeline"
}
