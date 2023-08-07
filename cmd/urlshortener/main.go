package main

import (
	"log"
	"net/http"

	"github.com/scetle/urlshortener/internal/database"
	"github.com/scetle/urlshortener/internal/handlers"
)

func main() {
  db, err := database.NewDB()
  if err != nil {
    log.Fatal(err)
  }

  err = database.Migrate(db.DB)
  if err != nil {
    log.Fatal(err)
  }

  mux := http.NewServeMux()
  mux.HandleFunc("/", handlers.IndexHandler)
  mux.HandleFunc("/shorten", handlers.ShortenHandler)

  if err = http.ListenAndServe(":8080", mux); err != nil {
    log.Fatal(err) 
  }
}
