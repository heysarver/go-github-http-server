// /internal/git.go
package internal

import (
	"log"
	"os"
	"os/exec"
)

func CloneOrPullRepo(config *Config) error {
	if _, err := os.Stat(config.Subfolder); os.IsNotExist(err) {
		log.Printf("Cloning repository %s into %s", config.RepoURL, config.Subfolder)
		cmd := exec.Command("git", "clone", "-b", config.BranchName, config.RepoURL, config.Subfolder)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	log.Printf("Pulling latest changes from repository %s", config.RepoURL)
	cmd := exec.Command("git", "-C", config.Subfolder, "pull", "origin", config.BranchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
