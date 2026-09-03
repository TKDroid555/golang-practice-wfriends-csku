package storage

import "github.com/noteapp/backend/pkg/models"

// Strategy Pattern: NoteStorage is an interface defining the operations
// for storing and retrieving notes. This allows us to switch between
// different storage strategies (e.g., Memory, PostgreSQL, Redis) without
// changing the business logic.
type NoteStorage interface {
	Save(note models.Note) error
	GetAll() ([]models.Note, error)
	GetByID(id string) (models.Note, error)
	Delete(id string) error
}
