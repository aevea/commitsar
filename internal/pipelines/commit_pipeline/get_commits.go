package commitpipeline

import (
	history "github.com/aevea/git/v2"
	"github.com/aevea/integrations"
)

func (pipeline *CommitPipeline) getCommits(gitRepo *history.Git) ([]history.SimpleCommit, error) {
	currentBranch := integrations.GetCurrentRef()

	if currentBranch == "" {
		gitBranch, err := gitRepo.CurrentBranch()

		if err != nil {
			return nil, err
		}

		currentBranch = gitBranch.Name().String()
	}

	pipeline.debugLogger.Printf("current branch %s", currentBranch)

	upstreamBranch := integrations.FindCompareBranch()

	if upstreamBranch == "" {
		upstreamBranch = "master"
	}

	pipeline.debugLogger.Printf("upstream branch %s", upstreamBranch)

	currentBranchCommit, err := gitRepo.LatestCommitOnBranch(currentBranch)

	if err != nil {
		return nil, err
	}

	currentBranchCommits, err := gitRepo.CommitsOnBranchSimple(currentBranchCommit.Hash)

	if err != nil {
		return nil, err
	}

	upstreamBranchCommit, err := gitRepo.LatestCommitOnBranch(upstreamBranch)

	if err != nil {
		return nil, err
	}

	upstreamBranchCommits, err := gitRepo.CommitsOnBranchSimple(upstreamBranchCommit.Hash)

	if err != nil {
		return nil, err
	}

	var diffedCommits []history.SimpleCommit

	for _, commit := range currentBranchCommits {
		if !contains(upstreamBranchCommits, commit) {
			diffedCommits = append(diffedCommits, commit)
		}
	}

	return diffedCommits, nil
}

func contains(s []history.SimpleCommit, e history.SimpleCommit) bool {
	for _, a := range s {
		if a.Hash == e.Hash {
			return true
		}
	}
	return false
}
