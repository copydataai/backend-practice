package services

import (
	"encoding/json"
	"net/http"
)
var (
	ErrBadRequest   = ErrorResponse{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrBadFetch     = ErrorResponse{StatusCode: http.StatusInternalServerError, Type: "Error to fetch data", Message: "Error to send query to DB"}
	ErrNotFound             = ErrorResponse{StatusCode: http.StatusNotFound, Type: "not found in DB", Message: "This data not exists"}
	ErrInvalidJSON  = ErrorResponse{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	// decide between conflict and bad request
	ErrAlreadyExists = ErrorResponse{StatusCode: http.StatusConflict, Type: "duplicate_entry", Message: "Another entity has the same value as this field"}
)

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e ErrorResponse) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
