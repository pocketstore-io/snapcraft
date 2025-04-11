package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the flags
	demoFlag := flag.Bool("demo", false, "Clone the demo repository")
	repoFlag := flag.String("repo", "", "The URL of the repository to clone (required if --demo is not used)")

	// Parse the flags
	flag.Parse()

	// Determine which repository to clone
	var repoURL string
	if *demoFlag {
		repoURL = "https://github.com/pocketstore-io/demo.git"
	} else if *repoFlag != "" {
		repoURL = *repoFlag
	} else {
		fmt.Println("Error: You must specify either --demo or --repo=<repository_url>")
		flag.Usage()
		os.Exit(1)
	}

	// Clone the repository
	fmt.Printf("Cloning repository: %s\n", repoURL)
	cmd := exec.Command("git", "clone", repoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to clone repository: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Repository cloned successfully.")
}