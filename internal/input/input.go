package input

import "fmt"

// ask for project name and collect value
func AskProjectName() (name string, err error) {
	fmt.Println("Enter your desired name:")
	fmt.Scanln(&name)
	// validate desired name
	if len(name) == 0 {
		name = "go-web-server"
	}
	return name, nil
}

// ask for Go path and collect value
func AskGoPath() (path string, err error) {
	fmt.Println("Enter your project path:")
	fmt.Scanln(&path)
	// validate project path
	return path, nil
}
