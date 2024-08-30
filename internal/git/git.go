package git

import (
	"fmt"
	"os/exec"
)

// do `git init` operation
func DoInit() (err error) {
	cmd := exec.Command("git", "init")

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing git repo: %s", err.Error())
	}

	fmt.Println("Git repository initialized successfully!")
	return nil
}

// do `git commit` operation with comment 'first commit'
func DoCommit() (err error) {
	return nil
}

// do `git push` operation to GitHub repository
func DoPush(repo string) (err error) {
	return nil
}
