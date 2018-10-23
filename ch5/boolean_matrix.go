///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//changeMatrix method
func changeMatrix(matrix [3][3]int) [3][3]int {
	var i int
	var j int
	var Rows [3]int
	var Columns [3]int

	var matrixChanged [3][3]int

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if matrix[i][j] == 1 {
				Rows[i] = 1
				Columns[j] = 1
			}

		}
	}

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if Rows[i] == 1 || Columns[j] == 1 {
				matrixChanged[i][j] = 1
			}

		}
	}

	return matrixChanged

}

//printMatrix method
func printMatrix(matrix [3][3]int) {
	var i int
	var j int
	//var k int
	for i = 0; i < 3; i++ {

		for j = 0; j < 3; j++ {

			fmt.Printf("%d", matrix[i][j])

		}
		fmt.Printf("\n")
	}

}

//main method
func main() {

	var matrix = [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	printMatrix(matrix)

	matrix = changeMatrix(matrix)

	printMatrix(matrix)

}
