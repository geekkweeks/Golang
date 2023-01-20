package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
}
