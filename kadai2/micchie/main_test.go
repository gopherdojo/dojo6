package main

import (
	"testing"
)

func TestValidationFormat(t *testing.T) {
	tests := []struct {
		format   string
		lower    string
		hasError bool
	}{
		{format: "JPG", lower: "jpg", hasError: false},
		{format: "jpg", lower: "jpg", hasError: false},
		{format: "PNG", lower: "png", hasError: false},
		{format: "png", lower: "png", hasError: false},
		{format: "GIF", lower: "", hasError: true},
		{format: "", lower: "", hasError: true},
	}

	for _, test := range tests {
		var lower string
		var hasError bool
		if l, err := ValidationFormat(test.format); err != nil {
			lower = l
			hasError = true
		}

		if hasError != test.hasError {
			t.Errorf("format: %#v (want: %v, %v, god: %v, %v)", test.format, test.lower, test.hasError, lower, hasError)
			continue
		}
	}
}
