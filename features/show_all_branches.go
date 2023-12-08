package features

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
)

func InitShowAllBranches(repository *git.Repository) {
	showAllBranches(repository)
}

func showAllBranches(repository *git.Repository) {
	branches, err := repository.Branches()
	if err != nil {
		fmt.Println(err)
	}
	showBranches(filterBranches(branches, "", "contains"))
}
