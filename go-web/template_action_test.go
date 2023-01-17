package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Personal struct {
	Title string
	Name  string
}

type Comparator struct {
	Title      string
	ValueOne   int
	ValueTwo   int
	ValueThree int
}

func TemplateWithAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/conditiontemplate.gohtml"))
	t.ExecuteTemplate(w, "conditiontemplate.gohtml", Personal{
		Title: "HI",
	})
}

func TemplateWithActionComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", Comparator{
		Title:      "Comparator",
		ValueOne:   100,
		ValueTwo:   40,
		ValueThree: 50,
	})
}

func TemplateWithActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/templatewithrange.gohtml"))
	t.ExecuteTemplate(w, "templatewithrange.gohtml", map[string]interface{}{
		"Title": "Range",
		"Skills": []string{
			".NET", "Go", "PHP", "SQL",
		},
	})
}

func TestTemplateWithAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWithAction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateWithActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWithActionComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateWithActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWithActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
