package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for migration
	_ "github.com/lib/pq"
)

// MigrateDB - is going to be responsible for
// migrating our database
func (d *Database) MigrateDB() error {
	fmt.Println("Migrating the database...")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating postgres driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postfres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("error applying migration: %w", err)
		}
	}

	// error := m.Force(1) //1 is migrations version number, you may use your latest version
	// if err != nil {
	// 	return error
	// }

	fmt.Println("Migration successful")

	return nil
}
