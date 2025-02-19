# Seobe

![Jenkins](https://img.shields.io/badge/jenkins-%232C5263.svg?style=for-the-badge&logo=jenkins&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
***
Tool that migrates repo from Jenkinsfile to Github Actions

## Requirements
- Docker

## Setup

1. Fill in the environment variables in the .env.local file based on the .example.env.local file.

2. Start the Docker instance:
   ```sh
   docker-compose up --build
   ```