//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing math,sort,os,csv,io,strconv and fmt packages

import (
	"fmt"
	"sort"
)

// Rune class
type Rune []rune

// Sort interface methods
func (r Rune) Len() int               { return len(r) }
func (r Rune) Swap(i int, j int)      { r[i], r[j] = r[j], r[i] }
func (r Rune) Less(i int, j int) bool { return r[i] < r[j] }

//ConvertStringToRuneSlice
func ConvertStringToRuneSlice(str string) []rune {
	var r []rune
	var value rune
	for _, value = range str {
		r = append(r, value)
	}
	return r
}

// CheckAnagram method
func CheckAnagram(str1 string, str2 string) bool {
	var r1 Rune = ConvertStringToRuneSlice(str1)
	var r2 Rune = ConvertStringToRuneSlice(str2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}

// main method
func main() {

	var s1 string
	s1 = "listen"

	var s2 string
	s2 = "silent"

	var check bool
	check = CheckAnagram(s1, s2)

	fmt.Println(check)

	var s3 string
	s3 = "integral"

	var s4 string
	s4 = "gradient"

	check = CheckAnagram(s3, s4)

	fmt.Println(check)

}
