//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing bytes,fmt,rand and time packages

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

//  MutliCursalMaze class
type MultiCursalMaze struct {
	cells          []byte
	hor            []byte
	vert           []byte
	cellsByRow     [][]byte
	horWallsByRow  [][]byte
	vertWallsByRow [][]byte
}

// GetMultiCursalMaze method
func GetMultiCursalMaze(rows int, cols int) *MultiCursalMaze {
	var cells []byte
	var hor []byte
	var vert []byte

	cells = make([]byte, rows*cols)
	hor = bytes.Repeat([]byte{'-'}, rows*cols)
	vert = bytes.Repeat([]byte{'|'}, rows*cols)
	var cellsByRow [][]byte
	var horWallsByRow [][]byte
	var vertWallsByRow [][]byte
	cellsByRow = make([][]byte, rows)
	horWallsByRow = make([][]byte, rows)
	vertWallsByRow = make([][]byte, rows)
	var i int
	for i = range horWallsByRow {
		cellsByRow[i] = cells[i*cols : (i+1)*cols]
		horWallsByRow[i] = hor[i*cols : (i+1)*cols]
		vertWallsByRow[i] = vert[i*cols : (i+1)*cols]
	}
	return &MultiCursalMaze{cells, hor, vert, cellsByRow, horWallsByRow, vertWallsByRow}
}

// PrintMultiCursalMaze  method
func (maze *MultiCursalMaze) PrintMultiCursalMaze() {
	var horWall []byte
	var horOpen []byte
	var verWall []byte
	var verOpen []byte
	var rightCorner []byte
	var rightWall []byte
	horWall = []byte("+---")
	horOpen = []byte("+   ")
	verWall = []byte("|   ")
	verOpen = []byte("    ")
	rightCorner = []byte("+\n")
	rightWall = []byte("|\n")
	var b []byte
	var r int
	var hw []byte
	for r, hw = range maze.horWallsByRow {
		var h byte
		for _, h = range hw {
			if h == '-' || r == 0 {
				b = append(b, horWall...)
			} else {
				b = append(b, horOpen...)
			}
		}
		b = append(b, rightCorner...)
		var c int
		var vw byte
		for c, vw = range maze.vertWallsByRow[r] {
			if vw == '|' || c == 0 {
				b = append(b, verWall...)
			} else {
				b = append(b, verOpen...)
			}
			if maze.cellsByRow[r][c] != 0 {
				b[len(b)-2] = maze.cellsByRow[r][c]
			}
		}
		b = append(b, rightWall...)
	}

	for _ = range maze.horWallsByRow[0] {
		b = append(b, horWall...)
	}
	b = append(b, rightCorner...)

	fmt.Println(string(b))
}

//GenerateMultiCursalMaze method
func (maze *MultiCursalMaze) GenerateMultiCursalMaze() {

	maze.RecurseMaze(rand.Intn(len(maze.cellsByRow)), rand.Intn(len(maze.cellsByRow[0])))
}

const (
	up = iota
	down
	right
	left
)

// RecurseMaze method
func (maze *MultiCursalMaze) RecurseMaze(row int, col int) {
	rand.Seed(time.Now().UnixNano())
	maze.cellsByRow[row][col] = ' '
	var wall int
	for _, wall = range rand.Perm(4) {
		switch wall {
		case up:
			if row > 0 && maze.cellsByRow[row-1][col] == 0 {
				maze.horWallsByRow[row][col] = 0
				maze.RecurseMaze(row-1, col)
			}
		case down:
			if row < len(maze.cellsByRow)-1 && maze.cellsByRow[row+1][col] == 0 {
				maze.horWallsByRow[row+1][col] = 0
				maze.RecurseMaze(row+1, col)
			}
		case left:
			if col > 0 && maze.cellsByRow[row][col-1] == 0 {
				maze.vertWallsByRow[row][col] = 0
				maze.RecurseMaze(row, col-1)
			}
		case right:
			if col < len(maze.cellsByRow[0])-1 && maze.cellsByRow[row][col+1] == 0 {
				maze.vertWallsByRow[row][col+1] = 0
				maze.RecurseMaze(row, col+1)
			}
		}
	}
}

//SolveMaze method
func (maze *MultiCursalMaze) SolveMaze(ra int, ca int, rz int, cz int) {
	var rSolve func(ra int, ca int, dir int) bool
	rSolve = func(r int, c int, dir int) bool {
		if r == rz && c == cz {
			maze.cellsByRow[r][c] = 'F'
			return true
		}
		if dir != down && maze.horWallsByRow[r][c] == 0 {
			if rSolve(r-1, c, up) {
				maze.cellsByRow[r][c] = '^'
				maze.horWallsByRow[r][c] = '^'
				return true
			}
		}
		if dir != up && r+1 < len(maze.horWallsByRow) && maze.horWallsByRow[r+1][c] == 0 {
			if rSolve(r+1, c, down) {
				maze.cellsByRow[r][c] = 'v'
				maze.horWallsByRow[r+1][c] = 'v'
				return true
			}
		}
		if dir != left && c+1 < len(maze.vertWallsByRow[0]) && maze.vertWallsByRow[r][c+1] == 0 {
			if rSolve(r, c+1, right) {
				maze.cellsByRow[r][c] = '>'
				maze.vertWallsByRow[r][c+1] = '>'
				return true
			}
		}
		if dir != right && maze.vertWallsByRow[r][c] == 0 {
			if rSolve(r, c-1, left) {
				maze.cellsByRow[r][c] = '<'
				maze.vertWallsByRow[r][c] = '<'
				return true
			}
		}
		return false
	}
	rSolve(ra, ca, -1)
	maze.cellsByRow[ra][ca] = 'S'
}

// main method
func main() {
	var maze *MultiCursalMaze
	const height = 6
	const width = 9
	maze = GetMultiCursalMaze(height, width)
	maze.GenerateMultiCursalMaze()

	maze.SolveMaze(
		rand.Intn(height), rand.Intn(width),
		rand.Intn(height), rand.Intn(width))

	maze.PrintMultiCursalMaze()
}
