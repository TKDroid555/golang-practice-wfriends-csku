package patterns

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/noteapp/backend/pkg/models"
)

// Factory Method Pattern: Encapsulates the logic of creating a new Note.
// This ensures that every note created follows the same initialization rules
// (e.g., generating a unique ID and setting the created_at timestamp).

type NoteFactory interface {
	CreateNote(title, content string) models.Note
}

type defaultNoteFactory struct{}

func NewNoteFactory() NoteFactory {
	return &defaultNoteFactory{}
}

func (f *defaultNoteFactory) CreateNote(title, content string) models.Note {
	return models.Note{
		ID:        generateID(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
