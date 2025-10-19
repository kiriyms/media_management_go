package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

// Open opens (or returns) the global DB connection. Call once at startup.
func Open(path string) error {
	if conn != nil {
		return nil
	}

	d, err := sql.Open("sqlite3", path+"?_busy_timeout=5000")
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	// enable WAL for better concurrent behavior
	if _, err := d.Exec("PRAGMA journal_mode = WAL;"); err != nil {
		d.Close()
		return fmt.Errorf("enable wal: %w", err)
	}

	// create table for stored tokens
	create := `CREATE TABLE IF NOT EXISTS tokens (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        token TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err := d.Exec(create); err != nil {
		d.Close()
		return fmt.Errorf("create tokens table: %w", err)
	}

	conn = d
	return nil
}

// InsertToken stores the JWT token string in the database.
func InsertToken(token string) error {
	if conn == nil {
		return fmt.Errorf("db not initialized")
	}

	_, err := conn.Exec("INSERT INTO tokens (token) VALUES (?)", token)
	if err != nil {
		return fmt.Errorf("insert token: %w", err)
	}
	return nil
}

// Close closes the global DB connection.
func Close() error {
	if conn == nil {
		return nil
	}
	err := conn.Close()
	conn = nil
	return err
}
