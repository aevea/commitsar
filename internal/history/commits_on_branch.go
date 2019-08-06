package history

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// CommitsOnBranch iterates through all references and returns commit hashes on given branch
func CommitsOnBranch(repo *git.Repository, branchName string) ([]plumbing.Hash, error) {
	branchRef := plumbing.NewBranchReferenceName(branchName)

	var commits []plumbing.Hash

	allRefs, refErr := repo.References()

	if refErr != nil {
		return nil, refErr
	}

	defer allRefs.Close()

	refIterErr := allRefs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name() == branchRef {
			commits = append(commits, ref.Hash())
		}

		return nil
	})

	if refIterErr != nil {
		return nil, refIterErr
	}

	return commits, nil
}
