package features

import (
	"fmt"
	"strings"

	constants "github.com/leedrum/git-kit/constants"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

func filterBranches(branches storer.ReferenceIter, filterContent string, filterType string) []plumbing.Reference {
	branchesFiltered := make([]plumbing.Reference, 0)
	branches.ForEach(func(branch *plumbing.Reference) error {
		if filterType == constants.FilterTypeContains {
			if strings.Contains(branch.Name().Short(), filterContent) {
				branchesFiltered = append(branchesFiltered, *branch)
			}
		} else if filterType == constants.FilterTypePrefix {
			if strings.HasPrefix(branch.Name().Short(), filterContent) {
				branchesFiltered = append(branchesFiltered, *branch)
			}
		}

		return nil
	})
	return branchesFiltered
}

func showBranches(branches []plumbing.Reference) {
	for _, branch := range branches {
		bName := branch.Name().Short()
		fmt.Println(bName)
	}
}
