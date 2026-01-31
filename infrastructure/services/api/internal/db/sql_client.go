package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
	"time"
)

// NewSQLClient initializes a new MySQL database connection using the provided environment variable getter.
//
// The function prioritizes a MYSQL_URL environment variable (e.g. "mysql://user:pass@host:port/db").
// If MYSQL_URL is unset it falls back to reading the following variables:
//   - MYSQL_HOST
//   - MYSQL_PORT
//   - MYSQL_USER
//   - MYSQL_PASSWORD
//   - MYSQL_DATABASE
func NewSQLClient(getenv func(string) string) (*sql.DB, error) {
	var dsn string

	if connStr := getenv("MYSQL_URL"); connStr != "" {
		u, err := url.Parse(connStr)
		if err != nil {
			return nil, fmt.Errorf("invalid MYSQL_URL: %w", err)
		}

		user := u.User.Username()
		password, _ := u.User.Password()
		hostAddr := u.Host
		dbname := strings.TrimPrefix(u.Path, "/")

		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			user, password, hostAddr, dbname,
		)

	} else {

		user := getenv("MYSQL_USER")
		password := getenv("MYSQL_PASSWORD")
		host := getenv("MYSQL_HOST")
		port := getenv("MYSQL_PORT")
		dbname := getenv("MYSQL_DATABASE")

		if user == "" || password == "" || host == "" || port == "" || dbname == "" {
			return nil, fmt.Errorf("missing required database variables")
		}

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			user, password, host, port, dbname,
		)

	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
