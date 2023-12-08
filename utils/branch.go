package utils

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func DeleteBranch(repository *git.Repository, nameBranch string) error {
	errDelete := repository.Storer.RemoveReference(plumbing.ReferenceName(nameBranch))

	if errDelete != nil {
		err_str := fmt.Sprintf("errDelete of branch `%s`", nameBranch)
		fmt.Println(err_str)
	}

	return nil
}

func PrintBranches(repository *git.Repository) ([]string, error) {
	branches, err := repository.Branches()
	if err != nil {
		return nil, err
	}

	branchNames := make([]string, 0)

	err = branches.ForEach(func(reference *plumbing.Reference) error {
		branchNames = append(branchNames, reference.Name().Short())
		return nil
	})

	if err != nil {
		return nil, err
	}

	return branchNames, nil
}
