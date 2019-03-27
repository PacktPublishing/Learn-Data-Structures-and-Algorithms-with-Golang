//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and strconv package
import (
	"fmt"
	"strconv"
)

// look_say method
func look_say(str string) (rstr string) {
	var cbyte byte
	cbyte = str[0]
	var inc int
	inc = 1
	var i int
	for i = 1; i < len(str); i++ {
		var dbyte byte
		dbyte = str[i]
		if dbyte == cbyte {
			inc++
			continue
		}
		rstr = rstr + strconv.Itoa(inc) + string(cbyte)
		cbyte = dbyte
		inc = 1
	}
	return rstr + strconv.Itoa(inc) + string(cbyte)
}

// main method
func main() {
	var str string
	str = "1"
	fmt.Println(str)
	var i int
	for i = 0; i < 8; i++ {
		str = look_say(str)
		fmt.Println(str)
	}
}
