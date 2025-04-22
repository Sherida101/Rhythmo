# Rhythmo — Build better habits with rhythm

📦 **File:** [docker-setup.md](https://github.com/Sherida101/Rhythmo/docs/md/docker-setup.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

## 🐬 `docker-setup.md`

Follow the steps below to set up Docker for **Rhythmo** project. This guide assumes that you are using an Ubuntu Operating system (OS).

If not, please refer to the [Docker installation guide](https://docs.docker.com/get-docker/) for your operating system.

### Install Docker
```markdown
sudo apt update
sudo apt install -y ca-certificates curl gnupg lsb-release

sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
  sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) \
  signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io

# Run Docker without sudo
sudo usermod -aG docker $USER
newgrp docker

# Enable Docker to start on boot
sudo systemctl enable docker
sudo systemctl status docker
sudo systemctl start docker

# Optional: Verify Docker installation
docker --version
docker info

# Optional: Test Docker installation
docker run hello-world

# Optional: Restart Docker service
sudo systemctl restart docker

# Optional: Install Docker compose plugin (Version 2 of Docker Compose)
# Note: Docker Compose is now included as a plugin in Docker Desktop and Docker Engine.
sudo apt-get install docker-compose-plugin

# Optional: Verify Docker compose plugin installation
docker compose version

# Optional: Test Docker service
docker ps
```

### Install Docker Compose
```markdown
sudo apt update
sudo apt install -y docker-compose

# Verify Docker Compose installation
docker compose version
```
