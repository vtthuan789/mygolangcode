package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection poll
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// maxOpenDbConn us the maximum connections to the database open at any given time
const maxOpenDbConn = 10

// maxIdleDbConn is the maximum connections can remain in the pool but remain idle
const maxIdleDbConn = 5

// maxDbLifeTime is the maximum lifetime for a database connection
const maxDbLifeTime = 5 * time.Minute

// ConnectSQL creates database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB tries to ping the database
func testDB(d *sql.DB) error {
	return d.Ping()
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
