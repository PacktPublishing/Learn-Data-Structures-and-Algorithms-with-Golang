//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
)

// GetWordLadders method
func GetWordLadders(start string, end string, wordList []string) [][]string {
	wordList = RemoveStartWord(wordList, start)
	var wordmap map[string][]string
	wordmap = map[string][]string{}
	var isEndWord bool
	isEndWord = false
	var cnt int
	cnt = 1
	var breadFirstSearch func([]string, []string)
	breadFirstSearch = func(wordList, nodes []string) {
		cnt++
		var newWordList []string
		newWordList = make([]string, 0, len(wordList))
		var newNodes []string
		newNodes = make([]string, 0, len(wordList))
		var w string
		for _, w = range wordList {
			var isEqual bool
			isEqual = false
			var n string
			for _, n = range nodes {
				if isSame(n, w) {
					wordmap[n] = append(wordmap[n], w)
					isEqual = true
				}
			}

			if isEqual {
				newNodes = append(newNodes, w)
				if w == end {
					isEndWord = true
				}
			} else {
				newWordList = append(newWordList, w)
			}
		}

		if isEndWord ||
			len(newWordList) == 0 ||
			len(newNodes) == 0 {
			return
		}

		breadFirstSearch(newWordList, newNodes)
	}

	var nodes []string
	nodes = []string{start}
	breadFirstSearch(wordList, nodes)

	var result [][]string
	result = DepthFirstSearch(cnt, wordmap, isEndWord, start, end)

	return result

}

//DepthFirstSearch method
func DepthFirstSearch(cnt int, wordmap map[string][]string, isEndWord bool, start string, end string) [][]string {
	var result [][]string
	result = [][]string{}
	if !isEndWord {

		return result
	}

	var path []string
	path = make([]string, cnt)
	path[0] = start

	var depthFirstSearch func(int)
	depthFirstSearch = func(idx int) {
		if idx == cnt {
			if path[idx-1] == end {
				result = append(result, CopyString(path))
			}
			return
		}
		var prev string
		prev = path[idx-1]
		var w string
		for _, w = range wordmap[prev] {
			path[idx] = w
			depthFirstSearch(idx + 1)
		}
	}

	depthFirstSearch(1)

	return result
}

// CopyString method
func CopyString(str []string) []string {
	var temp []string
	temp = make([]string, len(str))
	copy(temp, str)
	return temp
}

// RemoveStartWord method
func RemoveStartWord(wordList []string, start string) []string {
	var i int
	var size int
	i, size = 0, len(wordList)
	for ; i < size; i++ {
		if wordList[i] == start {
			break
		}
	}

	if i == size {
		return wordList
	}
	wordList[i] = wordList[size-1]
	return wordList[:size-1]
}

// isSame method
func isSame(a string, b string) bool {
	var same bool
	same = false
	var i int
	for i = range a {
		if a[i] != b[i] {
			if same {
				return false
			}
			same = true
		}
	}

	return true
}

// main method
func main() {

	var start string
	start = "COLD"

	var end string
	end = "WARM"
	var dict []string

	dict = append(dict, "CORD")
	dict = append(dict, "CORM")
	dict = append(dict, "WORM")
	dict = append(dict, "CARD")
	dict = append(dict, "WOLD")
	dict = append(dict, "WORD")
	dict = append(dict, "WARD")
	dict = append(dict, "WARM")
	dict = append(dict, "FARM")

	var wordLadders [][]string
	wordLadders = GetWordLadders(start, end, dict)

	var words []string
	for _, words = range wordLadders {
		//fmt.Println(start)
		var word string
		var i int
		for i, word = range words {

			if i == 0 {
				fmt.Print(word)
			} else {
				fmt.Print("->" + word)
			}
		}
		fmt.Println()
	}
}
