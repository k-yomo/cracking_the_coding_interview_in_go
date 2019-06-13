package __rotate_matrix

// 1.7 Rotate Matrix
// Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
// write a function to rotate the image by 90 degrees. Can you do this in place?
func rotateMatrix(matrix [][]int) [][]int {
	matrixLength := len(matrix)
	rotatedMatrix := make([][]int, matrixLength)
	for row := 0; row < matrixLength; row++ {
		rotatedMatrix[row] = make([]int, matrixLength)
		for i := 0; i < matrixLength; i++ {
			rotatedMatrix[row][i] = matrix[matrixLength-i-1][row]
		}
	}
	return rotatedMatrix
}

// this is more efficient than above by not allocating memory for copied array
func betterRotateMatrix(matrix [][]int) [][]int {
	matrixLength := len(matrix)
	// start from outermost layer and goes to inside
	for layer := 0; layer < matrixLength/2; layer++ {
		first := layer
		last := matrixLength - 1 - first
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]          // left -> top
			matrix[last-offset][first] = matrix[last][last-offset] // bottom -> left
			matrix[last][last-offset] = matrix[i][last]            // right -> bottom
			matrix[i][last] = top                                  // top -> right
		}
	}
	return matrix
}
