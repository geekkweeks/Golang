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

func MultipleParamWebHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	status := r.URL.Query().Get("status")

	fmt.Fprintf(w, "Hello %s %s %s", firstName, lastName, status)
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

func TestMultipleQueryParamHttpTest(t *testing.T) {
	paramQueryFirstName := "first_name=Alfan"
	paramQueryLastName := "last_name=Zahriyono"
	paramQueryStatus := "status=Active"
	request := httptest.NewRequest(http.MethodGet, baseUrl+"/home?"+paramQueryFirstName+"&"+paramQueryLastName+"&"+paramQueryStatus, nil)
	recorder := httptest.NewRecorder()

	MultipleParamWebHandler(recorder, request)

	response := recorder.Result()
	// Read from body response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)
}
