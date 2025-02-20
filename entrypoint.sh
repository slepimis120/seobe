#!/bin/bash

set -e

echo "Updating GitHub Actions Importer..."
gh actions-importer update

echo "Running migrate..."
gh actions-importer migrate jenkins -o output/audit --source-url https://9e83-82-117-210-226.ngrok-free.app/job/panoptikon --target-url https://github.com/slepimis120/panoptikon

tail -f /dev/null