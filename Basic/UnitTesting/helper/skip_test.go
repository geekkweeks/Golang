package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// Skip function will cancel next process/code
func TestByUsingSkip(t *testing.T) {
	// operting sytem
	if runtime.GOOS == "windows" {
		t.Skip("Can not run by Operating Sytem Windows")
	}

	result := 30
	require.Equal(t, 10, result, "Result mut be 10")
	fmt.Println("Tehst by using assertion is done")
}
