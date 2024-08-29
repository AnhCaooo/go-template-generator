// Created by Anh Cao on 27.08.2024.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// STEP 1: ask for which path the new repo will be located. If not given, then give default name 'go-web-server' and create it in current working directory.
	projectName, err := askProjectName()
	if err != nil {
		log.Fatal(err)
	}

	// STEP 2: navigate to one upper level and create folder, then navigate to.
	err = createProjectThenNavigate(projectName)
	if err != nil {
		log.Fatal(err)
	}
	// STEP 3: do git init
	err = doGitInit()
	if err != nil {
		log.Fatal(err)
	}

	// STEP 4: ask for project path (GitHub) in order to initialize Go module. Good to provide some examples.
	goPath, err := askGoPath()
	if err != nil {
		log.Fatal(err)
	}

	// STEP 5: initialize go module
	err = initializeGoModule(goPath)
	if err != nil {
		log.Fatal(err)
	}

	// STEP 6: do copy all files from template and paste to current working directory
	// STEP 7: fix all errors that happens after copying
	// STEP 8: do git commit and push to (GitHub)

}

func askProjectName() (name string, err error) {
	fmt.Println("Enter your desired name:")
	fmt.Scanln(&name)
	// validate desired name
	return name, nil
}

func askGoPath() (path string, err error) {
	fmt.Println("Enter your project path:")
	fmt.Scanln(&path)
	// validate project path
	return path, nil
}

func createProjectThenNavigate(name string) (err error) {
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

func initializeGoModule(goPath string) (err error) {
	cmd := exec.Command("go", "mod", "init", goPath)

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing go module: %s", err.Error())
	}
	fmt.Println("Go module initialized successfully!")
	return nil
}

func doGitInit() (err error) {
	cmd := exec.Command("git", "init")

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing git repo: %s", err.Error())
	}

	fmt.Println("Git repository initialized successfully!")
	return nil
}
