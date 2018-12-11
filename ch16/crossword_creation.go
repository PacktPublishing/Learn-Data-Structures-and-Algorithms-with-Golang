//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Direction int

var (
	ErrWrongWord             = errors.New("Wrong words ")
	ErrCrosswordNotGenerated = errors.New("crossword not generated")
)

const (
	Horizontal Direction = iota
	Vertical
	NumRetries      int = 10
	PositionRetries int = 1000
)

var alphabets = []rune("ABCDEFGHIJKLMNÑOPQRSTUVWXYZ")

type Point struct {
	XCoordinate int
	YCoordinate int
}

type Output struct {
	Word   string
	Direct Direction
	Pos    Point
}

type Solution struct {
	Width  int
	Height int
	Words  []Output
}

type CrosswordSlot struct {
	Point
	Letter rune
}

func (p Point) Equals(other Point) bool {
	return p.XCoordinate == other.XCoordinate && p.YCoordinate == other.YCoordinate
}

//GetSlots method
func (output Output) GetSlots() []CrosswordSlot {
	var result []CrosswordSlot
	var i int
	var l rune
	for i, l = range output.Word {
		var p Point
		p = output.Pos
		if output.Direct == Vertical {
			p.YCoordinate += i
		} else {
			p.XCoordinate += i
		}
		result = append(result, CrosswordSlot{Point: p, Letter: l})
	}
	return result
}

//CheckSimilar method
func (output Output) CheckSimilar(other Output) bool {
	var slots []CrosswordSlot
	slots = output.GetSlots()
	var otherSlot CrosswordSlot
	for _, otherSlot = range other.GetSlots() {
		var mySlot CrosswordSlot
		for _, mySlot = range slots {
			if mySlot.Point.Equals(otherSlot.Point) && mySlot.Letter != otherSlot.Letter {
				return true
			}
		}
	}
	return false
}

//CheckSolution method
func (s *Solution) CheckSolution(output Output) bool {
	if output.Direct == Vertical {
		if output.Pos.YCoordinate+len(output.Word) > s.Height {
			return false
		}
	} else {
		if output.Pos.XCoordinate+len(output.Word) > s.Width {
			return false
		}
	}
	return true
}

//CheckExists method
func (s *Solution) CheckExists(output Output) bool {
	var word Output
	for _, word = range s.Words {
		if word.CheckSimilar(output) {
			return true
		}
	}
	return false
}

// AddWord method
func (s *Solution) AddWord(word string) error {
	if len(word) > s.Height && len(word) > s.Width {
		return ErrWrongWord
	}
	var try int
	for try = 0; try < PositionRetries; try++ {
		var p Point
		p = Point{rand.Intn(s.Width), rand.Intn(s.Height)}
		direct := Vertical
		if rand.Int()%2 == 0 {
			direct = Horizontal
		}
		var output Output
		output = Output{Word: strings.ToUpper(word), Pos: p, Direct: direct}
		if s.CheckSolution(output) && !s.CheckExists(output) {
			s.Words = append(s.Words, output)
			return nil
		}
	}
	return ErrWrongWord
}

// GetCrossWordSolution
func GetCrossWordSolution(height int, width int, numWords int, dict WordDictionary) (*Solution, error) {
	if height < 1 || width < 1 {
		return nil, ErrCrosswordNotGenerated
	}
	var s *Solution
	s = &Solution{Width: width, Height: height}
	var i int
MainLoop:
	for i = 0; i < numWords; i++ {
		var try int
		for try = 0; try < NumRetries; try++ {
			var e error
			if e = s.AddWord(dict.ChooseWord()); e == nil {
				continue MainLoop
			}
		}
		return nil, ErrCrosswordNotGenerated
	}
	return s, nil
}

// main method
func main() {

	var dict WordDictionary

	var err error

	dict, err = NewWordDictionaryFromTextFile("dict_list.txt")

	if err != nil {
		panic(err)
	}

	var height int

	height = 13
	var width int
	width = 13

	var numwords int
	numwords = 7

	var solution *Solution

	solution, err = GetCrossWordSolution(height, width, numwords, dict)

	if err != nil {
		panic(err)
	}

	fmt.Println("CrossWord  Solution", solution.Words)

}
