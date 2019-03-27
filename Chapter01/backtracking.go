//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//findElementsWithSum  of k from arr of size
func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {

	var num int = 0

	if addValue > k {
		return -1
	}

	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {

			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}

	var i int
	for i = l; i < size; i++ {

		//fmt.Println(" m", m)
		combinations[m] = l

		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

// main method
func main() {

	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}

	var addedSum int = 18

	var combinations [19]int

	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)

	//fmt.Println(check)

	//var check2 bool = findElement(arr,9)

	//fmt.Println(check2)

}
