# GitHub Repo Serving HTTP Server

This project creates an HTTP server that serves files from a specified GitHub repository.

## Configurations
Use a `.env` file in the project root for configuration:

- `REPO_URL`: The GitHub repository URL to clone or pull from.
- `SUBFOLDER`: The local directory to store and serve the repository files.
- `PULL_FREQUENCY`: Frequency at which the repository is pulled for updates (e.g., `30s`).
- `BRANCH_NAME`: The branch name to pull from.

### Example `.env` File
```env
REPO_URL=git@github.com:username/repo.git
SUBFOLDER=./repo
PULL_FREQUENCY=30s
BRANCH_NAME=main
```
