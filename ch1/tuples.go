//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//gets the powerseries of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int) {

	return a * a, a * a * a

}

func powerSeriesN(a int) (square int, cube int) {

	square = a * a

	cube = square * a

	return

}

func powerSeriesE(a int) (int, int, error) {

	var square int = a * a

	var cube int = square * a

	return square, cube, nil

}

// main method
func main() {

	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, "Cube", cube)

	fmt.Println(powerSeriesN(4))

	fmt.Println(powerSeriesE(5))
}
