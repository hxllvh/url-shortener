package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" || r.Method != http.MethodGet {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles("web/templates/index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = tmpl.Execute(w, nil)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
