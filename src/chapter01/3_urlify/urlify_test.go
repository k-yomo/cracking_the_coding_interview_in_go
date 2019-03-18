package __urlify

import (
	"reflect"
	"testing"
)

func TestUrlify(t *testing.T) {
	testCases := []struct {
		str      string
		charLength int
		expected string
	}{
		{"Kanji Yomoda  ", 13, "Kanji%20Yomoda%20"},
		{"Kanji Yomoda  ", 5, "Kanji"},
	}
	for _, tc := range testCases {
		result := urlify(tc.str, tc.charLength)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %s", tc.expected, result, tc.str)
		}
	}
}
