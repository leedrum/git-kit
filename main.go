package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leedrum/git-kit/features"
	"gopkg.in/src-d/go-git.v4"
)

func mainOptions(list_features []string, repositoryPath string) {
	PrintOptions(list_features)
	var selection int
	fmt.Print("\nYour selection: ")
	fmt.Scanf("%d", &selection)

	if repositoryPath == "" {
		path, _ := os.Getwd()
		repositoryPath = path
	}

	repository, err := git.PlainOpen(repositoryPath)
	if err != nil {
		log.Fatal(err)
	}

	switch selection {
	case 0:
		return
	case 1:
		repositoryPath = features.InitRepositoryPath()
	case 2:
		features.InitShowAllBranches(repository)
	case 3:
		features.InitShowFilterBranches(repository)
	case 4:
		features.InitDeleteBranchesBySelect(repository)
	default:
		fmt.Printf("\n`%v` is not in the list options\n---------------\n", selection)
	}

	mainOptions(list_features, repositoryPath)
}

func main() {
	list_features := []string{}
	list_features = append(list_features, "Quit")
	list_features = append(list_features, "Change repository path")
	list_features = append(list_features, "Show all branches")
	list_features = append(list_features, "Show branches by filter")
	list_features = append(list_features, "Delete branches")

	mainOptions(list_features, "")
}

func PrintOptions(list_features []string) {
	fmt.Println("------ Git Kit Options ------ ")
	for i, text := range list_features {
		fmt.Printf("\n[%d] %s", i, text)
	}
	fmt.Println("\n---------------------")
}
