package cart

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Status      int    `json:"status"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

// ErrorNotFound Returns a Not found error
func ErrorNotFound(w http.ResponseWriter, description string) {
	err := &errorResponse{
		Status:      404,
		Error:       "Not found",
		Description: "Not found",
	}
	if description != "" {
		err.Description = description
	}
	json.NewEncoder(w).Encode(err)
}

// ErrorUnauthorised Returns an Unauthorised error
func ErrorUnauthorised(w http.ResponseWriter, description string) {
	err := &errorResponse{
		Status:      401,
		Error:       "Unauthorised",
		Description: "Unauthorised",
	}

	if description != "" {
		err.Description = description
	}

	json.NewEncoder(w).Encode(err)
}

// ErrorUnprocessable Returns an Unprocessable entity error:
func ErrorUnprocessable(w http.ResponseWriter, description string) {
	err := &errorResponse{
		Status:      422,
		Error:       "Unprocessable entity",
		Description: "Unprocessable entity",
	}

	if description != "" {
		err.Description = description
	}

	json.NewEncoder(w).Encode(err)
}
