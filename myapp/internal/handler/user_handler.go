package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"myapp/internal/model"
	"myapp/internal/service"
)

// UserHandler is the "Controller": it parses requests, delegates
// to the service layer, and writes JSON responses. There is no
// server-side "View" — the JSON response itself is what the
// frontend (React/Vue/mobile/etc.) consumes and renders.
type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Routes registers this handler's endpoints onto the given mux.
// Go 1.22+ supports method + path patterns natively, no router library needed.
func (h *UserHandler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", h.Create)
	mux.HandleFunc("GET /api/users", h.List)
	mux.HandleFunc("GET /api/users/{id}", h.Get)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, users)
}

// --- small JSON helpers, usually pulled into a shared internal/httpx package ---

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
