package root_runner

type Runner struct {
}

type RunnerOptions struct {
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

// New returns a new instance of a RootRunner with fallback for logging
func New() *Runner {

	return &Runner{}
}
