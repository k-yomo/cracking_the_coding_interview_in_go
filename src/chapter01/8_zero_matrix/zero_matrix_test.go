package __zero_matrix

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{0, 0, 0}, {4, 5, 0}, {7, 8, 0}},
		},
		{
			[][]int{{1, 2, 0, 4}, {5, 6, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{0, 0, 0, 0}, {0, 6, 0, 8}, {0, 0, 0, 0}, {0, 14, 0, 16}},
		},
	}
	for _, tc := range testCases {
		result := zeroMatrix(tc.matrix)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %v", tc.expected, result, tc.matrix)
		}
	}
}

func TestAnotherZeroMatrix(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{0, 0, 0}, {4, 5, 0}, {7, 8, 0}},
		},
		{
			[][]int{{1, 2, 0, 4}, {5, 6, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{0, 0, 0, 0}, {0, 6, 0, 8}, {0, 0, 0, 0}, {0, 14, 0, 16}},
		},
	}
	for _, tc := range testCases {
		result := anotherZeroMatrix(tc.matrix)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %v", tc.expected, result, tc.matrix)
		}
	}
}
