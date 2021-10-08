package main

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

// Handler is a URL shortener which implements the http.Handler interface.
type Handler struct {
	router *mux.Router

	s Shortener
}

type Shortener interface {
	Put(url string) (id string)
	Get(id string) (url string)
}

// NewHandler returns a new Handler which implements the URL shortener's REST API.
func NewHandler(s Shortener) *Handler {
	h := &Handler{
		router: mux.NewRouter(),
		// urls:   make(map[string]string),
		s: s,
	}
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
	var v URLResource
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if v.ShortURL != "" {
		http.Error(w, "you don't get to choose your short url", http.StatusBadRequest)
		return
	}
	id := h.s.Put(v.URL)
	v.ShortURL = path.Join(req.URL.Path, id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) HandleRedirect(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	url := h.s.Get(id)
	if url == "" {
		http.NotFound(w, req)
		return
	}
	http.Redirect(w, req, url, http.StatusFound)
}
