package render

import (
	"encoding/json"
	"net/http"
)

type RespondData struct {
	Message string `json:"message"`
}

// Respond renders a json response
func Respond(w http.ResponseWriter, v interface{}) {
	respond(w, v, http.StatusOK)
}

// Error renders an error as json
func Error(w http.ResponseWriter, err error, status int) {
	respond(w, &RespondData{err.Error()}, status)
}

// responds does the json response heavy lifting
func respond(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}
