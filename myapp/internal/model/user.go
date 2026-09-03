package model

// User is the core domain entity.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserRequest is the DTO the API accepts on creation.
// Kept separate from User so the wire format can evolve independently
// of the domain model (e.g. no client-supplied ID).
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
