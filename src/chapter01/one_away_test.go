package main
import (
"reflect"
"testing"
)

func TestIsOneAway(t *testing.T) {
	testCases := []struct {
		str1      string
		str2     string
		expected bool
	}{
		{"orange", "orang", true},
		{"orange", "orange1", true},
		{"orange", "oranga", true},
		{"orange", "oran", false},
		{"orange", "orangaa", false},
		{"orange", "oranbo", false},
	}
	for _, tc := range testCases {
		result := isOneAway(tc.str1, tc.str2)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %s, %s", tc.expected, result, tc.str1, tc.str2)
		}
	}
}

