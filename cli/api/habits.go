// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Sherida101/Rhythmo/cli/db"
	habitmodel "github.com/Sherida101/Rhythmo/cli/model"
	"github.com/joho/godotenv"
)

// FetchHabits fetches habits from the API.
// It takes an optional apiURL parameter. If not provided, it will use the API_URL from the .env file.
func FetchHabits(apiURL string) ([]habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// If no API URL is set, try to get it from .env
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
	}

	if apiURL == "" {
		// Fallback to DB
		fmt.Println("No API URL found. Fetching habits from the database.")
		return db.GetHabits()
	}

	// Otherwise, fetch from the API
	resp, err := http.Get(apiURL + "/api/habits")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from API: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var habits []habitmodel.Habit
	err = json.Unmarshal(body, &habits)
	return habits, err
}

// FetchHabit fetches a single habit by ID from the API.
func FetchHabit(apiURL string, id int) (habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return habitmodel.Habit{}, fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/habits/%d", apiURL, id))
	if err != nil {
		return habitmodel.Habit{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var habit habitmodel.Habit
	err = json.Unmarshal(body, &habit)
	return habit, err
}

// CreateHabit creates a new habit using the API.
func CreateHabit(apiURL string, habit habitmodel.Habit) (habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return habitmodel.Habit{}, fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	body, err := json.Marshal(habit)
	if err != nil {
		return habitmodel.Habit{}, err
	}

	resp, err := http.Post(apiURL+"/api/habits", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return habitmodel.Habit{}, err
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	var newHabit habitmodel.Habit
	err = json.Unmarshal(body, &newHabit)
	return newHabit, err
}

// UpdateHabit updates an existing habit using the API.
func UpdateHabit(apiURL string, id int, habit habitmodel.Habit) (habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return habitmodel.Habit{}, fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	body, err := json.Marshal(habit)
	if err != nil {
		return habitmodel.Habit{}, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/habits/%d", apiURL, id), bytes.NewBuffer(body))
	if err != nil {
		return habitmodel.Habit{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return habitmodel.Habit{}, err
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	var updatedHabit habitmodel.Habit
	err = json.Unmarshal(body, &updatedHabit)
	return updatedHabit, err
}

// DeleteHabit deletes a habit using the API.
func DeleteHabit(apiURL string, id int) error {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/habits/%d", apiURL, id), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete habit: %s", resp.Status)
	}
	return nil
}

// FetchHabitByName fetches a single habit by name from the API.
func FetchHabitByName(apiURL string, name string) (habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return habitmodel.Habit{}, fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/habits/name/%s", apiURL, name))
	if err != nil {
		return habitmodel.Habit{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var habit habitmodel.Habit
	err = json.Unmarshal(body, &habit)
	return habit, err
}

// FetchHabitByCategory fetches a single habit by category from the API.
func FetchHabitByCategory(apiURL string, category string) ([]habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return nil, fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/habits/category/%s", apiURL, category))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var habits []habitmodel.Habit
	err = json.Unmarshal(body, &habits)
	return habits, err
}

// FetchHabitByStatus fetches a single habit by status from the API.
func FetchHabitByStatus(apiURL string, status string) ([]habitmodel.Habit, error) {
	// Load .env file if not already loaded
	_ = godotenv.Load()

	// Fallback to .env API_URL if not provided
	if apiURL == "" {
		apiURL = os.Getenv("API_URL")
		if apiURL == "" {
			return nil, fmt.Errorf("API URL not provided and missing in .env")
		}
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/habits/status/%s", apiURL, status))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var habits []habitmodel.Habit
	err = json.Unmarshal(body, &habits)
	return habits, err
}
