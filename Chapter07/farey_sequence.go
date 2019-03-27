///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// fraction class
type fraction struct {
	numerator   int
	denominator int
}

// string method
func (frac fraction) String() string {
	return fmt.Sprintf("%d/%d", frac.numerator, frac.denominator)
}

// g method
func g(l fraction, r fraction, num int) {
	var frac fraction
	frac = fraction{l.numerator + r.numerator, l.denominator + r.denominator}
	if frac.denominator <= num {
		g(l, frac, num)
		fmt.Print(frac, " ")
		g(frac, r, num)
	}
}

// main method
func main() {
	var num int
	var l fraction
	var r fraction

	for num = 1; num <= 11; num++ {
		l = fraction{0, 1}
		r = fraction{1, 1}
		fmt.Printf("F(%d): %s ", num, l)
		g(l, r, num)
		fmt.Println(r)
	}
	var primes [1001]bool
	var p int
	for _, p = range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31} {
		for num = p * 2; num <= 1000; num += p {
			primes[num] = true
		}
	}
	var totients [1001]int
	var i int
	for i = range totients {
		totients[i] = 1
	}
	for num = 2; num <= 1000; num++ {
		if !primes[num] {
			totients[num] = num - 1
			var a int
			var f int
			var r int
			for a = num * 2; a <= 1000; a += num {
				f = num - 1
				for r = a / num; r%num == 0; r /= num {
					f *= num
				}
				totients[a] *= f
			}
		}
	}
	var sum int
	for num, sum = 1, 1; num <= 1000; num++ {
		sum += totients[num]
		if num%100 == 0 {
			fmt.Printf("|F(%d)|: %d\n", num, sum)
		}
	}
}
