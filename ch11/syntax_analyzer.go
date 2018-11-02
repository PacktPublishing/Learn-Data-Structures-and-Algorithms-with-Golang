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

	var fileSet *token.FileSet
	fileSet = token.NewFileSet()

	const example = `package sub

import (
	"fmt"
	"strings"
)

func printHello() {
	fmt.Println("Hello World")
}`

	var file *ast.File
	var err error

	file, err = parser.ParseFile(fileSet, "", example, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	var spec *ast.ImportSpec
	for _, spec = range file.Imports {
		fmt.Println(spec.Path.Value)
	}

}
