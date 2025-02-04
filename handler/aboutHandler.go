package handler

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	tempAbout, err := template.ParseFiles("templates/about.html")
	if err != nil {
		ErrorFiveHandler(w, r, err)

		return
	}
	w.WriteHeader(http.StatusOK)
	tempAbout.Execute(w, nil)
}
