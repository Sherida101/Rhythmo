# Rhythmo — Build better habits with rhythm

📦 **File:** [full-stack-walkthrough.md](https://github.com/Sherida101/Rhythmo/docs/md/full-stack-walkthrough.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

## 🧱 `full-stack-walkthrough.md`

```markdown
## How the Go CLI talks to the Spring Boot API

The **Rhythmo** project consists of two main components: a **Go CLI** application and a **Spring Boot API**. These two components interact via Hypertext Transfer Protocol (HTTP) requests.

The Go Command-line interface (CLI) communicates with the Spring Boot Application Programming Interface (API) by making Representational State Transferful (RESTful) requests to the endpoints exposed by the Spring Boot server. These requests can be used to fetch, create, update or delete resources like "Habits" in the application.

### Go CLI
- The Go CLI uses the `pgx` package for PostgreSQL connections and the `net/http` package for making HTTP requests to the Spring Boot API.

- HTTP requests are sent to the Spring Boot API for actions such as creating or updating habits.

- The Go CLI can also send GET requests to retrieve habits based on specific parameters such as user ID or category.

### Spring Boot API
- The Spring Boot application exposes RESTful API endpoints, which can handle HTTP requests from the Go CLI.

- The API is built with Spring Boot, using controllers to manage routes and services to handle business logic.

- The API uses **JPA** and **PostgreSQL** to manage the Habit entities and data persistence.

The data exchange between the Go CLI and the Spring Boot API follows the **JSON format**, making it easy to send and receive structured data.

## How to run both sides

### 1. Running the Go CLI
Please ensure that you have **Go** installed on your system. To run the Go CLI, follow these steps:

- Clone the Go repository:
  ```bash
  git clone https://github.com/Sherida101/Rhytmo.git
  cd Rhythmo
```
Set up your environment by creating a .env file containing your PostgreSQL database credentials:

```bash
DB_USER=<your_db_user>
DB_PASSWORD=<your_db_password>
DB_HOST=<your_db_host>
DB_PORT=<your_db_port>
DB_NAME=<your_db_name>
```

**Install the necessary Go dependencies**

```bash
go mod tidy
```

**Run the Go CLI**

```bash
go run main.go
```

### 2. Running the Spring Boot API
Please Ensure that you have Java and Maven installed on your system. To run the Spring Boot API, follow these steps:

**Clone the Spring Boot repository**

```bash
git clone https://github.com/Sherida101/Rhythmo.git
cd Rhythmo
```

**Set up the PostgreSQL database connection in the application.properties or application.yml file**

```bash
spring.datasource.url=jdbc:postgresql://localhost:5432/<your_db_name>
spring.datasource.username=<your_db_user>
spring.datasource.password=<your_db_password>
```

**Build and run the Spring Boot application**

```bash
mvn clean install
mvn spring-boot:run
```

This will start the Spring Boot application on port 8080 by default. You can access the API at `http://localhost:8080`.

## How they are connected (URL, ports, formats)
- Go CLI communicates with Spring Boot API using HTTP requests.

- Spring Boot API URL: http://localhost:8080

    - Example endpoints:

        - GET /habits - Retrieve all habits

        - POST /habits - Create a new habit

        - PUT /habits/{id} - Update a habit

        - DELETE /habits/{id} - Delete a habit

- The Go CLI makes requests to these endpoints with data formatted as JSON.

- Both sides use PostgreSQL for persistent storage, ensuring the data is stored and shared between the two applications.

**Example of a Go CLI request**

```go
resp, err := http.Post("http://localhost:8080/habits", "application/json", bytes.NewBuffer(jsonData))
```
**Example of a Spring Boot API response**

```json
{
  "id": 1,
  "name": "Exercise",
  "category": "Health",
  "status": "Active",
  "createdAt": "2025-04-21T12:00:00Z"
}
```

## Screenshots or terminal output

1. **Example Output from Go CLI**
```bash
$ go run main.go
Successfully created habit: Exercise
Habit ID: 1
Status: Active
```

2. **Example Output from Spring Boot API**

```bash
$ mvn spring-boot:run
2025-04-21 12:00:00.000  INFO 12345 --- [           main] c.r.a.RhythmoApplication                : Started RhythmoApplication in 5.678 seconds (JVM running for 6.234)
2025-04-21 12:00:10.123  INFO 12345 --- [io-8080-exec-1] c.r.a.controller.HabitController        : Created habit with ID 1
```

This shows the Go CLI successfully communicating with the Spring Boot API and creating a new habit.
