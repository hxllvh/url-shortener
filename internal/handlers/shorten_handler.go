package handlers

import (
  "github.com/scetle/url-shortener/internal/service"
  "github.com/scetle/url-shortener/internal/database"
  "github.com/scetle/url-shortener/internal/models"
  "net/http"
  "fmt"
  "html/template"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles("web/templates/index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  db, err := database.NewDB()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  originalURL := r.FormValue("url")
  shortURL := fmt.Sprintf("localhost:8080/%s", service.ShortenURL(originalURL))
  
  dataURL := models.DataURL{
    OriginalURL: originalURL,
    ShortURL: shortURL,
  }
  mainURL := models.URL{
    OriginalURL: originalURL,
    ShortURL: shortURL,
  }

  existingURL, err := db.AddURL(mainURL)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  } else if existingURL != "" {
    dataURL.ExistingURL = existingURL
  }

  tmpl.Execute(w, dataURL)
}
