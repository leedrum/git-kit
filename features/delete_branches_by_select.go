package features

import (
	"fmt"
	"log"

	componentsUI "github.com/leedrum/git-kit/components"
	utils "github.com/leedrum/git-kit/utils"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func InitDeleteBranchesBySelect(repository *git.Repository) {
	showUIDeleteBranchesBySelect(repository)
}

func showUIDeleteBranchesBySelect(repository *git.Repository) {
	branchNames, err := utils.PrintBranches(repository)
	if err != nil {
		log.Fatal(err)
		return
	}

	branchNamesDelete := componentsUI.HandleBranchSelection(branchNames)
	branches, err := repository.Branches()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = branches.ForEach(func(reference *plumbing.Reference) error {
		shortNameBranch := reference.Name().Short()
		nameBranch := string(reference.Name())
		if utils.IsSliceContain(branchNamesDelete, shortNameBranch) {
			return utils.DeleteBranch(repository, nameBranch)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("===== Delete successfully =====")
		fmt.Println("Deleted branches: ", branchNamesDelete)
		fmt.Println("================================")
	}
}
