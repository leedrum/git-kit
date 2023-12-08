package features

import (
	"fmt"

	constants "github.com/leedrum/git-kit/constants"
	"gopkg.in/src-d/go-git.v4"
)

func InitShowFilterBranches(repository *git.Repository) {
	getCondition()
	filterTxt := ""
	fmt.Println("Please enter keyword contains in branch name: ")
	fmt.Scanf("%s", &filterTxt)
	ShowBranchesFiltered(repository, filterTxt, constants.FilterTypeContains)
	fmt.Println("You selected `Delete branches by name`")
}

func ShowBranchesFiltered(repository *git.Repository, filterContent string, filterType string) {
	branches, err := repository.Branches()
	if err != nil {
		fmt.Println(err)
	}
	branchesFiltered := filterBranches(branches, filterContent, filterType)
	showBranches(branchesFiltered)
}

func getCondition() string {
	selection := 1
	fmt.Println("Please select filter type: ")
	fmt.Println("1. Contains")
	fmt.Println("2. Prefix")
	fmt.Print("Your selection: ")
	fmt.Scanf("%d", &selection)

	switch selection {
	case 1:
		return constants.FilterTypeContains
	case 2:
		return constants.FilterTypePrefix
	default:
		return constants.FilterTypeContains
	}
}
