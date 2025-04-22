# Rhythmo

## 📚 Table of Contents

- [🔍 Overview](#overview)
- [🧠 Objectives](#objectives)
- [🛠️ Prerequisites](#prerequisites)
- [🚀 Setup: Your First Go + Java Project](#setup-your-first-go--java-project)
- [▶️ Running the App](#running-the-app)
- [🗂️ Project Structure](#project-structure)
- [🧠 Want to Build Project from Scratch?](#want-to-build-project-from-scratch)
- [🤝 Contributing & License](#contributing--license)
- [🙌 Support](#support)
- [🔗 Connect with Me](#connect-with-me)

## Overview

Rhythmo is a Full-stack habit tracker that helps you build and maintain good habits. It features a command-line interface (CLI) built with Go, a web application powered by Java Spring Boot and PostgreSQL, and a REST API for seamless integration between the two.

The CLI allows you to manage your habits directly from the terminal, while the web app provides a user-friendly interface for tracking your progress.

## Objectives

This project serves as a learning resource for developers looking to explore the following:

|     |                                                 |
| --- | :---------------------------------------------: |
| ✅  |        Writing clean, idiomatic Go code         |
| ✅  |    Building a full REST API with Spring Boot    |
| ✅  | Designing a dual-mode application (CLI and web) |
| ✅  |       Structuring multi-language projects       |
| ✅  |            CI/CD with GitHub Actions            |

## Prerequisites

- Go 1.24+
- Java 17+
- Maven
- PostgreSQL
- Docker

---

## Setup: Your First Go + Java Project

### 1. Clone the Repository

```bash
git clone https://github.com/Sherida101/Rhythmo.git
cd Rhythmo
```

### 2. Install Dependencies

#### Go CLI

```bash
cd cli
go mod tidy
```

#### Java Web App

```bash
cd web
./mvnw clean install
```

### 3. Set Up Database

- Create a PostgreSQL database named `rhythmo`
- Run the SQL scripts in `./scripts/init.sql` to set up the initial database schema.
- Update the database connection settings in `web/src/main/resources/application.properties` if needed.

---

## Running the App

### CLI

```bash
cd cli
go run main.go
```

See [CLI documentation](docs/cli.md) for details on how to use the CLI.

### 🌐 Web

```bash
cd web
./mvnw spring-boot:run
```

### Using VSCode

1. Open VS Code command palette: `Ctrl` + `Shift` + `P` → **Debug: Select and Start Debugging** or click the dropdown next to the green ▶ Run button (upper-left)

2. Pick: **Run Rhythmo Go CLI** or **Run Rhythmo Spring Boot Web App**

3. The app will start running, and you can access it at `http://localhost:8080` in your web browser.

4. To stop the server, press `Ctrl` + `C` in the terminal or click the red square stop button in the debug panel.

### Using Makefile

The project includes a `Makefile` for easy management of tasks. Here are some common commands:

```bash
# Build & run Java web app
make run-web

# Run Go CLI
make cli

# Build Go CLI binary
make build-cli

# Start everything with Docker
make docker-up

# Stop Docker services
make docker-down
```

## API Usage

The REST Client extension can be used to test the API endpoints directly from the VSCode editor. See [API documentation](docs/api.md) for details on how to use the REST API.

## Project Structure

See [architecture.md](docs/architecture.md) for a detailed overview of the project structure.

## Want to Build Project from Scratch?

Check out the guides:

- [📘 Setup a Go CLI Project](docs/go-setup.md)
- [📗 Create a Java Spring Boot REST API](docs/java-setup.md)
- [🐬 Setup Docker for  Ubuntu-based systems](docs/docker-setup.md)
- [🧱 Full-Stack Walkthrough](docs/full-stack-walkthrough.md)

## Contributing & License

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute to this project.

If you find any issues or have suggestions, feel free to open an issue.

The project is open-sourced under the [MIT License](LICENSE). You can use, modify, and distribute this project as long as you include the original license.

---

## Support

![developer](https://img.shields.io/badge/Developed%20By%20%3A-ASPTechnologies%20Incorporations-blue) | [!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://buymeacoffee.com/asptechinc)

If you like this project, please consider supporting the developer.
There are no advertisements nor in-app purchases.

Your support will help keep the project free and updated. Thank you!

Star ⭐ the repository if you like what you see 😉.

## Connect with Me

[<img align="center" alt="Sherida Providence | YouTube" width="28px" src="https://firebasestorage.googleapis.com/v0/b/web-johannesmilke.appspot.com/o/other%2Fsocial%2Fyoutube.png?alt=media" />](https://www.youtube.com/obAZ9eizOU77HaEoLn0jHA?sub_confirmation=1)&ensp;YouTube: [@Sherida Providence](https://www.youtube.com/obAZ9eizOU77HaEoLn0jHA?sub_confirmation=1 'YouTube Sherida Providence')

[<img align="center" alt="AaliyahProvidence | LinkedIn" width="28px" src="https://firebasestorage.googleapis.com/v0/b/web-johannesmilke.appspot.com/o/other%2Fsocial%2Flinkedin.png?alt=media" />](https://linkedin.com/in/aaliyah-providence-0355b321a/)&ensp;LinkedIn: [@Aaliyah Providence](https://linkedin.com/in/aaliyah-providence-0355b321a/ 'LinkedIn Aaliyah Providence')

[<img align="center" alt="Sherida101 | GitHub" width="28px" src="https://firebasestorage.googleapis.com/v0/b/web-johannesmilke.appspot.com/o/other%2Fsocial%2Fgithub.png?alt=media" />](https://github.com/Sherida101)&ensp;GitHub: [@Sherida101](https://github.com/Sherida101 'GitHub Sherida101')
