package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// adding payload
func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// json response
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// error response
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// func ParseJSON(r *http.Request, v any) error {
// 	if r.Body == nil {
// 		return fmt.Errorf("missing request body")
// 	}

// 	return json.NewDecoder(r.Body).Decode(v)
// }

// func GetTokenFromRequest(r *http.Request) string {
// 	tokenAuth := r.Header.Get("Authorization")
// 	tokenQuery := r.URL.Query().Get("token")

// 	if tokenAuth != "" {
// 		return tokenAuth
// 	}

// 	if tokenQuery != "" {
// 		return tokenQuery
// 	}

// 	return ""
// }
