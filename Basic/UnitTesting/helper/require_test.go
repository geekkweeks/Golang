package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Require Vs Assertion
// Require will call function FailNow(). It means after the error occure, the testing will stop (next code will not be execute)
// Assertion will call function Fail(). It means after the error occure, the next code will be executed
func TestByUsingRequire(t *testing.T) {
	result := 30
	require.Equal(t, 10, result, "Result mut be 10")
	fmt.Println("Tehst by using assertion is done")
}
