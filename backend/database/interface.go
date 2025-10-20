package database

// Database defines the interface for all database operations.
type Database interface {
	// Insert functions
	AddToken(tokenHash string) (string, error)
	AddNote(note string) (string, error)
	AddLink(link, imgPath string) (string, error)

	// Retrieval functions
	GetToken(tokenHash string) (*Token, error)
	GetNotes() ([]Note, error)
	GetLinks() ([]Link, error)

	// Update functions
	UpdateNote(id, newNote string) error

	// Delete functions
	DeleteToken(id string) error
	DeleteNote(id string) error
	DeleteLink(id string) error
}
