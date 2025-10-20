package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// MustOpen opens (and initializes) the database. Logs fatal on any error.
func MustOpen(path string) {
	if db != nil {
		return
	}

	d, err := sql.Open("sqlite3", path+"?_busy_timeout=5000&_foreign_keys=on")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	// Enable WAL for better concurrency
	if _, err := d.Exec(`PRAGMA journal_mode = WAL;`); err != nil {
		d.Close()
		log.Fatal("failed to enable WAL:", err)
	}

	schema := []string{
		`CREATE TABLE IF NOT EXISTS Session (
			id TEXT PRIMARY KEY,
			token_hash TEXT NOT NULL,
			createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
			updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS Note (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			note TEXT NOT NULL,
			createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
			updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS Link (
			id TEXT PRIMARY KEY,
			link TEXT NOT NULL,
			img_path TEXT,
			createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
			updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, stmt := range schema {
		if _, err := d.Exec(stmt); err != nil {
			d.Close()
			log.Fatal("failed to create tables:", err)
		}
	}

	db = d
	log.Println("database initialized successfully")
}

//
// ─── INSERT FUNCTIONS ─────────────────────────────────────────────────────────────
//

// AddToken inserts a token_hash into Session table. Returns the new record ID.
func AddToken(tokenHash string) (string, error) {
	if db == nil {
		return "", fmt.Errorf("database not initialized")
	}
	id := uuid.New().String()
	_, err := db.Exec(
		`INSERT INTO Session (id, token_hash, createdAt, updatedAt) VALUES (?, ?, ?, ?)`,
		id, tokenHash, time.Now(), time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("insert token: %w", err)
	}
	return id, nil
}

// AddNote inserts a sanitized note into Note table. Returns the new record ID.
func AddNote(title, note string) (string, error) {
	if db == nil {
		return "", fmt.Errorf("database not initialized")
	}

	id := uuid.New().String()
	_, err := db.Exec(
		`INSERT INTO Note (id, title, note, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)`,
		id, title, note, time.Now(), time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("insert note: %w", err)
	}
	return id, nil
}

// AddLink inserts a link and optional img_path into Link table. Returns the new record ID.
func AddLink(link, imgPath string) (string, error) {
	if db == nil {
		return "", fmt.Errorf("database not initialized")
	}
	id := uuid.New().String()
	_, err := db.Exec(
		`INSERT INTO Link (id, link, img_path, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)`,
		id, link, imgPath, time.Now(), time.Now(),
	)
	if err != nil {
		return "", fmt.Errorf("insert link: %w", err)
	}
	return id, nil
}

//
// ─── RETRIEVAL FUNCTIONS ─────────────────────────────────────────────────────────
//

// GetToken retrieves a full Token record by token_hash.
func GetToken(tokenHash string) (*Token, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var t Token
	err := db.QueryRow(`SELECT id, token_hash, createdAt, updatedAt FROM Session WHERE token_hash = ?`, tokenHash).
		Scan(&t.ID, &t.TokenHash, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("token not found")
		}
		return nil, fmt.Errorf("query token: %w", err)
	}
	return &t, nil
}

// GetNotes retrieves all notes from the Note table.
func GetNotes() ([]Note, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	rows, err := db.Query(`SELECT id, note, createdAt, updatedAt FROM Note ORDER BY createdAt DESC`)
	if err != nil {
		return nil, fmt.Errorf("query notes: %w", err)
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Note, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan note: %w", err)
		}
		notes = append(notes, n)
	}
	return notes, nil
}

// GetLinks retrieves all links from the Link table.
func GetLinks() ([]Link, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	rows, err := db.Query(`SELECT id, link, img_path, createdAt, updatedAt FROM Link ORDER BY createdAt DESC`)
	if err != nil {
		return nil, fmt.Errorf("query links: %w", err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var l Link
		if err := rows.Scan(&l.ID, &l.Link, &l.ImgPath, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan link: %w", err)
		}
		links = append(links, l)
	}
	return links, nil
}

//
// ─── UPDATE FUNCTION (NOTE ONLY) ─────────────────────────────────────────────────
//

// UpdateNote updates an existing note and refreshes updatedAt.
func UpdateNote(id, newNote string) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	_, err := db.Exec(
		`UPDATE Note SET note = ?, updatedAt = ? WHERE id = ?`,
		newNote, time.Now(), id,
	)
	if err != nil {
		return fmt.Errorf("update note: %w", err)
	}
	return nil
}

//
// ─── DELETE FUNCTIONS ─────────────────────────────────────────────────────────────
//

// DeleteToken removes a Session by ID.
func DeleteToken(id string) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	_, err := db.Exec(`DELETE FROM Session WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete token: %w", err)
	}
	return nil
}

// DeleteNote removes a Note by ID.
func DeleteNote(id string) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	_, err := db.Exec(`DELETE FROM Note WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete note: %w", err)
	}
	return nil
}

// DeleteLink removes a Link by ID.
func DeleteLink(id string) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	_, err := db.Exec(`DELETE FROM Link WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete link: %w", err)
	}
	return nil
}

//
// ─── DATA STRUCTURES ──────────────────────────────────────────────────────────────
//

// Note represents a single note record.
type Note struct {
	ID        string
	Title     string
	Note      string
	CreatedAt string
	UpdatedAt string
}

// Link represents a single link record.
type Link struct {
	ID        string
	Link      string
	ImgPath   string
	CreatedAt string
	UpdatedAt string
}

// Token represents a session token record.
type Token struct {
	ID        string
	TokenHash string
	CreatedAt string
	UpdatedAt string
}

//
// ─── CLOSE ────────────────────────────────────────────────────────────────────────
//

// Close safely closes the database connection.
func Close() error {
	if db == nil {
		return nil
	}
	err := db.Close()
	db = nil
	return err
}
