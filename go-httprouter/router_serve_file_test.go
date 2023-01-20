package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {

	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	// url must with *filepath  hardcoded by ServeFiles function
	router.ServeFiles("/files/*filepath", http.FS(directory))

	targetURL := "http://localhost:3000/files/testfile.txt"
	request := httptest.NewRequest(http.MethodGet, targetURL, nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hi Golang developer", string(body))

}
