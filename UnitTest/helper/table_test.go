package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string //sub test name
		request string
		expect string
	}{
		{
			name: "TestingAlfan",
			request: "alfan",
			expect: "hello alfan",
		},
		{
			name: "TestingZahriyono",
			request: "zahriyono",
			expect: "hello zahriyono",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expect, result)
		})

	}


}
