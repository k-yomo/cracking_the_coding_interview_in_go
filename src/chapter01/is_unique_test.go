package main

import (
	"reflect"
	"testing"
)

func TestIsUnique(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"orange", true},
		{"apple", false},
	}
	for _, tc := range testCases {
		result := isUnique(tc.str)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %s", tc.expected, result, tc.str)
		}
	}
}
