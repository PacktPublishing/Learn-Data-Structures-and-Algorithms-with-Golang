///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//prints the matrix in zig zag fashion
func PrintZigZag(n int) []int {
	var zigzag []int
	zigzag = make([]int, n*n)
	var i int
	i = 0
	var m int
	m = n * 2
	var p int
	for p = 1; p <= m; p++ {
		var x int
		x = p - n
		if x < 0 {
			x = 0
		}
		var y int
		y = p - 1
		if y > n-1 {
			y = n - 1
		}
		var j int
		j = m - p
		if j > p {
			j = p
		}
		var k int
		for k = 0; k < j; k++ {
			if p&1 == 0 {
				zigzag[(x+k)*n+y-k] = i
			} else {
				zigzag[(y-k)*n+x+k] = i
			}
			i++
		}
	}

	return zigzag
}

// main method
func main() {
	var n int
	n = 5
	var length int
	length = 2
	var i int
	var sketch int
	for i, sketch = range PrintZigZag(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}
}
