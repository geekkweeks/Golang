package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string  `json:"product_id""`
	Name     string  `json:"product_name""`
	Price    float32 `json:"product_price""`
	Quantity int     `json:"product_quantity""`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "QW123",
		Name:     "VAPE",
		Price:    234.23,
		Quantity: 12,
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"product_id":"QW123","product_name":"VAPE","product_price":234.23,"product_quantity":12}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
