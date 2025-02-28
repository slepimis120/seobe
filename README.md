# Seobe

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Jenkins](https://img.shields.io/badge/jenkins-%232C5263.svg?style=for-the-badge&logo=jenkins&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

***
Tool that migrates a repo from Jenkinsfile to Github Actions

## Requirements
- Docker

***

## Setup

1. Fill in the environment variables in the .env.local file based on the .example.env.local file.

2. Start the Docker instance:
   ```sh
   docker-compose up --build
   ```
***

## How it works
When you run docker-compose up --build, the following steps occur behind the scenes:

1. **Docker Image Build**
2. **GitHub CLI Setup**
   - The GitHub access token from the .env.local file is used to authenticate the GitHub CLI.
3. **GitHub Actions Importer**
   - This tool is used to facilitate the migration from Jenkins to GitHub Actions.
4. **Entrypoint Script**
   - This script is executed as the entry point of the container.
5. **Migration process**
   - The gh-actions-importer is run to start the migration process.
6. **Go script**
   - The go script is executed to implement certain features that gh actions-importer (such as caching for Java projects, setting up Java project version etc.).
7**Branch and Pull Request Creation**
   - A new branch is created in the repository.
   - A pull request is created to merge the changes.

***

## Additional Information
- [Why Github Actions Importer](docs/why-gh-actions-importer.md)
- [Why Manual YML Parsing](docs/why-manual-yml-parsing.md)