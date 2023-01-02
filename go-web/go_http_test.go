package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var baseUrl = "http://localhost:8080"

func GoWebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello All developers!")
}

func TestHttpTest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, baseUrl+"/home", nil)
	recorder := httptest.NewRecorder()

	GoWebHandler(recorder, request)

	response := recorder.Result()
	// Read from body response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)
}
