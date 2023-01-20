package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterWithParams(t *testing.T) {

	employeeText := "employeeId"
	router := httprouter.New()
	router.GET("/employee/:employeeId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Get employeeid value from url
		employeeId := params.ByName("employeeId")
		text := employeeText + ":" + employeeId
		fmt.Fprint(writer, text)
	})

	employeeId := "15"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/employee/"+employeeId, nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "employeeId:15", string(body))

}

func TestRouterWithNamePrameter(t *testing.T) {

	router := httprouter.New()
	router.GET("/employees/:id/department/:departmentId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		departmentId := params.ByName("departmentId")
		text := "EmployeeID: " + id + " DepartmentId:" + departmentId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/employees/20/department/12", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "EmployeeID: 20 DepartmentId:12", string(body))
}

func TestRouterWithCatchAllPrameter(t *testing.T) {

	router := httprouter.New()
	router.GET("/employees/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/employees/profile/picture", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /profile/picture", string(body))
}
