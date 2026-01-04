package http

import (
	"encoding/json"
	"net/http"
)

// HealthHandler returns JSON { "status": "OK" }
// Why JSON? Structured responses are standard in microservices.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "Working fine in Notification"})
}

// AuthorizeHandler: keep it stubbed for now.
// We'll teach JSON parsing and 400 vs 501 soon.
func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
