package main

import (
	"fmt"
	"os/exec"
)

func doGitInit() (err error) {
	cmd := exec.Command("git", "init")

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing git repo: %s", err.Error())
	}

	fmt.Println("Git repository initialized successfully!")
	return nil
}
