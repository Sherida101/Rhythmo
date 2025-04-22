# Purpose:
# This script cleans up the Go project by cleaning build caches
# formatting the Go code, checking for syntax errors and ensuring
# dependencies are tidy.

# Usage:
# chmod +x ./scripts/clean-go.sh
# ./scripts/clean-go.sh

#!/bin/bash

# Save the current directory
original_dir=$(pwd)

# Define the path to the log-errors.sh script
LOG_ERRORS_SCRIPT="./log-errors.sh"

# Function to log errors using log-errors.sh
log_error() {
  "$LOG_ERRORS_SCRIPT" "$1"
}

# Function to execute a command and log stderr if it fails
run_and_log() {
  local output
  output=$("$@" 2>&1)
  local status=$?
  if [ $status -ne 0 ]; then
    log_error "$output"
  fi
  return $status
}

# Navigate to the CLI directory
cd cli || exit

# Clean build caches
echo "Cleaning Go build caches..."
run_and_log go clean -cache -modcache -i -r

# Format Go code
echo "Formatting Go code..."
run_and_log go fmt ./...

# Check for syntax errors
echo "Checking Go code for syntax errors..."
syntax_errors=$(go vet ./... 2>&1)
if [ -n "$syntax_errors"  ]; then
  log_error "Syntax errors found in Go code:"
  log_error "$syntax_errors"
else
  echo "No syntax errors."
fi

# Ensure dependencies are tidy
echo "Tidying Go module dependencies..."
run_and_log go mod tidy

echo "Go project cleaned and validated."

# Return to the original directory
cd "$original_dir" || exit
