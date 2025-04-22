#!/bin/bash
# Purpose:
# This script cleans up the Go project by cleaning compiled classes,
# compiling Java files and running tests.

# Usage:
# chmod +x ./scripts/clean-java.sh
# ./scripts/clean-java.sh

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

# Navigate to the web directory
cd web || exit

# Clean compiled classes
echo "Cleaning compiled Java classes..."
find . -name "*.class" -exec rm -f {} \;

# Compile Java files
echo "Compiling Java files..."
javac $(find . -name "*.java")

# Run tests (if you have a test class named TestRunner)
echo "Running tests..."
java TestRunner

echo "Java project cleaned and validated."

# Return to the original directory
cd "$original_dir" || exit
