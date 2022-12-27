package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
// recommendeed using assertion to avoid if else
// the behaviour oth assert is equal to fail()
// if error found, still process the code below
*/
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Zahri")
	assert.Equal(t, "hello Alfan", result, "Result must be 'hello Alfan")
	fmt.Println("TestHelloWorldAssert is done")
}

/*
// recommendeed using assertion to avoid if else
// the behaviour oth assert is equal to failNow()
// if error found, the all process below will not be executed
*/
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Zahri")
	require.Equal(t, "hello Alfan", result, "Result must be 'hello Alfan")
	fmt.Println("TestHelloWorldAssert is done")
}