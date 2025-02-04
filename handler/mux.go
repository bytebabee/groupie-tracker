package handler

import (
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"strings"
)

func TempSelector(artists []models.Artists) *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	if _, err := http.Dir("static").Open("."); err != nil {
		fmt.Printf("Directory style may not be accessible: %v", err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/home":
			HomeHandler(w, r, artists)
		case "/discover":
			DiscoverHandler(w, r, artists) // Pass artists to the DiscoverHandler
		case "/about":
			AboutHandler(w, r)
		default:
			if strings.HasPrefix(r.URL.Path, "/artist/") {
				ArtistHandler(w, r, artists)
				return
			} else {
				ErrorHandler(w, r)
			}
		}
	})
	return mux
}
