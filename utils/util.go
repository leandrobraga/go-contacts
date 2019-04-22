package utils

import (
	"encoding/json"
	"net/http"
)

// Message build json message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond return json message
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(data)
}
