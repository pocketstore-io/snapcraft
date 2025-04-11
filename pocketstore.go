package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Step 1: Check if Docker is installed
	if !isCommandAvailable("docker") {
		fmt.Println("Docker is not installed. Please install Docker and try again.")
		return
	}

	// Step 2: Check if Docker Compose is installed
	if !isCommandAvailable("docker compose") {
		fmt.Println("Docker Compose is not installed. Please install Docker Compose and try again.")
		return
	}

	// Step 3: Check if Git is installed
	if !isCommandAvailable("git") {
		fmt.Println("Git is not installed. Please install Git and try again.")
		return
	}

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

	fmt.Printf("Cloning repository: %s\n", repoURL)
	cmd := exec.Command("git", "clone", repoURL, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to clone repository: %v\n", err)
		os.Exit(1)
	}

	// Step 4: Run `docker compose up` in the target directory
	fmt.Println("Running 'docker compose up' in the target directory...")
	cmd = exec.Command("docker", "compose", "up")
	cmd.Dir = targetDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run 'docker compose up': %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Docker Compose is up and running!")
}

// Helper function to check if a command is available in the system
func isCommandAvailable(command string) bool {
	cmd := exec.Command("sh", "-c", "command -v "+command)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}