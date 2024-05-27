// /main.go
package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go_github_http_server/internal"
)

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Continuing without it.")
	}

	// Command-line flags
	configFile := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	// Initialize configuration
	config, err := internal.NewConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// Git repository setup and initial pull
	if err := internal.CloneOrPullRepo(config); err != nil {
		log.Fatalf("Failed to setup repo: %v", err)
	}

	// Initialize the Gin router
	router := gin.Default()

	// Serve static files from the repository folder
	path, err := filepath.Abs(config.Subfolder)
	if err != nil {
		log.Fatalf("Invalid subfolder path: %v", err)
	}
	router.Static("/", path)

	// Set up server in a goroutine to listen in the background
	go func() {
		if err := router.Run(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Start auto-pull function with user-defined frequency
	autoPullInterval, err := time.ParseDuration(config.PullFrequency)
	if err != nil {
		log.Fatalf("Invalid pull frequency: %v", err)
	}
	ticker := time.NewTicker(autoPullInterval)
	go func() {
		for range ticker.C {
			if err := internal.CloneOrPullRepo(config); err != nil {
				log.Printf("Failed to pull repo: %v", err)
			}
		}
	}()

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down server...")
	ticker.Stop()
}
