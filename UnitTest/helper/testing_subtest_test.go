package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("zahriyonoTesting", func(t *testing.T) {
		result := HelloWorld("zahriyono")
		require.Equal(t, "hello Alfan", result, "Result must be 'hello Alfan")
	})

	t.Run("RonaldoTesting", func(t *testing.T) {
		result := HelloWorld("Ronaldo")
		require.Equal(t, "hello Alfan", result, "Result must be 'hello Alfan")
	})
}