package main

import (
	"fmt"
	"os"

	"github.com/adev73/dbml-go/parser"
	"github.com/adev73/dbml-go/scanner"
)

func main() {
	f, _ := os.Open("test.dbml")
	s := scanner.NewScanner(f)
	parser := parser.NewParser(s)
	dbml, err := parser.Parse()
	fmt.Printf("%#v, %v", dbml, err)
}
