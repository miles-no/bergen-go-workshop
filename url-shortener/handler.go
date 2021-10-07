package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is a URL shortener which implements the http.Handler interface.
type Handler struct {
	router *mux.Router
}

// NewHandler returns a new Handler which implements the URL shortener's REST API.
func NewHandler() *Handler {
	h := &Handler{router: mux.NewRouter()}
	h.router.HandleFunc("/urls", h.HandleCreate).Methods("POST")
	h.router.HandleFunc("/urls/{id}", h.HandleRedirect).Methods("GET")
	return h
}

// ServeHTTP implements the http.Handler interface required by the http.Server.
func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Pass the request to the router.
	h.router.ServeHTTP(w, req)
}

// URLResource defines the JSON resource of the 'POST /urls'-endpoint.
//
// The same resource type is used for both the request body and the response.
type URLResource struct {
	// URL is the fully-qualified URL to be shortened.
	// It should be included in both the request body and the response.
	URL string `json:"url"`

	// ShortURL is the shortened URL created by this service.
	// If set by a requesting client, the service should respond with
	// status 401 (Bad request).
	ShortURL string `json:"short_url"`
}

func (h *Handler) HandleCreate(w http.ResponseWriter, req *http.Request) {
	log.Print("Got create request")
	http.Error(w, "TODO", http.StatusNotImplemented)
}

func (h *Handler) HandleRedirect(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Got request for id %s", id)
	http.Error(w, "TODO", http.StatusNotImplemented)
}
