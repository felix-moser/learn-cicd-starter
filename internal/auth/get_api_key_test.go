package auth

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type test struct {
	input        http.Header
	outputString string
	outputErr    error
}

func TestGetAPIKey(t *testing.T) {
	tests := []test{
		// auth header api key == ""
		{
			input:        nil,
			outputString: "",
			outputErr:    ErrNoAuthHeaderIncluded,
		},
		// Malformed auth cause no Key
		{
			input: http.Header{
				"Authorization": []string{"Bearer"},
			},
			outputString: "",
			outputErr:    errors.New("malformed authorization header"),
		},
		// Best case
		{
			input: http.Header{
				"Authorization": []string{"ApiKey 123"},
			},
			outputString: "123",
			outputErr:    nil,
		},
	}

	for i, tc := range tests {
		fmt.Printf("Running test %v\n", i)
		got, gotErr := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.outputString, got) || !reflect.DeepEqual(tc.outputErr, gotErr) {
			t.Fatalf("expected output: %v got: %v; expected error: %v got %v", tc.outputString, got, tc.outputErr, gotErr)
		}
	}
}
