#!/bin/bash

set -e

dockerd &

sleep 5

echo "Updating GitHub Actions Importer..."
gh actions-importer update

echo "Running migrate..."
gh actions-importer dry-run jenkins --output-dir output/audit --source-url "${JENKINS_INSTANCE_URL}/job/${JENKINS_JOB_NAME}" --enable-features actions/cache

echo "Creating new branch"
sha=$(curl -s -H "Authorization: token $GITHUB_ACCESS_TOKEN" \
    "https://api.github.com/repos/$OWNER/$REPO/branches/$MAIN_BRANCH" | jq -r '.commit.sha')

curl -X POST -H "Authorization: token $GITHUB_ACCESS_TOKEN" \
    -d "{\"ref\": \"refs/heads/seobe/jenkins2github\", \"sha\": \"$sha\"}" \
    "https://api.github.com/repos/$OWNER/$REPO/git/refs"

cd /app/output/audit
REPO_FOLDER=$(find . -type d -mindepth 1 -maxdepth 1 -name "*$REPO*")
YML_FILE=$(find "$REPO_FOLDER/.github/workflows/" -type f -name "*.yml" -print -quit)

echo "Running Go script with path to YAML file"
go run /app/seobe.go --path "$YML_FILE"

if [ -n "$YML_FILE" ]; then
    WORKFLOW_CONTENT=$(cat "$YML_FILE" | base64 | tr -d '\n')

    echo "Adding new .github file to GitHub repository"
    API_URL="https://api.github.com/repos/$OWNER/$REPO/contents/.github/workflows/$(basename $YML_FILE)"

    curl -X PUT -H "Authorization: token $GITHUB_ACCESS_TOKEN" \
        -d '{
                "message": "Add .github/workflows/'$(basename $YML_FILE)' file",
                "content": "'$WORKFLOW_CONTENT'",
                "branch": "seobe/jenkins2github"
            }' \
        "$API_URL"

    echo "Creating pull request"
    curl -X POST -H "Authorization: token $GITHUB_ACCESS_TOKEN" \
        -d '{
                "title": "Migrate Jenkins to GitHub Actions",
                "head": "seobe/jenkins2github",
                "base": "'$MAIN_BRANCH'",
                "body": "This pull request migrates the Jenkins pipeline to GitHub Actions. Brought to you by **slepimis120/seobe**"
            }' \
        "https://api.github.com/repos/$OWNER/$REPO/pulls"
else
    echo "No .yml file found in the .github/workflows directory."
    exit 1
fi