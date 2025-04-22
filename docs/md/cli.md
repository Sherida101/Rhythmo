# Rhythmo — Build better habits with rhythm

📦 **File:** [cli.md](https://github.com/Sherida101/Rhythmo/docs/md/cli.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

## 🖥️ `docs/CLI.md`

```markdown
# ⚙️ Rhythmo CLI documentation

This is the command-line interface (CLI) tool written in Go to manage your daily habits from the terminal.

## 🚀 How to Run

```bash
cd cli
go run main.go


## 🚀 How to Build

```bash
go build -o rhythmo-cli
./rhythmo-cli
```

## 📦 Commands (Planned)

| Command           |      Description             |      Example             | Run                 |
|-------------------|:----------------------------:|:------------------------:|:-------------------:|
| `list`            | View all habits              | rhythmo list             | go run main.go list |
| `add "Read Book"` | Add a new habit              | rhythmo add "Read Book"  |                     |
| `done 3`          | Mark habit ID 3 as completed | rhythmo done 3           |                     |
| `delete 5`        | Delete habit with ID 5       | rhythmo delete 5         |                     |
| `sync`            | Sync with web API            | rhythmo sync             |                     |

## 🌐 Web API Integration

The CLI can synchronise with the Spring Boot REST API.

See [API documentation](api.md) for details.

## 🧪 Running Tests

### Go CLI
```bash
go test ./...
```

### Java backend
```bash
./mvnw test
```
