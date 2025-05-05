package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"backend/config"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// Error Messages
	databaseOpenError string = "Could not open database connection: %s"
	databasePingError string = "Could not ping database: %s"

	// Database Default Configuration
	maxOpenConns = 20
	maxIdleConns = 10
	maxLifeTime  = 5 * time.Minute
	maxIdleTime  = 5 * time.Minute
)

// Slice of database Options
var options = []string{
	"charset=utf8",
	// "parseTime=true",
	"loc=Local",
	"multiStatements=true",
}

// NewMySQLDatabase open and returns a new MySQL connection
func NewMySQLDatabase(cfg config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		// user:password@protocol(host:port)/database?options
		"%s:%s@tcp(%s:%s)/%s?%s",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database,
		strings.Join(options, "&"),
	)

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf(databaseOpenError, err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf(databasePingError, err)
	}

	database.SetMaxIdleConns(maxIdleConns)
	database.SetMaxOpenConns(maxOpenConns)
	database.SetConnMaxLifetime(maxLifeTime)
	database.SetConnMaxIdleTime(maxIdleTime)

	return database
}
