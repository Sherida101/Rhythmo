package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/crypto/bcrypt"

	habitmodel "github.com/Sherida101/Rhythmo/cli/model"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect() *pgx.Conn {
	// Load environment variables from .env file
	_ = godotenv.Load(".env")

	// Construct the Data Source Name (DSN)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Establish a connection to the database
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil { // Log the error using the log-error.sh script
		logMessage := fmt.Sprintf("Database connection failed: %v", err)

		// Execute the log-errors.sh script with the error message
		cmd := exec.Command("bash", "scripts/log-errors.sh", logMessage)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if cmdErr := cmd.Run(); cmdErr != nil {
			log.Printf("Failed to execute log-errors.sh: %v\n", cmdErr)
		}

		// Terminate the application
		os.Exit(1)
	}
	return conn
}

// Create
// AddUser adds a new user to the database
func AddUser(username string, password string) error {
	conn := Connect()
	defer conn.Close(context.Background())

	// Check if the user already exists
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := conn.QueryRow(context.Background(), query, username).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if exists {
		return fmt.Errorf("user with username %s already exists", username)
	}

	// Check if the password is empty
	if password == "" {
		return fmt.Errorf("hashed password is empty")
	}

	// Hash the password if it exists
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Convert the hashed password to a string
	password = string(hashedPassword)

	// Save the hashed password to the database
	query = `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err = conn.Exec(context.Background(), query, username, password)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}
	return nil
}

// AddHabit adds a new habit to the database
func AddHabit(name string, frequency int) error {
	conn := Connect()
	defer conn.Close(context.Background())

	query := `INSERT INTO habits (name, frequency) VALUES ($1, $2)`
	_, err := conn.Exec(context.Background(), query, name, frequency)
	if err != nil {
		return fmt.Errorf("failed to add habit: %w", err)
	}
	return nil
}

// Get
// GetHabits retrieves all habits from the database
func GetHabits() ([]habitmodel.Habit, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	query := `SELECT id, name, frequency FROM habits`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve habits: %w", err)
	}
	defer rows.Close()

	var habits []habitmodel.Habit
	for rows.Next() {
		var habit habitmodel.Habit
		err = rows.Scan(&habit.ID, &habit.Name, &habit.Frequency)
		if err != nil {
			return nil, fmt.Errorf("failed to scan habit: %w", err)
		}
		habits = append(habits, habit)
	}
	return habits, nil
}

// GetHabit retrieves a specific habit by ID
func GetHabit(id string) (habitmodel.Habit, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	query := `SELECT id, name, frequency FROM habits WHERE id = $1`
	var habit habitmodel.Habit
	err := conn.QueryRow(context.Background(), query, id).Scan(&habit.ID, &habit.Name, &habit.Frequency)
	if err != nil {
		return habitmodel.Habit{}, fmt.Errorf("failed to retrieve habit: %w", err)
	}
	return habit, nil
}

// GetHabitByName retrieves a specific habit by name
func GetHabitByName(name string) (habitmodel.Habit, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	query := `SELECT id, name, frequency FROM habits WHERE name = $1`
	var habit habitmodel.Habit
	err := conn.QueryRow(context.Background(), query, name).Scan(&habit.ID, &habit.Name, &habit.Frequency)
	if err != nil {
		return habitmodel.Habit{}, fmt.Errorf("failed to retrieve habit: %w", err)
	}
	return habit, nil
}

// GetHabitByID retrieves a specific habit by ID
func GetHabitByID(id string) (habitmodel.Habit, error) {
	conn := Connect()
	defer conn.Close(context.Background())

	query := `SELECT id, name, frequency FROM habits WHERE id = $1`
	var habit habitmodel.Habit
	err := conn.QueryRow(context.Background(), query, id).Scan(&habit.ID, &habit.Name, &habit.Frequency)
	if err != nil {
		return habitmodel.Habit{}, fmt.Errorf("failed to retrieve habit: %w", err)
	}
	return habit, nil
}

// Update
// UpdateHabit updates the details of a specific habit
func UpdateHabit(id string, name string, frequency int) error {
	conn := Connect()
	defer conn.Close(context.Background())

	query := `UPDATE habits SET name = $1, frequency = $2 WHERE id = $3`
	_, err := conn.Exec(context.Background(), query, name, frequency, id)
	if err != nil {
		return fmt.Errorf("failed to update habit: %w", err)
	}
	return nil
}

// MarkHabitAsCompleted marks a habit as completed for a specific date
// and updates the lastDone date in the database
func MarkHabitAsCompleted(id string, date string) error {
	conn := Connect()
	defer conn.Close(context.Background())

	// Check if the habit exists
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM habits WHERE id = $1)`
	err := conn.QueryRow(context.Background(), query, id).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if habit exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("habit with ID %s does not exist", id)
	}

	// Update the habit status to "done"
	// and set the lastDone to the provided date
	// or the current date if not provided
	query = `UPDATE habits SET status = 'done', lastDone = $1 WHERE id = $2`
	_, err = conn.Exec(context.Background(), query, date, id)
	if err != nil {
		return fmt.Errorf("failed to mark habit as completed: %w", err)
	}
	return nil
}

// Delete
// DeleteHabit deletes a habit from the database
func DeleteHabit(id string) error {
	conn := Connect()
	defer conn.Close(context.Background())

	// Check if the habit exists
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM habits WHERE id = $1)`
	err := conn.QueryRow(context.Background(), query, id).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if habit exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("habit with ID %s does not exist", id)
	}

	// Delete the habit
	query = `DELETE FROM habits WHERE id = $1`
	_, err = conn.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete habit: %w", err)
	}
	return nil
}
