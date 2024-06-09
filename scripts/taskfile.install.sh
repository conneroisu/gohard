#!/bin/bash
# file: taskfile.install.sh
# url: https://github.com/conneroisu/hardgo/scripts/taskfile.dev.requirements.sh
# title: Installing Development Requirements
# description: This script installs the required development tools for the project.

# Check if the command, brew, exists, if not install it
command -v brew >/dev/null 2>&1 || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Check if the command, go, exists, if not install it
command -v go >/dev/null 2>&1 || brew install go

# Check if the command, gum, exists, if not install it
command -v gum >/dev/null 2>&1 || go install github.com/charmbracelet/gum@latest

# Check if the command, sqlite3, exists, if not install it
command -v sqlite3 >/dev/null 2>&1 || gum spin --spinner dot --title "Installing SQLite3" --show-output -- \
    brew install sqlite
