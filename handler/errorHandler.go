package handler

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	Error error
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("templates/404.html")
	if err != nil {
		ErrorFiveHandler(w, r, err)
		return
	}
	log.Print("Error: Page not found")
	w.WriteHeader(http.StatusNotFound)
	temp.Execute(w, nil)
}

func ErrorFiveHandler(w http.ResponseWriter, r *http.Request, errmsg error) {

	temp, err := template.ParseFiles("templates/500.html")
	if err != nil {
		log.Printf("Error : %v", err)
		http.Error(w, "Error while parsing the 500 template", http.StatusInternalServerError)
		return
	}
	log.Printf("Error : %v", errmsg)
	w.WriteHeader(http.StatusInternalServerError)
	temp.Execute(w, nil)
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/badrequest.html")
	if err != nil {
		ErrorFiveHandler(w, r, err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	temp.Execute(w, nil)
}

func BadFetchHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/fecthissue.html")
	if err != nil {
		ErrorFiveHandler(w, r, err)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	temp.Execute(w, nil)
}
