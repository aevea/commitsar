package cmd

import "strings"

// IdentifySameBranch breaks up the reference names and tries to identify if the branches are the same
func IdentifySameBranch(branchA, branchB string) bool {
	splitBranchA := strings.Split(branchA, "/")

	splitBranchB := strings.Split(branchB, "/")

	return splitBranchA[(len(splitBranchA)-1)] == splitBranchB[len(splitBranchB)-1]
}
