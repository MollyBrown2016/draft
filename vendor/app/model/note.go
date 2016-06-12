package model

import (
	"fmt"
	"time"

	"app/shared/database"
)

// Note table contains the information for each note
type Note struct {
	ID        uint32    `db:"id"`
	Content   string    `db:"content"`
	UID       uint32    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   uint8     `db:"deleted"`
}

// NoteID returns the note id
func (u *Note) NoteID() string {
	r := fmt.Sprintf("%v", u.ID)
	return r
}

// NoteByID gets note by ID
func NoteByID(userID string, noteID string) (Note, error) {
	var err error

	result := Note{}
	err = database.SQL.Get(&result, "SELECT id, content, user_id, created_at, updated_at, deleted FROM note WHERE id = ? AND user_id = ? LIMIT 1", noteID, userID)

	return result, standardizeError(err)
}

// NotesByUserID gets all notes for a user
func NotesByUserID(userID string) ([]Note, error) {
	var err error

	var result []Note
	err = database.SQL.Select(&result, "SELECT id, content, user_id, created_at, updated_at, deleted FROM note WHERE user_id = ?", userID)

	return result, standardizeError(err)
}

// NoteCreate creates a note
func NoteCreate(content string, userID string) error {
	var err error
	_, err = database.SQL.Exec("INSERT INTO note (content, user_id) VALUES (?,?)", content, userID)

	return standardizeError(err)
}

// NoteUpdate updates a note
func NoteUpdate(content string, userID string, noteID string) error {
	_, err := database.SQL.Exec("UPDATE note SET content=? WHERE id = ? AND user_id = ? LIMIT 1", content, noteID, userID)

	return standardizeError(err)
}

// NoteDelete deletes a note
func NoteDelete(userID string, noteID string) error {
	var err error
	_, err = database.SQL.Exec("DELETE FROM note WHERE id = ? AND user_id = ?", noteID, userID)

	return standardizeError(err)
}
