package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Employee struct {
	Id          int
	Name        string
	Department  string
	Salary      int
	IsPermanent bool
}

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestLogJson(t *testing.T) {
	logJson("Alfan")
	logJson(23)
	logJson(true)
	logJson([]string{"MTK", "BAHASA", "IPS"})

	//encode object json
	employeeObj := Employee{
		Id:          23,
		Name:        "Alfan Zahriyono",
		Department:  "IT",
		Salary:      23000,
		IsPermanent: true,
	}
	logJson(employeeObj)
}
