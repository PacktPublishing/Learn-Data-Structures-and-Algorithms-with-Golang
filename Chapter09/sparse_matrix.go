///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//List of List
type LOL struct {
	Row    int
	Column int
	Value  float64
}

//Sparse Matrix
type SparseMatrix struct {
	cells []LOL
	shape [2]int
}

// Shape method
func (sparseMatrix *SparseMatrix) Shape() (int, int) {
	return sparseMatrix.shape[0], sparseMatrix.shape[1]
}

// NumNonZero method
func (sparseMatrix *SparseMatrix) NumNonZero() int {
	return len(sparseMatrix.cells)
}

// Less Than method
func LessThan(lol LOL, i int, j int) bool {

	if lol.Row < i && lol.Column < j {

		return true
	}

	return false
}

// Equal method
func Equal(lol LOL, i int, j int) bool {

	if lol.Row == i && lol.Column == j {

		return true
	}

	return false

}

// GetValue method
func (sparseMatrix *SparseMatrix) GetValue(i int, j int) float64 {
	var lol LOL
	for _, lol = range sparseMatrix.cells {
		if LessThan(lol, i, j) {
			continue
		}
		if Equal(lol, i, j) {
			return lol.Value
		}
		return 0.0
	}
	return 0.0
}

//SetValue method
func (sparseMatrix *SparseMatrix) SetValue(i int, j int, value float64) {

	var lol LOL
	var index int
	for index, lol = range sparseMatrix.cells {
		if LessThan(lol, i, j) {
			continue
		}
		if Equal(lol, i, j) {
			sparseMatrix.cells[index].Value = value
			return
		}

		sparseMatrix.cells = append(sparseMatrix.cells, LOL{})
		var k int
		for k = len(sparseMatrix.cells) - 2; k >= index; k-- {
			sparseMatrix.cells[k+1] = sparseMatrix.cells[k]
		}
		sparseMatrix.cells[index] = LOL{
			Row:    i,
			Column: j,
			Value:  value,
		}
		return
	}
	sparseMatrix.cells = append(sparseMatrix.cells, LOL{
		Row:    i,
		Column: j,
		Value:  value,
	})
}

// New SparseMatrix method
func NewSparseMatrix(m int, n int) *SparseMatrix {
	return &SparseMatrix{
		cells: []LOL{},
		shape: [2]int{m, n},
	}
}

// main method
func main() {

	var sparseMatrix *SparseMatrix

	sparseMatrix = NewSparseMatrix(3, 3)

	sparseMatrix.SetValue(1, 1, 2.0)
	sparseMatrix.SetValue(1, 3, 3.0)

	fmt.Println(sparseMatrix)
	fmt.Println(sparseMatrix.NumNonZero())
}
