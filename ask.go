package main

import "fmt"

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
