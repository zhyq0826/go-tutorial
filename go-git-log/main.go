package main

import (
	"flag"
	"fmt"
)

func scan(path string) {
	fmt.Printf("Found folders:\n\n")
	respositories := recursiveScanFolder(path)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, respositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

func stats(email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a folder to scan for git")
	flag.StringVar(&email, "email", "zhyq0826@gmail.com", "the email to scan")
	flag.Parse()
	if folder != "" {
		scan(folder)
		// return
	}

	stats(email)
}
