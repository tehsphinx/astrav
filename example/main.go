package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/dbg"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "file_to_parse.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	fNode := astrav.New(f)

	printTrees(fNode, f)
}

func printTrees(fNode astrav.Node, f *ast.File) {
	fNode.Walk(func(node astrav.Node) bool {
		dbg.Green(fmt.Sprintf("%s%T", strings.Repeat("\t", node.Level()), node))
		//for _, p := range node.Parents() {
		//	fmt.Printf("%T;", p)
		//}
		//fmt.Println("")
		return true
	})

	fmt.Println("\n\n")

	var v visitor
	ast.Walk(v, f)
}

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}
