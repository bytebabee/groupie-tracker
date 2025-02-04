package handler

import (
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, artists []models.Artists) {
	if artists == nil {
		BadFetchHandler(w, r)
		return
	}
	tempHome, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorFiveHandler(w, r, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	tempHome.Execute(w, nil)

}
