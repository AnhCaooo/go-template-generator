// Created by Anh Cao on 27.08.2024.

package main

import (
	"log"

	"github.com/AnhCaooo/go-web-server-template-generator/internal/actions"
	"github.com/AnhCaooo/go-web-server-template-generator/internal/git"
	"github.com/AnhCaooo/go-web-server-template-generator/internal/input"
)

func main() {
	// STEP 1: ask for which path the new repo will be located. If not given, then give default name 'go-web-server' and create it in current working directory.
	projectName, err := input.AskProjectName()
	if err != nil {
		log.Fatal(err)
	}

	// STEP 2: navigate to one upper level and create folder, then navigate to.
	err = actions.CreateProjectThenNavigate(projectName)
	if err != nil {
		log.Fatal(err)
	}
	// STEP 3: do git init
	err = git.DoInit()
	if err != nil {
		log.Fatal(err)
	}

	// STEP 4: ask for project path (GitHub) in order to initialize Go module. Good to provide some examples.
	goPath, err := input.AskGoPath()
	if err != nil {
		log.Fatal(err)
	}

	// STEP 5: initialize go module
	err = actions.InitializeGoModule(goPath)
	if err != nil {
		log.Fatal(err)
	}

	// STEP 6: do copy all files from template and paste to current working directory
	// STEP 7: fix all errors that happens after copying
	// STEP 8: do git commit and push to (GitHub)

}
