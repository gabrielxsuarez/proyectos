package api

import (
	"net/http"

	json "github.com/json-iterator/go"
)

/* ========== RESPONSE ========== */
type ApiHealthResponse struct {
	Status string `json:"status"`
}

/* ========== ENDPOINT ========== */
func ApiHealthEndpoint(w http.ResponseWriter, r *http.Request) {
	response := ApiHealthResponse{
		Status: "UP",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
