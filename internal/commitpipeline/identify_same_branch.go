package commitpipeline

import (
	"fmt"

	history "github.com/aevea/git/v4"
	"github.com/apex/log"
)

// IdentifySameBranch breaks up the reference names and tries to identify if the branches are the same
func IdentifySameBranch(branchA, branchB string, gitRepo *history.Git) (bool, error) {
	log.Debugf("Getting latest commit on branch '%s'...", branchA)
	commitBranchA, err := gitRepo.LatestCommitOnBranch(branchA)

	if err != nil {
		log.Errorf("Failed to get latest commit on branch '%s': %v", branchA, err)
		return false, fmt.Errorf("failed to get latest commit on branch '%s': %w", branchA, err)
	}

	log.Debugf("Branch '%s' latest commit hash: %s", branchA, commitBranchA.Hash)

	log.Debugf("Getting latest commit on upstream branch '%s'...", branchB)
	commitBranchB, err := gitRepo.LatestCommitOnBranch(branchB)

	if err != nil {
		log.Errorf("Failed to get latest commit on upstream branch '%s': %v", branchB, err)
		return false, fmt.Errorf("failed to get latest commit on upstream branch '%s' (ensure the branch exists and is fetched): %w", branchB, err)
	}

	log.Debugf("Branch '%s' latest commit hash: %s", branchB, commitBranchB.Hash)

	return commitBranchA.Hash == commitBranchB.Hash, nil
}
