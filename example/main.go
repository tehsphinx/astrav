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
	file = flag.String("file", "file_to_parse.go", "file to parse")
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
		//dbg.Green(string(node.GetSource()))
		//for _, p := range node.Parents() {
		//	fmt.Printf("%T;", p)
		//}
		//fmt.Println("")
		return true
	})
}
