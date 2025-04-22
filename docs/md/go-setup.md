# Rhythmo — Build better habits with rhythm

📦 **File:** [go-setup.md](https://github.com/Sherida101/Rhythmo/docs/md/go-setup.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

## 📘 `go-setup.md`

📌  Please note that Linux Ubuntu/Debian, Vim and VSCode are used as the OS and IDE in this guide. You can use any OS or IDE of your choice, but the commands may vary slightly.

```markdown
## 📘 Setup a Go CLI Project

### Install VSCode extensions

Tool	Extension
Go	Go (by Go Team at Google)
REST Client (for testing APIs inside VS Code)

### Install Go
Download the Latest Go Version at https://go.dev/dl/. Replace `1.24.2` with the latest version.

```bash
# Download Go
wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz

# Remove Old version of Go (if installed)
sudo rm -rf /usr/local/go

# Extract and Install Go
sudo tar -C /usr/local -xzf go1.24.2.linux-amd64.tar.gz

# Add Go to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> .bashrc
source .bashrc

# Verify Go Installation
go version

```

## Create Go project and initialise it with `go mod`

```bash
# Create a new directory for the project
mkdir Rhythmo && cd Rhythmo

# Create a new Go module. replace `github.com/yourusername/Rhythmo` with your own module name or repository URL
mkdir cli && cd cli
go mod init github.com/yourusername/Rhythmo/cli
touch main.go

# Copy/paste the code from `main.go` below
# into the created `main.go` file and save it.
[main.go](cli/main.go)

# Run the Go CLI/Test it
cd cli
go run main.go
```

## cobra-cli usage
```bash
# Install npm

# Navigate to the cli directory
cd cli
# Install cobra-cli
go get -u github.com/spf13/cobra@latest
cobra-cli init
```
## Writing first command
```bash
# Create a new command
cobra-cli add add
cobra-cli add list
```
```bash

## Calling REST APIs
