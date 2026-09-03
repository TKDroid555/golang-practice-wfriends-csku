package repository

import (
	"fmt"
	"sync"

	"myapp/internal/model"
)

// UserRepository defines how the service layer talks to storage.
// Because it's an interface, swapping the in-memory implementation
// below for a real Postgres/MySQL one later requires zero changes
// to the service or handler layers.
type UserRepository interface {
	Create(u model.User) (model.User, error)
	GetByID(id int) (model.User, error)
	List() ([]model.User, error)
}

// inMemoryUserRepository is a placeholder implementation so this
// boilerplate runs out of the box with no database required.
// Replace with a real implementation backed by database/sql, pgx, etc.
type inMemoryUserRepository struct {
	mu     sync.Mutex
	nextID int
	users  map[int]model.User
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		nextID: 1,
		users:  make(map[int]model.User),
	}
}

func (r *inMemoryUserRepository) Create(u model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u.ID = r.nextID
	r.nextID++
	r.users[u.ID] = u
	return u, nil
}

func (r *inMemoryUserRepository) GetByID(id int) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, ok := r.users[id]
	if !ok {
		return model.User{}, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (r *inMemoryUserRepository) List() ([]model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]model.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}
	return users, nil
}
