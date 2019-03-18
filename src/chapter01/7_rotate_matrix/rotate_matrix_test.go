package __rotate_matrix

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
	}
	for _, tc := range testCases {
		result := rotateMatrix(tc.matrix)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %v", tc.expected, result, tc.matrix)
		}
	}
}

func TestBetterRotateMatrix(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			[][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
		},
	}
	for _, tc := range testCases {
		result := betterRotateMatrix(tc.matrix)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("failed test expected: %v, got: %v, params: %v", tc.expected, result, tc.matrix)
		}
	}
}
