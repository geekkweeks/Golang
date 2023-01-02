package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMultipleEndpoint(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "My Home")
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "About my self")
	})

	mux.HandleFunc("/contact", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Contact me")
	})

	mux.HandleFunc("/admin/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Admin page")
	})

	mux.HandleFunc("/admin/setting", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Admin setting")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
