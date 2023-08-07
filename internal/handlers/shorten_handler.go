package handlers

import (
  "github.com/scetle/urlshortener/internal/service"
  "github.com/scetle/urlshortener/internal/database"
  "github.com/scetle/urlshortener/internal/models"
  "net/http"
  "fmt"
  "html/template"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.NotFound(w, r)
    return
  }
  originalURL := r.FormValue("url")
  tmpl, err := template.ParseFiles("web/templates/index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  shortURL := fmt.Sprintf("localhost:8080/%s", service.ShortenURL(originalURL))

  db, err := database.NewDB()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  url := models.URL{
    OriginalURL: originalURL,
    ShortURL:    shortURL,
  }

  existingURL, err := database.CheckIfExists(db.DB, url.OriginalURL)
  if err == nil {
    http.Error(w, err.Error(), http.StatusUnprocessableEntity)
    fmt.Fprintf(w, "This URL has already been shortened: %s", existingURL)
  } else {
    err = database.AddURL(db.DB, url)
    if err != nil {
      fmt.Println("no")
    } else {
      tmpl.Execute(w, url)
    }
  }
}
