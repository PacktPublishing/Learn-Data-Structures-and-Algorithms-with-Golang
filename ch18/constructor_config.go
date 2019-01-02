//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing flag and fmt packages
import (
	"bytes"
	"fmt"
	"log"
)

//Table class
type Table struct {
	rows   int
	cols   int
	Logger *log.Logger
}

//WriteOutput method
func (t *Table) WriteOutput() {
	fmt.Println("Writing Output")
	t.Logger.Println("Writing to a file")
}

//TableConfig class
type TableConfig struct {
	rows   int
	cols   int
	logger *log.Logger
}

// New method
func New(tConfig TableConfig) Table {

	if tConfig.logger != nil {
		var buf bytes.Buffer
		tConfig.logger = log.New(&buf, "logger", log.Lshortfile)
	}
	var t Table
	t = Table{tConfig.rows, tConfig.cols, tConfig.logger}
	return t
}

// main method
func main() {

	var buf bytes.Buffer

	var logger *log.Logger

	logger = log.New(&buf, "logger", log.Lshortfile)
	var table Table
	var rows int
	rows = 3
	var cols int
	cols = 4
	table = New(TableConfig{rows: rows, cols: cols, logger: logger})

	fmt.Println(table)

}
