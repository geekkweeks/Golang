package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can run on windows")
	}

	result := HelloWorld("zahri")
	require.Equal(t, "hello Alfan", result, "Result must be 'hello Alfan")
	fmt.Println("TestHelloWorldAssert is done")
}