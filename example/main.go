package main

import (
	"flag"
	"fmt"
	"go/token"
	"log"
	"strings"

	"github.com/tehsphinx/astrav"
)

var (
	file = flag.String("file", "./1/example.go", "file to parse")
)

func main() {
	flag.Parse()

	fs := token.NewFileSet()

	fNode, err := astrav.NewFile(*file, fs)
	if err != nil {
		log.Fatal(err)
	}

	printTrees(fNode)
}

func printTrees(fNode astrav.Node) {
	fNode.Walk(func(node astrav.Node) bool {
		fmt.Printf("%s%T\n", strings.Repeat("\t", node.Level()), node)
		return true
	})
}
