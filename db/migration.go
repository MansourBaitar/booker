package db

import (
	"app/internal"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

const migrationSource = "file://db/migrations"

// ApplyMigrations defines a function which will apply the database migrations
func ApplyMigrations(db *sql.DB) error {
	logger := internal.NewDefaultLogger()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationSource, "postgres", driver)
	if err != nil {
		return err
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return err
	}

	logger.Printf("Current migration (version=%d, dirty=%t)", version, dirty)
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	if err == migrate.ErrNoChange {
		logger.Printf("Latest migration already applied (version=%d)", version)
	} else {
		version, _, _ = m.Version()
		logger.Printf("Applied migration version=%d", version)
	}

	return nil
}
