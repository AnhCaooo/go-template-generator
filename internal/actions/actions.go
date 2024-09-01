package actions

import (
	"fmt"
	"os"
	"os/exec"

	cp "github.com/otiai10/copy"
)

// create a new project with given path in one upper level of current directory then navigate to it
func CreateProjectThenNavigate(desiredPath string) (err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory. Error: %s", err.Error())
	}

	// template directory
	templateDir := currentDir + "/template"
	// get directory one level up
	parentDir := currentDir + "/.."

	// Create the new folder
	targetDir := parentDir + "/" + desiredPath
	err = os.Mkdir(targetDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error while creating folder: %s", err.Error())
	}

	err = cp.Copy(templateDir, targetDir)
	if err != nil {
		return fmt.Errorf("error while copying: %s", err.Error())
	}

	// Navigate to the newly created folder
	err = os.Chdir(targetDir)
	if err != nil {
		return fmt.Errorf("error navigating to new folder: %s", err.Error())
	}

	fmt.Printf("New project '%s' created and navigated to in '%s'\n", desiredPath, parentDir)
	return nil
}

// initialize Go module with given project path
func InitializeGoModule(goModulePath string) (err error) {
	// Check if module is already initialized
	if _, err := exec.Command("go", "mod", "init").CombinedOutput(); err == nil {
		fmt.Println("Module is already initialized.")
		return nil
	}

	// Create a new module
	cmd := exec.Command("go", "mod", "init", goModulePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize module: %w", err)
	}

	fmt.Printf("Module '%s' initialized.\n", goModulePath)
	return nil
}
