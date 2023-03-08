package cmd

import (
	"app/internal"
	"flag"
	"os"
	"strconv"
)

type Context struct {
	// ExecuteMigrations toggles the execution of the automatic SQL migrations
	ExecuteMigrations bool

	Database struct {
		User       string
		Pwd        string
		Database   string
		SslEnabled bool
		Host       string
		Port       uint16
	}
}

func GetContext() *Context {
	ctx := &Context{}
	parseFlags(ctx)
	parseEnv(ctx)
	return ctx
}

func parseFlags(ctx *Context) {
	flag.BoolVar(&ctx.ExecuteMigrations, "migrate", true, "Execute database migrations on startup")
	flag.Parse()
}

func parseEnv(ctx *Context) {
	logger := internal.NewDefaultLogger()
	ctx.Database.User = envOrDefault("DB_USER", "postgres")
	ctx.Database.Pwd = envOrDefault("DB_PASSWORD", "postgres")
	ctx.Database.Database = envOrDefault("DB_NAME", "booker")
	ctx.Database.SslEnabled = envOrDefault("DB_SSL", "false") == "true"
	ctx.Database.Host = envOrDefault("DB_HOST", "localhost")

	port, err := strconv.ParseUint(envOrDefault("DB_PORT", "5432"), 10, 16)
	if err != nil {
		logger.Printf("Unable to parse database port")
		os.Exit(1)
	}

	ctx.Database.Port = uint16(port)
}

func envOrDefault(envName string, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
