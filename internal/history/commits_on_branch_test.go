package history

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func createBranch(repo *git.Repository) {
	headRef, _ := repo.Head()

	ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())

	storerErr := repo.Storer.SetReference(ref)

	if storerErr != nil {
		log.Println(storerErr)
	}

	worktree, _ := repo.Worktree()

	checkoutErr := worktree.Checkout(&git.CheckoutOptions{
		Branch: ref.Name(),
	})

	if checkoutErr != nil {
		log.Println(checkoutErr)
	}

}

func TestCommitsOnBranch(t *testing.T) {
	repo := setupRepo()
	createCommit(repo, "test commit on master")
	createBranch(repo)
	createCommit(repo, "commit on new branch")

	commits, err := CommitsOnBranch(repo, "my-branch")

	assert.Equal(t, len(commits), 1)

	commit, commitErr := repo.CommitObject(commits[0])

	assert.NoError(t, commitErr)
	assert.Equal(t, commit.Message, "commit on new branch")
	assert.Equal(t, err, nil)

}
