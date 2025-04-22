// Rhythmo — Build better habits with rhythmo
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sherida101/Rhythmo/cli/cmd"
	"github.com/Sherida101/Rhythmo/cli/config"
	"github.com/Sherida101/Rhythmo/cli/db"
)

func main() {
	config.LoadEnv()
	cmd.Execute()

	// Establish a connection to the database
	conn := db.Connect()
	defer conn.Close(context.Background())

	// Example query
	var greeting string
	err := conn.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)
}
