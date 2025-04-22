# Rhythmo — Build better habits with rhythm

📦 **File:** [java-setup.md](https://github.com/Sherida101/Rhythmo/docs/md/java-setup.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

## 📗 `java-setup.md`

📌  Please note that Linux Ubuntu/Debian, Vim, Curl and VSCode are used as the OS and IDE in this guide. You can use any OS or IDE of your choice, but the commands may vary slightly.

```markdown
# 📗 Create a Java Spring Boot REST API

### Perquisites
Follow the `[Go Setup](docs/go-setup.md)` guide to set up the Go CLI first.

### Install VSCode extensions

Tool	Extension
Java	Java Extension Pack (by Microsoft)
Maven	Maven for Java
Spring Boot	Spring Boot Extension Pack

## Install JDK & Maven
```bash
sudo apt update
sudo apt install openjdk-17-jdk

# Set JAVA_HOME environment variable
echo 'export JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64' >> .bashrc
source .bashrc

# Verify Java Installation
java -version
```

## Create Spring Boot project
```bash
# Navigate to the project directory
cd Rhythmo

# Create a new directory for the Java project
mkdir web && cd web

# Create a new Spring Boot project using curl
curl https://start.spring.io/starter.zip \
  -d dependencies=web,devtools,thymeleaf,data-jpa,h2 \
  -d name=rhythmo-web \
  -o web.zip && unzip web.zip && rm web.zip

# Create a basic controller. Copy/paste the code from `HomeController.java` below
# into the created `HomeController.java` file and save it.
[HomeController.java](web/src/main/java/com/rhythmo/app/controller/HomeController.java)

# Run the Web app/Test it
cd web
./mvnw spring-boot:run
```
## Errors and Solutions
### Error: `VSCode not recognising changes in the project structure immediately`
- **Cause:** This is a common issue with VSCode where it does not immediately recognize changes in the project structure, especially after creating new files or directories.

- **Solution:**  Restart the VSCode IDE to refresh the project structure.

1. Open the Command Palette (Ctrl+Shift+P or Cmd+Shift+P on macOS).
2. Type Java: Clean the Java language server workspace and select it.
3. Restart VSCode.

or run the following
```bash
mvn clean
mvn install
```
## Creating models, repositories, REST controller

## Running the API
