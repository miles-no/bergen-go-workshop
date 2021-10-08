package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/miles-no/bergen-go-workshop/url-shortener/shortener"
)

func main() {
	addr := "localhost:8080"
	if str := os.Getenv("HTTP_ADDR"); str != "" {
		addr = str
	}

	// shortener := shortener.NewInMemory()
	// shortener, err := shortener.NewOnDisk("./data")
	shortener, err := shortener.NewRedis("redis:6379")
	if err != nil {
		log.Fatalf("Failed to init shortener: %s", err)
	}
	srv := http.Server{
		Addr:         addr,
		Handler:      NewHandler(shortener),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Listening on %s...", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
