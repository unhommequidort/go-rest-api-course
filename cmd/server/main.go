// Package main is the entry into the application
package main

import (
	"context"
	"fmt"

	"github.com/unhommequidort/go-rest-api-course/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("error migrating database: %w", err)
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(), "1ce125f4-b827-4ef9-98df-6aa1fbb9bc60"))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
