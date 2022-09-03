package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByUsingAssertion(t *testing.T) {
	result := 30
	assert.Equal(t, 10, result, "Result mut be 10")
	fmt.Println("Tehst by using assertion is done")
}
