package db

import (
	"app/cmd"
	"app/internal"
	"database/sql"
	"fmt"
	"os"
	"time"
)

const (
	waitDuration = 5 * time.Second
	maxRetries   = 10
)

type DatabaseConfig struct {
	User         string
	Password     string
	DatabaseName string
	SslMode      string
	Host         string
	Port         uint16
}

func createConnectionString(config *DatabaseConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User, config.Password,
		config.Host, config.Port,
		config.DatabaseName, config.SslMode)
}

func NewConfigFromContext(ctx *cmd.Context) *DatabaseConfig {
	sslMode := "disable"
	if ctx.Database.SslEnabled {
		sslMode = "require"
	}

	return &DatabaseConfig{
		User:         ctx.Database.User,
		Password:     ctx.Database.Pwd,
		DatabaseName: ctx.Database.Database,
		SslMode:      sslMode,
		Host:         ctx.Database.Host,
		Port:         ctx.Database.Port,
	}
}

// NewConnection defines a utility method for creating a new database connection
func NewConnection(config *DatabaseConfig) *sql.DB {
	cs := createConnectionString(config)
	db := tryCreateConnection(cs, 0)
	return db
}

func tryCreateConnection(connectionString string, try uint8) *sql.DB {
	logger := internal.NewDefaultLogger()

	if try > maxRetries {
		logger.Printf("Max retries reached, unable to connect to database")
		os.Exit(1)
	}

	logger.Printf("Trying to connect to database (try=%d)", try)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Printf("Failed opening connection, retrying in %d seconds: %v", waitDuration/time.Second, err)
		time.Sleep(waitDuration)
		return tryCreateConnection(connectionString, try+1)
	}

	if err = db.Ping(); err != nil {
		logger.Printf("Failed to ping database: %v", err)
		time.Sleep(waitDuration)
		return tryCreateConnection(connectionString, try+1)
	}

	logger.Printf("Successfully connected to database")
	return db
}
