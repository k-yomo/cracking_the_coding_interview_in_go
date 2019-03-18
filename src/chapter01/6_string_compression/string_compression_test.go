package __string_compression

import (
	"reflect"
	"testing"
)

func TestCompressString(t *testing.T) {
	testCases := []struct {
		str        string
		expected   string
	}{
		{"aaaabbcddda", "a4b2c1d3a1"},
		{"abbbbccc", "a1b4c3"},
		{"abbccc", "abbccc"},
	}
	for _, tc := range testCases {
		result := compressString(tc.str)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %s", tc.expected, result, tc.str)
		}
	}
}
