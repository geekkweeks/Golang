package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Hello %s", name)
}

func TestQueryParamHttpTest(t *testing.T) {
	paramQuery := "?name=Alfan"
	request := httptest.NewRequest(http.MethodGet, baseUrl+"/home"+paramQuery, nil)
	recorder := httptest.NewRecorder()

	HelloWebHandler(recorder, request)

	response := recorder.Result()
	// Read from body response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)
}
