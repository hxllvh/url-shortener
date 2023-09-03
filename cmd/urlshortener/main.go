package main

import (
	"log"
	"net/http"

	"github.com/scetle/url-shortener/internal/database"
	"github.com/scetle/url-shortener/internal/handlers"
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
  mux.HandleFunc("/index", handlers.IndexHandler)
  mux.HandleFunc("/shorten", handlers.ShortenHandler)
  mux.HandleFunc("/", handlers.RedirectHandler)
  if err = http.ListenAndServe(":8080", mux); err != nil {
    log.Fatal(err) 
  }
}
