//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and math packages
import (
	"fmt"
	"math"
)

var FirstPeg *[]int
var ThirdPeg *[]int
var SecondPeg *[]int

//InitialisePegs method
func InitialisePegs(num int) {
	FirstPeg = CreateNPegs(num)
	ThirdPeg = &([]int{})
	SecondPeg = &([]int{})
	return
}

//CreateNPegs method
func CreateNPegs(num int) *[]int {
	var disks []int
	disks = []int{num}
	var i int
	for i = 1; i < num; i++ {
		disks = append(disks, (num - i))
	}
	return &disks
}

// HanoiAlgo method
func HanoiAlgo(n int, first *[]int, second *[]int, third *[]int) {
	if n > 0 {
		HanoiAlgo(n-1, first, third, second)
		if len(*first) > 0 {
			var disk int
			disk = (*first)[len(*first)-1]
			*first = (*first)[:(len(*first) - 1)]
			*third = append((*third), disk)
      fmt.Println("First Peg disks", *first)
      fmt.Println("Second Peg disks",*second)
			fmt.Println("third Peg disks",*third)
		}
		HanoiAlgo(n-1, second, first, third)
	}
}

// RunAlgo method
func RunAlgo() {
	var order float64
	order = math.Exp2(float64(len(*(FirstPeg)))) - 1
	fmt.Println("Order of Hanoi Algo", *(FirstPeg), " is ", order)
	HanoiAlgo(len(*(FirstPeg)), FirstPeg, SecondPeg, ThirdPeg)
}

// main method
func main() {

	InitialisePegs(5)
	RunAlgo()
}
