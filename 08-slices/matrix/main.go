package main

import "fmt"

func main() {
	fmt.Println(createMatrix(50, 50))
}

// Before refactor
// func createMatrix(rows, cols int) [][]int {
// 	matrix := make([][]int, 0, rows*cols)

// 	for i := 0; i < rows; i++ {
// 		row := make([]int, 0)
// 		for j := 0; j < cols; j++ {
// 			row = append(row, i*j)
// 		}
// 		matrix = append(matrix, row)
// 	}

// 	return matrix
// }

// After refactor
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}

	return matrix
}
