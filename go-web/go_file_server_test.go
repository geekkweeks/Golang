package go_web

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	// directory files
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// access the file location by url
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// create server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
