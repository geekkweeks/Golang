package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Address struct {
	Street string
}

type Page struct {
	Title       string
	Name        string
	Description string
	Address     Address
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title":       "Mahir Go",
		"Name":        "Alfan",
		"Description": "Ok nih",
		"Address": map[string]interface{}{
			"Street": "Jalan Kemayoran",
		},
	})
}

func TemplateDataMapWithStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "HTML WITH STRUCT",
		Name:  "Developer Tampan",
		Address: Address{
			Street: "Jalan Kebagusan",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Print(string(body))
}

func TestTemplateDataMapWithStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMapWithStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Print(string(body))
}
