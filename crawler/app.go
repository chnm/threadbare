package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/ratelimit"
)

type Config struct {
	dbstr string
}

type App struct {
	DB       *sql.DB
	Config   *Config
	Client   *http.Client
	Limiters struct {
		Items ratelimit.Limiter
	}
}

func (app *App) Init() error {
	log.Info("Starting the API crawler")

	app.Config = &Config{}

	dbstr, ok := os.LookupEnv("THREADBARE_DB_STR")
	if !ok {
		return errors.New("THREADBARE_DB_STR environment variable not set")
	}
	app.Config.dbstr = dbstr

	policy := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10)

	// Connect to the database
	var db *sql.DB
	dbConnect := func() error {
		d, err := sql.Open("pgx", os.Getenv("THREADBARE_DB_STR"))
		if err != nil {
			return fmt.Errorf("failed to dial the database: %w", err)
		}
		if err := d.Ping(); err != nil {
			return fmt.Errorf("failed to ping the database: %w", err)
		}

		db = d
		return nil
	}

	log.Infof("Attempting to connect to the database.")
	err := backoff.Retry(dbConnect, policy)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	app.DB = db
	log.Info("Connected to the database successfully.")

	// Initialize the HTTP client
	app.Client = &http.Client{}

	return nil
}

// Close the database connection
func (app *App) Shutdown() {
	err := app.DB.Close()
	if err != nil {
		log.Errorf("Failed to close the database: ", err)
	} else {
		log.Info("Closed the database successfully.")
	}
	log.Info("Shutdown the API crawler.")
}
