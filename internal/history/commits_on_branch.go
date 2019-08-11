package history

import (
	"errors"
	"log"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// ErrCommonCommitFound is used for identifying when the iterator has reached the common commit
var ErrCommonCommitFound = errors.New("common commit found")

// CommitsOnBranch iterates through all references and returns commit hashes on given branch
func CommitsOnBranch(
	repo *git.Repository,
	branchRef plumbing.ReferenceName,
	compareRef plumbing.ReferenceName,
) ([]plumbing.Hash, error) {

	var commits []plumbing.Hash
	var branchHash plumbing.Hash
	var compareHash plumbing.Hash

	// common ancestor between two branches based on MergeBase
	var commonHash plumbing.Hash

	allRefs, refErr := repo.References()

	if refErr != nil {
		return nil, refErr
	}

	defer allRefs.Close()

	refIterErr := allRefs.ForEach(func(ref *plumbing.Reference) error {
		log.Println(branchRef, ref.Name().String())
		if ref.Name() == branchRef {
			branchHash = ref.Hash()
		}

		if ref.Name() == compareRef {
			compareHash = ref.Hash()
		}

		return nil
	})

	if refIterErr != nil {
		return nil, refIterErr
	}

	branchCommit, _ := repo.CommitObject(branchHash)

	compareCommit, _ := repo.CommitObject(compareHash)

	diffCommits, _ := branchCommit.MergeBase(compareCommit)

	commonHash = diffCommits[len(diffCommits)-1].Hash

	commitIter, logErr := repo.Log(&git.LogOptions{
		From: branchCommit.Hash,
	})

	if logErr != nil {
		return nil, logErr
	}

	iterErr := commitIter.ForEach(func(commit *object.Commit) error {
		if commit.Hash == commonHash {
			return ErrCommonCommitFound
		}
		commits = append(commits, commit.Hash)
		return nil
	})

	if iterErr != nil {
		if iterErr == ErrCommonCommitFound {
			return commits, nil
		}
		return nil, iterErr
	}

	return commits, nil
}
