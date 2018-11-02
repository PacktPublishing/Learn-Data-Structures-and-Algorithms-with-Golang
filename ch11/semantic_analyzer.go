//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, parser and token package
import (
	//"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

// main method
func main() {

	var fileSet *token.FileSet
	fileSet = token.NewFileSet()

	var file *ast.File
	var err error
	file, err = parser.ParseFile(fileSet, "check_type.go", nil, 0)
	if err != nil {

	}
	printer.Fprint(os.Stdout, fileSet, file)

	ast.Walk(new(MethodChecker), file)

}

//Method Checker class
type MethodChecker struct {
}

func (checker *MethodChecker) Visit(astNode ast.Node) (visitor ast.Visitor) {

	switch nodeType := astNode.(type) {
	case *ast.FuncDecl:
		nodeType.Name = ast.NewIdent(strings.Title(nodeType.Name.Name))
	}

	return checker
}
