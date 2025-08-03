#!/bin/bash

set -e

VERSION="v1.0.0"
APP_NAME="gatekeep"

echo "🔨 Building binaries..."

# Create bin directory
mkdir -p bin

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/${APP_NAME}.exe main.go
zip ${APP_NAME}-windows.zip bin/${APP_NAME}.exe LICENSE README.md

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o bin/${APP_NAME}-darwin main.go
zip ${APP_NAME}-macos.zip bin/${APP_NAME}-darwin LICENSE README.md

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bin/${APP_NAME}-linux main.go
zip ${APP_NAME}-linux.zip bin/${APP_NAME}-linux LICENSE README.md

echo "✅ Builds complete and zipped."

echo "📂 Files created:"
ls -1 *.zip

echo "🌍 Opening GitHub Releases page..."
REPO_URL=$(git config --get remote.origin.url | sed 's/\.git$//')
open "$REPO_URL/releases/new?tag=$VERSION&title=$VERSION" 2>/dev/null || xdg-open "$REPO_URL/releases/new?tag=$VERSION&title=$VERSION"

echo "🚀 Ready to upload zip files to the new release!"
