package rest

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /li", createLink)

	// TODO: GET - REDIRECT
}
