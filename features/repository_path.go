package features

import "fmt"

func InitRepositoryPath() string {
	repositoryPath := ""
	fmt.Print("Please enter the path of your repository: ")
	fmt.Scanf("%s", &repositoryPath)

	return repositoryPath
}
