// Package main is the entry into the application
package main

import (
	"context"
	"fmt"

	"github.com/unhommequidort/go-rest-api-course/internal/db"
)

// Run - is going to be reponsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("error connecting to database: %w", err)
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		fmt.Println("error pinging database: %w", err)
		return err
	}
	fmt.Println("database connection successful")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
