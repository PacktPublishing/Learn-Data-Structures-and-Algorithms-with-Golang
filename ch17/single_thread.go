//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
)

func finder(arr []string) []string {

	var array []string
	var str string
	fmt.Print("inside finder: ")
	for _, str = range arr {

		if str == "ore" {

			fmt.Print(str+" ")
			array = append(array, str)
		}

	}
	fmt.Println()
	return array

}

func miner(arr []string) []string {
	var str string
	fmt.Print("inside miner : ")
	for _, str = range arr {
		fmt.Print("mined", str," ")

	}
	fmt.Println()
	return arr

}

func smelter(arr []string) []string {
	var str string
	fmt.Print("inside smelter: ")
	for _, str = range arr {
		fmt.Print("smeltered", str, " ")

	}
	fmt.Println()
	return arr

}

func main() {
	var theMine []string
	theMine = []string{"rock", "ore", "ore", "rock", "ore"}
	var foundOre []string
	foundOre = finder(theMine)
	var minedOre []string
	minedOre = miner(foundOre)
	smelter(minedOre)
}
