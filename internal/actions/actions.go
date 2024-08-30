package actions

import (
	"fmt"
	"os"
	"os/exec"
)

// create a new project with given name in one upper level of current directory then navigate to it
func CreateProjectThenNavigate(name string) (err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory. Error: %s", err.Error())
	}

	// Navigate one level up
	parentDir := currentDir + "/.."

	// Create the new folder
	err = os.Mkdir(parentDir+"/"+name, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error while creating folder: %s", err.Error())
	}

	// Navigate to the newly created folder
	err = os.Chdir(parentDir + "/" + name)
	if err != nil {
		return fmt.Errorf("error navigating to new folder: %s", err.Error())
	}

	fmt.Printf("New project '%s' created and navigated to in '%s'\n", name, parentDir)
	return nil
}

// initialize Go module with given project path
func InitializeGoModule(goPath string) (err error) {
	cmd := exec.Command("go", "mod", "init", goPath)

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing go module: %s", err.Error())
	}
	fmt.Println("Go module initialized successfully!")
	return nil
}
