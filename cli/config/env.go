// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	_ = godotenv.Load()
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
