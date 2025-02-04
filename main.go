package main

import (
	"fmt"
	"groupie-tracker/handler"
	"groupie-tracker/utilities"
)

func main() {
	// fetch data once before use them
	artists, _ := utilities.FetchArtistsData()

	// listen a random port
	listener, port, _ := utilities.ListenPort()
	defer listener.Close()

	fmt.Printf("Server is listening on http://localhost:%v\n", port)

	// serve the random port
	err := utilities.Serve(listener, handler.TempSelector(artists))
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
