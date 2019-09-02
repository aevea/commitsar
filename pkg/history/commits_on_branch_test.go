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

func createTestHistory(repo *git.Repository) {
	createCommit(repo, "test commit on master")
	createBranch(repo)
	createCommit(repo, "commit on new branch")
	createCommit(repo, "second commit on new branch")
	createCommit(repo, "third commit on new branch")
}

func TestCommitsOnBranch(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	headRef, _ := repo.Head()

	commits, err := CommitsOnBranch(repo, headRef.Hash(), "master")

	assert.Equal(t, 3, len(commits))

	commit, commitErr := repo.CommitObject(commits[0])

	assert.NoError(t, commitErr)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.Equal(t, err, nil)

}

func TestCommitsOnBranchWithMasterMerge(t *testing.T) {
	repo, _ := git.PlainOpen("../../testdata/commits_on_branch_test")
	headRef, _ := repo.Head()

	commits, err := CommitsOnBranch(repo, headRef.Hash(), "master")

	assert.Equal(t, 2, len(commits))

	lastCommit, lastCommitErr := repo.CommitObject(commits[0])
	assert.NoError(t, lastCommitErr)
	assert.Equal(t, "Merge branch 'master' into behind-master\n", lastCommit.Message)

	penultimateCommit, penultimateCommitErr := repo.CommitObject(commits[1])
	assert.NoError(t, penultimateCommitErr)
	assert.Equal(t, "first commit on behind-master branch\n", penultimateCommit.Message)

	assert.Equal(t, err, nil)

}
