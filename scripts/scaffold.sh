# To run this script, use the command:
# ./scripts/scaffold.sh [go|java|md] [optional: filename]

# Purpose:
# This script adds a header to files in the Rhythmo project.
# It checks if the header already exists and only adds it if it doesn't.
# The script is designed to be run from the root of the project directory.

# Usage:
# Add headers to all files of a type:
#   ./scripts/scaffold.sh go
#   ./scripts/scaffold.sh java
#   ./scripts/scaffold.sh md

# Add headers to a specific file:
#   ./scripts/scaffold.sh go path/to/file.go
#   ./scripts/scaffold.sh java path/to/file.java
#   ./scripts/scaffold.sh md path/to/file.md
#!/bin/bash

# === CONFIG ===
REPO_URL="https://github.com/Sherida101/Rhythmo"

# === HEADER GENERATORS ===

go_header() {
cat <<EOF
// Rhythmo — Build better habits with rhythm
// $REPO_URL
//
// See LICENSE for copyright details.

EOF
}

java_header() {
cat <<EOF
/*
 Rhythmo — Build better habits with rhythm
 $REPO_URL

 See LICENSE for copyright details.
*/
EOF
}

md_header() {
cat <<EOF
# Rhythmo — Build better habits with rhythm

📦 **File:** [$FILENAME]($REPO_URL/$FILE_DIRNAME/$FILENAME)

🔗 Repo: [$REPO_URL]($REPO_URL)

---
EOF
}

# === APPLY HEADER TO FILE ===

apply_header() {
  local file=$1
  local extension=$2
  local current_header

  case "$extension" in
    go)
      current_header="// Rhythmo"
      ;;
    java)
      current_header="Rhythmo — Build"
      ;;
    md)
      current_header="# Rhythmo"
      ;;
  esac

  if grep -q "$current_header" "$file"; then
    echo "✔️  Header already exists in $file"
  else
    case "$extension" in
      go)
        { go_header; cat "$file"; } > temp && mv temp "$file"
        ;;
      java)
        { java_header; echo ""; cat "$file"; } > temp && mv temp "$file"
        ;;
      md)
        FILENAME=$(basename "$file")
        FILE_DIRNAME=$(dirname "$file")
        eval "md_header" > temp
        echo "" >> temp
        cat "$file" >> temp
        mv temp "$file"
        ;;
    esac
    echo "✅ Header added to $file"
  fi
}

# === MAIN ENTRY ===

TYPE=$1
FILE=$2

if [ -z "$TYPE" ]; then
  echo "Usage: ./scripts/scaffold.sh [go|java|md] [optional: filename]"
  exit 1
fi

# One file
if [ -n "$FILE" ]; then
  apply_header "$FILE" "$TYPE"
else
  # All matching files
  echo "🔍 Scanning for *.$TYPE files..."
  find . -type f -name "*.$TYPE" | while read -r file; do
    apply_header "$file" "$TYPE"
  done
fi
