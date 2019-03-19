package __zero_matrix

// 1.8 Zero Matrix
// Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.
func zeroMatrix(matrix [][]int) [][]int {
	rowLength := len(matrix)
	colLength := len(matrix[0])
	var zeroPositions []map[string]int

	for x := 0; x < rowLength; x++ {
		for y := 0; y < colLength; y++ {
			if matrix[x][y] == 0 {
				zeroPositions = append(zeroPositions, map[string]int{"x": x, "y": y})
			}
		}
	}

	for i := 0; i < len(zeroPositions); i++ {
		x := zeroPositions[i]["x"]
		y := zeroPositions[i]["y"]
		for j := 0; j < colLength; j++ {
			matrix[x][j] = 0
		}
		for j := 0; j < rowLength; j++ {
			matrix[j][y] = 0
		}
	}
	return matrix
}

// This is better in terms of space complexity.
func anotherZeroMatrix(matrix [][]int) [][]int {
	rowLength := len(matrix[0])
	colLength := len(matrix)
	firstRowHasZero := false
	firstColHasZero := false

	// Check if first row has a zero
	for x := 0; x < colLength; x++ {
		if matrix[0][x] == 0 {
			firstRowHasZero = true
			break
		}
	}
	// Check if first column has a zero
	for y := 0; y < rowLength; y++ {
		if matrix[y][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// Check for zeros in the rest of the array
	for y := 1; y < rowLength; y++ {
		for x := 1; x < colLength; x++ {
			if matrix[y][x] == 0 {
				matrix[y][0] = 0
				matrix[0][x] = 0
			}
		}
	}

	// set zero based on values in first column
	for y := 0; y < rowLength; y++ {
		if matrix[y][0] == 0 {
			setRowZero(matrix, y)
		}
	}
	// set zero based on values in first row
	for x := 0; x < colLength; x++ {
		if matrix[0][x] == 0 {
			setColZero(matrix, x)
		}
	}

	// set zero to first row and column if necessary
	if firstRowHasZero {
		setRowZero(matrix, 0)
	}
	if firstColHasZero {
		setColZero(matrix, 0)
	}

	return matrix
}

func setRowZero(matrix [][]int, row int) {
	for x := 0; x < len(matrix[0]); x++ {
		matrix[row][x] = 0
	}
}

func setColZero(matrix [][]int, column int) {
	for y := 0; y < len(matrix); y++ {
		matrix[y][column] = 0
	}
}
