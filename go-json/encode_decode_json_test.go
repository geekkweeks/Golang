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
	Skills      []string
	Benefits    []EmployeeBenefit
}

type EmployeeBenefit struct {
	BenefitName string
	IsExpired   bool
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

func TestDecodeJson(t *testing.T) {
	jsonString := `{"Id":35,"Name":"Alfan Zah","Department":"IT","Salary":23000,"IsPermanent":true}`
	jsonBytes := []byte(jsonString)

	employee := &Employee{}

	err := json.Unmarshal(jsonBytes, employee)
	if err != nil {
		panic(err)
	}

	fmt.Println(employee)
	fmt.Println(employee.Id)
	fmt.Println(employee.Name)

}

func TestJSONArray(t *testing.T) {
	employeeObj := Employee{
		Id:          23,
		Name:        "Alfan Zahriyono",
		Department:  "IT",
		Salary:      23000,
		IsPermanent: true,
		Skills:      []string{"C#", "ReactJS", "JS"},
	}
	logJson(employeeObj)
}

func TestDecodeJSONArray(t *testing.T) {
	jsonString := `{"Id":23,"Name":"Alfan Zahriyono","Department":"IT","Salary":23000,"IsPermanent":true,"Skills":["C#","ReactJS","JS"]}`
	jsonBytes := []byte(jsonString)

	employee := &Employee{}

	err := json.Unmarshal(jsonBytes, employee)
	if err != nil {
		panic(err)
	}

	fmt.Println(employee)
	fmt.Println(employee.Skills[0])

}

func TestEncodeAdvanceJson(t *testing.T) {
	employeeObj := Employee{
		Id:          23,
		Name:        "Alfan Zahriyono",
		Department:  "IT",
		Salary:      23000,
		IsPermanent: true,
		Skills:      []string{"C#", "SQL", ".NET"},
		Benefits: []EmployeeBenefit{
			{
				BenefitName: "Private Insurance",
				IsExpired:   true,
			},
			{
				BenefitName: "Bonuses",
				IsExpired:   false,
			},
		},
	}
	logJson(employeeObj)
}

func TestDecodeAdvanceJson(t *testing.T) {
	jsonString := `{"Id":23,"Name":"Alfan Zahriyono","Department":"IT","Salary":23000,"IsPermanent":true,"Skills":["C#","SQL",".NET"],"Benefits":[{"BenefitName":"Private Insurance","IsExpired":true},{"BenefitName":"Bonuses","IsExpired":false}]}`
	jsonByte := []byte(jsonString)

	employee := &Employee{}
	err := json.Unmarshal(jsonByte, employee)
	if err != nil {
		panic(err)
	}
	fmt.Println(employee)
	fmt.Println(employee.Benefits)
}
