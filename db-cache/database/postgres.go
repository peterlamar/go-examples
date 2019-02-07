package database

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	connectionTimeout = "5" // DB Connection Timeout in Sec
)

var (
	db *sqlx.DB
	// MaxConnections the max number of open database connections
	MaxConnections = 100
)

// ConnectDB loads up the postgres database
func ConnectDB() {
	log.Trace("Connecting to Postgresql.")
	var err error

	connStr := buildPGEnvString()

	db, err = sqlx.Connect("postgres", connStr) // Connect to database

	if err != nil {
		log.Fatal("Error validating Postgresql string: ", err)
	}
	db.SetMaxOpenConns(MaxConnections)
	log.Trace("Connected to Postgresql.")
}

// Build the environment string for the database connection
// Change to all optional with defaults connecting to container
func buildPGEnvString() (rtnString string) {
	if os.Getenv("POSTGRES_HOST") == "" {
		log.Fatal("Environment variable POSTGRES_HOST is not set")
	}
	if os.Getenv("POSTGRES_PASSWORD") == "" {
		log.Fatal("Environment variable POSTGRES_PASSWORD is not set")
	}
	if os.Getenv("POSTGRES_DB") == "" {
		log.Fatal("Environment variable POSTGRES_DB is not set")
	}
	if os.Getenv("POSTGRES_USER") == "" {
		log.Fatal("Environment variable POSTGRES_HOST is not set")
	}
	if os.Getenv("POSTGRES_PORT") == "" {
		os.Setenv("POSTGRES_PORT", "5432")
	}

	var conTime string
	if os.Getenv("CONNECTION_TIMEOUT") == "" {
		conTime = connectionTimeout
	} else {
		conTime = os.Getenv("CONNECTION_TIMEOUT")
	}

	rtnString += "host=" + os.Getenv("POSTGRES_HOST") + " "
	rtnString += "password=" + os.Getenv("POSTGRES_PASSWORD") + " "
	rtnString += "dbname=" + os.Getenv("POSTGRES_DB") + " "
	rtnString += "user=" + os.Getenv("POSTGRES_USER") + " "
	rtnString += "port=" + os.Getenv("POSTGRES_PORT") + " "
	rtnString += "sslmode=disable connect_timeout=" + conTime

	return
}

// GetDB is used to get the database connection
func GetDB() *sqlx.DB {
	return db
}

// SetDB is used to set the database connection
func SetDB(dbx *sqlx.DB) {
	db = dbx
}
