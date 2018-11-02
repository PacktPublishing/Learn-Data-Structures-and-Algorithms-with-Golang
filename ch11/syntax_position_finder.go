//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, ast,parser and token package

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// main method
func main() {

	var fileSet *token.FileSet
	fileSet = token.NewFileSet()

	const example = `package position

import (
	  "fmt"
    "go/token"
 )

type pos = token.Pos

const wrongPosition = token.NoPos


func check(position pos) bool {
	return position != wrongPosition
}

func main() {
	fmt.Println(check(wrongPosition) == wrongPositionIsValid())
}
`
	var file *ast.File
	var err error
	file, err = parser.ParseFile(fileSet, "example.go", example, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, declarations := range file.Decls {
		position := declarations.Pos()
		relativePosition := fileSet.Position(position)
		absolutePosition := fileSet.PositionFor(position, false)

		var goKeyWord string
		goKeyWord = "func"

		if gen, ok := declarations.(*ast.GenDecl); ok {
			goKeyWord = gen.Tok.String()
		}

		var lineNumber string
		lineNumber = relativePosition.String()
		if relativePosition != absolutePosition {
			lineNumber += "[" + absolutePosition.String() + "]"
		}

		fmt.Printf("%s: %s\n", lineNumber, goKeyWord)
	}

}
