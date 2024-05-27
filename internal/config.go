// /internal/config.go
package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	RepoURL       string `json:"repo_url"`
	Dest          string `json:"Dest"`
	PullFrequency string `json:"pull_frequency"`
	BranchName    string `json:"branch_name"`
}

func NewConfig(configFile string) (*Config, error) {
	// Default configuration
	config := &Config{
		RepoURL:       "git@github.com:username/repo.git",
		Dest:          "./repo",
		PullFrequency: "30s",
		BranchName:    "main",
	}

	// Load from .env
	loadEnvConfig(config)

	// Load from config file if provided
	if configFile != "" {
		loadConfigFile(configFile, config)
	}

	// Override values with environment variables if they exist
	overrideWithEnv(config)

	return config, nil
}

func loadEnvConfig(config *Config) {
	config.RepoURL = getEnv("REPO_URL", config.RepoURL)
	config.Dest = getEnv("DEST", config.Dest)
	config.PullFrequency = getEnv("PULL_FREQUENCY", config.PullFrequency)
	config.BranchName = getEnv("BRANCH_NAME", config.BranchName)
}

func loadConfigFile(filename string, config *Config) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
		return
	}
	json.Unmarshal(data, config)
}

func overrideWithEnv(config *Config) {
	config.RepoURL = getEnv("REPO_URL", config.RepoURL)
	config.Dest = getEnv("DEST", config.Dest)
	config.PullFrequency = getEnv("PULL_FREQUENCY", config.PullFrequency)
	config.BranchName = getEnv("BRANCH_NAME", config.BranchName)
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
