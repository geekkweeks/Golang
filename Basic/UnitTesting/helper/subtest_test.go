package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
	"unicode/utf8"
)

func TestSubTest(t *testing.T) {
	t.Run("SubTestOne", func(t *testing.T) {
		result := 40
		require.Equal(t, 30, result, "Result must be 30")
	})
	t.Run("SubTestTwo", func(t *testing.T) {
		result := "jakartabandung"
		require.Equal(t, 10, utf8.RuneCountInString(result), "Max chars 10")
	})
}
