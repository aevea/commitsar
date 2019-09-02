package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
)

func TestBranchDiffCommits(t *testing.T) {
	repo := setupRepo()
	createTestHistory(repo)

	testGit := &Git{repo: repo}

	commits, err := testGit.BranchDiffCommits("my-branch", "master")

	commit, _ := repo.CommitObject(commits[0])

	assert.NoError(t, err)
	assert.Equal(t, "third commit on new branch", commit.Message)
	assert.Equal(t, 3, len(commits))
}

func TestBranchDiffCommitsWithMasterMerge(t *testing.T) {
	repo, _ := git.PlainOpen("../../testdata/commits_on_branch_test")
	testGit := &Git{repo: repo, Debug: true}

	commits, err := testGit.BranchDiffCommits("behind-master", "master")

	assert.Equal(t, 2, len(commits))

	commit, _ := repo.CommitObject(commits[1])

	assert.Equal(t, "first commit on branch\n", commit.Message)

	assert.Equal(t, err, nil)

}
