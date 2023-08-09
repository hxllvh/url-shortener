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

  dataURL := models.DataURL{
    OriginalURL: r.FormValue("url"),
    ShortURL: fmt.Sprintf("localhost:8080/%s", service.ShortenURL(r.FormValue("url"))),
  }

  existingURL := database.CheckIfExists(db.DB, dataURL.OriginalURL)
  if existingURL != "" {
    dataURL.ExistingURL = existingURL
    tmpl.Execute(w, dataURL)
    return
  }
  if err := database.AddURL(db.DB, models.URL{OriginalURL: dataURL.OriginalURL, ShortURL: dataURL.ShortURL}); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  tmpl.Execute(w, dataURL)
}
