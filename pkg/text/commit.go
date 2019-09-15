package text

// Commit is a parsed commit message + hash of the commit
type Commit struct {
	Category string
	Scope    string
	Breaking bool
	Heading  string
	Body     string
	Footer   string
	Hash     [20]byte
}
