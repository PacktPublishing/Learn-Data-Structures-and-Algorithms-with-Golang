///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"math/rand"
)

//main method
func main() {

	var threedarray [2][2][2]int

	var i int

	var j int

	var k int

	for i = 0; i < 2; i++ {

		for j = 0; j < 2; j++ {

			for k = 0; k < 2; k++ {

				threedarray[i][j][k] = rand.Intn(3)
			}
		}
	}

	fmt.Println(threedarray)
}
