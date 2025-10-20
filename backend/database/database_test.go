package database

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// setupTestDB creates a temporary SQLite database for testing.
func setupTestDB(t *testing.T) {
	t.Helper()

	// Use in-memory SQLite to avoid touching disk or prod data
	MustOpen(":memory:")

	// Clean up automatically after test finishes
	t.Cleanup(func() {
		if err := Close(); err != nil {
			t.Fatalf("failed to close test DB: %v", err)
		}
	})
}

// TestDatabaseSetup ensures that the schema initializes correctly.
func TestDatabaseSetup(t *testing.T) {
	setupTestDB(t)

	// Verify tables exist
	tables := []string{"Session", "Note", "Link"}
	for _, tbl := range tables {
		var name string
		err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tbl).Scan(&name)
		if err != nil {
			t.Fatalf("table %s not found: %v", tbl, err)
		}
	}
}

// TestAddAndGetToken tests insertion and retrieval of a token.
func TestAddAndGetToken(t *testing.T) {
	setupTestDB(t)

	id, err := AddToken("abc123hash")
	if err != nil {
		t.Fatalf("AddToken failed: %v", err)
	}

	token, err := GetToken("abc123hash")
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}

	if token.ID != id {
		t.Errorf("expected ID %s, got %s", id, token.ID)
	}
	if token.TokenHash != "abc123hash" {
		t.Errorf("expected TokenHash %s, got %s", "abc123hash", token.TokenHash)
	}
}

// TestAddGetUpdateDeleteNote tests full CRUD lifecycle for Note.
func TestAddGetUpdateDeleteNote(t *testing.T) {
	setupTestDB(t)

	// Create
	id, err := AddNote("Test note")
	if err != nil {
		t.Fatalf("AddNote failed: %v", err)
	}

	// Read
	notes, err := GetNotes()
	if err != nil {
		t.Fatalf("GetNotes failed: %v", err)
	}
	if len(notes) != 1 || notes[0].ID != id {
		t.Fatalf("expected one note with ID %s, got %+v", id, notes)
	}

	// Update
	err = UpdateNote(id, "Updated note text")
	if err != nil {
		t.Fatalf("UpdateNote failed: %v", err)
	}

	notes, _ = GetNotes()
	if notes[0].Note != "Updated note text" {
		t.Errorf("note not updated, got: %s", notes[0].Note)
	}

	// Delete
	if err := DeleteNote(id); err != nil {
		t.Fatalf("DeleteNote failed: %v", err)
	}

	notes, _ = GetNotes()
	if len(notes) != 0 {
		t.Fatalf("expected 0 notes after delete, got %d", len(notes))
	}
}

// TestAddGetDeleteLink tests CRUD for Link.
func TestAddGetDeleteLink(t *testing.T) {
	setupTestDB(t)

	id, err := AddLink("https://example.com", "/path/to/img.png")
	if err != nil {
		t.Fatalf("AddLink failed: %v", err)
	}

	links, err := GetLinks()
	if err != nil {
		t.Fatalf("GetLinks failed: %v", err)
	}

	if len(links) != 1 || links[0].ID != id {
		t.Fatalf("expected one link with ID %s, got %+v", id, links)
	}

	if err := DeleteLink(id); err != nil {
		t.Fatalf("DeleteLink failed: %v", err)
	}

	links, _ = GetLinks()
	if len(links) != 0 {
		t.Fatalf("expected 0 links after delete, got %d", len(links))
	}
}
