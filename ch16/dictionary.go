//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing bufio, errors, io, math/rand,os, strings and time packages

import (
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
	//"fmt"
)
//WordDictionary interface
type WordDictionary interface {
	ChooseWord() string

	ChooseWords(n int) []string
}
// Words class
type Words struct {
	words []string
}

var (
	ErrLoadFailure = errors.New("Failed to load word dictionary input")
)

var vowels = map[string]string{
	"Á": "A",
	"É": "E",
	"Í": "I",
	"Ó": "O",
	"Ú": "U",
}

//NewWordDictionaryFromTextFile method
func NewWordDictionaryFromTextFile(filenamestr string) (WordDictionary, error) {

	var file *os.File
	var err error
	file, err = os.Open(filenamestr)
	var words []string
	if err != nil {
		return nil, ErrLoadFailure
	}
	defer file.Close()
	words = readInputWords(file)
	rand.Seed(time.Now().UnixNano())
	return &Words{words: words}, nil
}

//readInputWords
func readInputWords(reader io.Reader) []string {
	var words []string

	var bufReader *bufio.Reader
	bufReader = bufio.NewReader(reader)
	for {
		var str string
		var err error

		if str, err = bufReader.ReadString('\n'); err == nil && len(str) > 5 {
			str = strings.TrimSpace(str)
			str = strings.ToUpper(str)
			for invalid, vowel := range vowels {
				str = strings.Replace(str, invalid, vowel, len(str))
			}
			words = append(words, str)
		} else if err != nil {
			break
		}
	}
	return words
}
// ChooseWord method
func (words *Words) ChooseWord() string {

	return words.words[rand.Intn(len(words.words))]
}
// ChooseWords method
func (words *Words) ChooseWords(n int) []string {
	var result []string
	var i int
	for i = 0; i < n; i++ {
		result = append(result, words.ChooseWord())
	}
	return result
}
