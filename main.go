package main

import (
	"app/cmd"
	"app/db"
	"app/internal"
	"app/internal/http"
	"embed"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

//go:embed web/dist/*
var files embed.FS

func main() {
	logger := internal.NewDefaultLogger()

	ctx := cmd.GetContext()
	conn := db.NewConnection(db.NewConfigFromContext(ctx))

	if ctx.ExecuteMigrations {
		logger.Printf("Migrations enabled, executing migration scripts")
		if err := db.ApplyMigrations(conn); err != nil {
			logger.Printf("Failed executing migrations: %v", err)
			os.Exit(1)
		}
	}

	s := http.NewHttpRouter(&files, conn)
	s.Start(3000)

}
