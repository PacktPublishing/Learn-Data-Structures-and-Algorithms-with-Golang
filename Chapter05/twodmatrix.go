///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// add method
func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var sum [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

// subtract method
func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var difference [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference

}

// multiply method
func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var n int
	var product [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			var productSum int = 0
			for n = 0; n < 2; n++ {
				productSum = productSum + matrix1[l][n]*matrix2[n][m]

			}
			product[l][m] = productSum
		}
	}
	return product
}

// transpose method
func transpose(matrix1 [2][2]int) [2][2]int {
	var m int
	var l int
	var transMatrix [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

// determinant method
func determinant(matrix1 [2][2]int) float64 {

	var det float64

	det = det + float64(((matrix1[0][0] * matrix1[1][1]) - (matrix1[0][1] * matrix1[1][0])))

	return det
}
// inverse method
func inverse(matrix [2][2]int) [][]float64 {

	var det float64

	det = determinant(matrix)

	var invmatrix float64
	invmatrix[0][0] = matrix[1][1] / det
	invmatrix[0][1] = -1 * matrix[0][1] / det
	invmatrix[1][0] = -1 * matrix[1][0] / det
	invmatrix[1][1] = matrix[0][0] / det

	return invmatrix

}

//main method
func main() {

	var matrix1 = [2][2]int{
		{4, 5},
		{1, 2}}
	var matrix2 = [2][2]int{
		{6, 7},
		{3, 4}}
	var sum [2][2]int
	sum = add(matrix1, matrix2)

	fmt.Println(sum)

	var difference [2][2]int
	difference = subtract(matrix1, matrix2)

	fmt.Println(difference)

	var product [2][2]int
	product = multiply(matrix1, matrix2)

	fmt.Println(product)

}
