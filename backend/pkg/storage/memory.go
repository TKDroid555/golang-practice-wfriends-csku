package storage

import (
	"errors"
	"sync"

	"github.com/noteapp/backend/pkg/models"
)

// Singleton Pattern: Ensure only one instance of the MemoryStorage exists.
var (
	instance *MemoryStorage
	once     sync.Once
)

// MemoryStorage is a concrete strategy of NoteStorage.
type MemoryStorage struct {
	mu    sync.RWMutex
	notes map[string]models.Note
}

// GetMemoryStorage returns the singleton instance of MemoryStorage.
func GetMemoryStorage() *MemoryStorage {
	once.Do(func() {
		instance = &MemoryStorage{
			notes: make(map[string]models.Note),
		}
	})
	return instance
}

// Save implements NoteStorage.Save
func (m *MemoryStorage) Save(note models.Note) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.notes[note.ID] = note
	return nil
}

// GetAll implements NoteStorage.GetAll
func (m *MemoryStorage) GetAll() ([]models.Note, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	result := make([]models.Note, 0, len(m.notes))
	for _, n := range m.notes {
		result = append(result, n)
	}
	return result, nil
}

// GetByID implements NoteStorage.GetByID
func (m *MemoryStorage) GetByID(id string) (models.Note, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	if note, ok := m.notes[id]; ok {
		return note, nil
	}
	return models.Note{}, errors.New("note not found")
}

// Delete implements NoteStorage.Delete
func (m *MemoryStorage) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if _, ok := m.notes[id]; !ok {
		return errors.New("note not found")
	}
	delete(m.notes, id)
	return nil
}
