package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// middleware
type LogMiddleware struct {
	Handler http.Handler
}

// Error handler by middleware
type ErrorHandler struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Logic before exute handler
	fmt.Println("Before execute Handler")

	middleware.Handler.ServeHTTP(w, r)

	// Logic after exute handler
	fmt.Println("After execute Handler")
}

func (handler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
	}()
	handler.Handler.ServeHTTP(w, r)

}

func TestMidddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(w, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("foo Executed")
		fmt.Fprint(w, "Hello Foo")
	})
	// for testing error purpose
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic Executed")
		panic("Something went wrong")

	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
