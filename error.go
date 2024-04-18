package main

import "net/http"

// handler for "GET /vq/err"
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
