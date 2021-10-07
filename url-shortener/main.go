package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "localhost:8080"
	if str := os.Getenv("HTTP_ADDR"); str != "" {
		addr = str
	}
	srv := http.Server{
		Addr:         addr,
		Handler:      NewHandler(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Listening on %s...", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
