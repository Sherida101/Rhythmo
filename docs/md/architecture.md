# Rhythmo — Build better habits with rhythm

📦 **File:** [architecture.md](https://github.com/Sherida101/Rhythmo/docs/md/architecture.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

# 🗂️ Project Structure
This file outlines the structure of the Rhythmo project, including the directory layout and the purpose of each component.

```markdown
rhythmo/
├── .github/                            ← GitHub Actions workflows
│   └── workflows/
│       └── build.yml                   ← CI workflow
├── cli/                                ← Go CLI tool
|   ├── api/                            ← Handles HTTP calls to Java backend
│   |   └── habits.go
|   ├──cmd/                             ← Cobra CLI commands
│   |  ├── add.go
│   |  ├── list.go
│   |  └── root.go
|   ├── model/                          ← Shared Go structs
│   |   └── habit.go
|   ├── main.go                         ← Main entrypoint
|   └── go.mod                          ← Go module file
├── config/
│   └── env.go                          ← for loading env vars
├── docs/                               ← Documentation
│   ├── api.md                          ← API documentation
│   ├── architecture.md                 ← Architecture overview
│   ├── cli.md                          ← CLI documentation
│   ├── full-stack-walkthrough.md       ← Full-stack walkthrough
│   ├── go-setup.md                     ← Go setup instructions
│   └── java-setup.md                   ← Java setup instructions
│
├── db/                                 ← Database files
│   ├── init.sql                        ← SQL script to initialise the database
│   └── schema.sql                      ← SQL script to create the database schema
├── web/                                ← Java Spring Boot web app
│   ├── src/
│   │   ├── main/
│   │   │   ├── java/                   ← Java source code
│   │   │   └── resources/              ← Static files (HTML, CSS, JS)
│   │   └── test/                       ← Java test code
│   └── pom.xml                         ← Maven build file
├── .gitignore                          ← Git ignore file
├── .env                                ← Environment variables
├── CONTRIBUTING.md                     ← Contribution guidelines
├── LICENSE                             ← License information
├── README.md                           ← Project overview
```

## Java web package layout

```markdown
web/src/main/java/com/rhythmo/app/
├── controller/            ← MVC & REST controllers
│   └── HabitApiController.java
├── model/                 ← Data entities
│   └── Habit.java
├── repository/            ← JPA repository interfaces
│   └── HabitRepository.java
├── service/               ← Business logic
├── RhythmoApplication.java ← Main entrypoint and application class
```


| Type             |      File                 |      Path                   |
|------------------|:-------------------------:|:---------------------------:|
| Entity/Model     | `Habit.java`              | com.rhythmo.app.model       |
| Repository       | `HabitRepository.java`    | com.rhythmo.app.repository  |
| REST Controller  | `HabitApiController.java` | com.rhythmo.app.controller  |
