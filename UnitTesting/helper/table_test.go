package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "TestOne",
			request:  "Alfan",
			expected: "Hello Alfan",
		},
		{
			name:     "TestTwo",
			request:  "AlfanZah",
			expected: "Hello AlfanZah",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}
