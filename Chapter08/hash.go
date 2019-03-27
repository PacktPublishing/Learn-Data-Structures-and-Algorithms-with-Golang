//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	//	"strconv"
	"crypto/sha1"
	"hash"
)

//CreateHash method
func CreateHash(byteStr []byte) []byte {
	var hashVal hash.Hash
	hashVal = sha1.New()
	hashVal.Write(byteStr)

	var bytes []byte

	bytes = hashVal.Sum(nil)
	return bytes
}

// Create hash for Multiple Values method
func CreateHashMultiple(byteStr1 []byte, byteStr2 []byte) []byte {
	return xor(CreateHash(byteStr1), CreateHash(byteStr2))
}

// XOR method
func xor(byteStr1 []byte, byteStr2 []byte) []byte {
	var xorbytes []byte
	xorbytes = make([]byte, len(byteStr1))
	var i int
	for i = 0; i < len(byteStr1); i++ {
		xorbytes[i] = byteStr1[i] ^ byteStr2[i]
	}
	return xorbytes
}

// main method
func main() {

	var bytes []byte
	bytes = CreateHashMultiple([]byte("Check"), []byte("Hash"))

	fmt.Printf("%x\n", bytes)
}
