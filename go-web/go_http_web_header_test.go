package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"
)

func RequestHeaderHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprintf(w, contentType)
}

func ResponseHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Alfan zah")
	fmt.Fprintf(w, "Ok")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	fmt.Println(response.Header.Get("x-powered-by"))
}
