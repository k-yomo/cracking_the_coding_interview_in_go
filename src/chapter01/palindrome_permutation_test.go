package main

import (
	"reflect"
	"testing"
)

func TestcheckPalindromePermutation(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"taco cat", true},
		{"taco cats", false},
	}
	for _, tc := range testCases {
		result := checkPalindromePermutation(tc.str)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %s", tc.expected, result, tc.str)
		}
	}
}
