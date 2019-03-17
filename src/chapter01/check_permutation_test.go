package main

import (
	"reflect"
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	testCases := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "tets", true},
		{"test", "ttes", true},
		{"test", "tesb", false},
		{"test", "testa", false},
		{"test", "", false},
	}
	for _, tc := range testCases {
		result := checkPermutation(tc.str1, tc.str2)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %s, %s", tc.expected, result, tc.str1, tc.str2)
		}
	}
}
