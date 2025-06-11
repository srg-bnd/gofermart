package response

import (
	"encoding/json"
	"net/http"
)

type StatusCode = int

func JSON(w http.ResponseWriter, status StatusCode, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, status StatusCode, message string) {
	JSON(w, status, map[string]string{"error": message})
}
