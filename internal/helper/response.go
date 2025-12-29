package helper

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

func RespondJSON(w http.ResponseWriter, status int, success bool, data interface{}, errors []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := APIResponse{
		Success: success,
		Data:    data,
		Errors:  errors,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
