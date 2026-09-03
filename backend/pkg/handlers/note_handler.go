package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/noteapp/backend/pkg/patterns"
	"github.com/noteapp/backend/pkg/storage"
)

type NoteHandler struct {
	store   storage.NoteStorage
	factory patterns.NoteFactory
}

func NewNoteHandler(store storage.NoteStorage) *NoteHandler {
	return &NoteHandler{
		store:   store,
		factory: patterns.NewNoteFactory(),
	}
}

func (h *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Simple router logic
	if r.URL.Path == "/api/notes" {
		switch r.Method {
		case http.MethodGet:
			h.handleGetNotes(w, r)
		case http.MethodPost:
			h.handleCreateNote(w, r)
		default:
			patterns.NewResponseBuilder(w).Status(http.StatusMethodNotAllowed).Error("Method not allowed").Send()
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api/notes/") {
		id := strings.TrimPrefix(r.URL.Path, "/api/notes/")
		if id != "" {
			switch r.Method {
			case http.MethodDelete:
				h.handleDeleteNote(w, r, id)
			default:
				patterns.NewResponseBuilder(w).Status(http.StatusMethodNotAllowed).Error("Method not allowed").Send()
			}
			return
		}
	}

	patterns.NewResponseBuilder(w).Status(http.StatusNotFound).Error("Not found").Send()
}

func (h *NoteHandler) handleGetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.store.GetAll()
	if err != nil {
		patterns.NewResponseBuilder(w).Status(http.StatusInternalServerError).Error(err.Error()).Send()
		return
	}

	patterns.NewResponseBuilder(w).Status(http.StatusOK).Data(notes).Send()
}

func (h *NoteHandler) handleCreateNote(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		patterns.NewResponseBuilder(w).Status(http.StatusBadRequest).Error("Invalid request payload").Send()
		return
	}

	if req.Title == "" || req.Content == "" {
		patterns.NewResponseBuilder(w).Status(http.StatusBadRequest).Error("Title and Content are required").Send()
		return
	}

	// Use Factory Pattern to create the note
	note := h.factory.CreateNote(req.Title, req.Content)

	if err := h.store.Save(note); err != nil {
		patterns.NewResponseBuilder(w).Status(http.StatusInternalServerError).Error(err.Error()).Send()
		return
	}

	// Use Builder Pattern to send response
	patterns.NewResponseBuilder(w).Status(http.StatusCreated).Message("Note created successfully").Data(note).Send()
}

func (h *NoteHandler) handleDeleteNote(w http.ResponseWriter, r *http.Request, id string) {
	err := h.store.Delete(id)
	if err != nil {
		patterns.NewResponseBuilder(w).Status(http.StatusNotFound).Error("Note not found").Send()
		return
	}

	patterns.NewResponseBuilder(w).Status(http.StatusOK).Message("Note deleted successfully").Send()
}
