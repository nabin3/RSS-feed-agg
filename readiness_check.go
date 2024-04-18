package main

import "net/http"

// Defining handler for "GET /v1/readiness"
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type respData struct {
		Status string `json:"status"`
	}

	respondWithJSON(w, 200, respData{Status: "ok"})
}
