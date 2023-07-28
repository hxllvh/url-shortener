package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/scetle/urlshortener/internal/database"
	"github.com/scetle/urlshortener/internal/models"
	"github.com/scetle/urlshortener/internal/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  
  if r.Method == http.MethodGet {
    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  } else {
    http.NotFound(w, r)
  }
} 

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost {
    originalURL := r.FormValue("url")
    tmpl, err := template.ParseFiles("web/templates/index.html")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    
    shortURL := fmt.Sprintf("localhost:8080/%s", utils.ShortenURL(originalURL))

    dbConnection, err := database.NewDB()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
 
    url := models.URL{
      OriginalURL: originalURL,
      ShortURL:    shortURL,
    }

    existingURL, err := database.CheckIfExists(dbConnection.GetDB(), url.OriginalURL)
    if err == nil {
      http.Error(w, err.Error(), http.StatusUnprocessableEntity)
      fmt.Fprintf(w, "This URL has already been shortened: %s", existingURL)
    } else {
      err = database.AddURL(dbConnection.GetDB(), url)
      if err != nil {
        fmt.Println("no")
      } else {
        tmpl.Execute(w, url)
      }
    }

  } else {
    http.NotFound(w, r)
  }
}
