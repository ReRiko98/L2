package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\", false},
	}

	for _, tc := range testCases {
		result, err := UnpackString(tc.input)
		if (err != nil) != tc.err {
			t.Errorf("Expected error %v, but got: %v", tc.err, err)
		}

		if result != tc.expected {
			t.Errorf("Expected result %q, but got: %q", tc.expected, result)
		}
	}
}
