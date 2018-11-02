//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, parser and token package

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// main method
func main() {

	const example = `
package variablechecking
const cons = 10.0
var Y = g(3.14)*3 + d
`

	var fileSet *token.FileSet
	fileSet = token.NewFileSet()
	var file *ast.File
	var err error
	file, err = parser.ParseFile(fileSet, "example.go", example, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(file, func(node ast.Node) bool {
		var str string
		switch nodeType := node.(type) {
		case *ast.BasicLit:
			str = nodeType.Value
		case *ast.Ident:
			str = nodeType.Name
		}
		if str != "" {
			fmt.Printf("%s:\t%s\n", fileSet.Position(node.Pos()), str)
		}
		return true
	})

}
