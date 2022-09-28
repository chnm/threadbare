package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"

	"github.com/hashicorp/go-retryablehttp"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/ratelimit"
)

type Config struct {
	dbstr    string
	loglevel string
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

	ll, ok := os.LookupEnv("THREADBARE_LOG_LEVEL")
	if !ok {
		return errors.New("THREADBARE_LOG_LEVEL environment variable not set")
	}
	app.Config.loglevel = ll

	// Set logging level
	switch app.Config.loglevel {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	}

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

	// Set up client for HTTP requests. Automatically retry.
	rc := retryablehttp.NewClient()
	rc.RetryWaitMin = 10 * time.Second
	rc.RetryWaitMax = 2 * time.Minute
	rc.RetryMax = 6
	rc.HTTPClient.Timeout = apiTimeout * time.Second
	rc.Logger = nil
	app.Client = rc.StandardClient()

	// Create limiters
	il := ratelimit.New(200-20, ratelimit.Per(60*time.Second)) // 200 requests/minute
	app.Limiters.Items = il

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
