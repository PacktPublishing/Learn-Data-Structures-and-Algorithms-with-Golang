//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing testing packages

import (
	"testing"
)

func TestAddition(test *testing.T) {


	cases := []struct{ integer1 , integer2 , resultSum int }{
		{1, 1, 2},
		{1, -1, 0},
		{1, 0, 1},
		{0, 0, 0},
		{3, 2, 1},
	}

	for _, cas := range cases {
		var sum int
		var expected int
		sum = cas.integer1 + cas.integer2
		expected = cas.resultSum
		if sum != expected {
			test.Errorf("%d + %d = %d, expected %d", cas.integer1, cas.integer2, sum, expected)
		}
	}

}
