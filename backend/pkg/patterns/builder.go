package patterns

import (
	"encoding/json"
	"net/http"
)

// Builder Pattern: Used to build complex HTTP JSON responses step by step.

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ResponseBuilder struct {
	w       http.ResponseWriter
	status  int
	payload APIResponse
}

func NewResponseBuilder(w http.ResponseWriter) *ResponseBuilder {
	return &ResponseBuilder{
		w:      w,
		status: http.StatusOK,
		payload: APIResponse{
			Success: true,
		},
	}
}

func (b *ResponseBuilder) Status(status int) *ResponseBuilder {
	b.status = status
	if status >= 400 {
		b.payload.Success = false
	}
	return b
}

func (b *ResponseBuilder) Message(message string) *ResponseBuilder {
	b.payload.Message = message
	return b
}

func (b *ResponseBuilder) Data(data interface{}) *ResponseBuilder {
	b.payload.Data = data
	return b
}

func (b *ResponseBuilder) Error(errStr string) *ResponseBuilder {
	b.payload.Error = errStr
	b.payload.Success = false
	return b
}

func (b *ResponseBuilder) Send() {
	b.w.Header().Set("Content-Type", "application/json")
	b.w.WriteHeader(b.status)
	json.NewEncoder(b.w).Encode(b.payload)
}
