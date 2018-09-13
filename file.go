package astrav

import (
	"go/parser"
	"go/token"
)

//NewFile creates a new file from file path
func NewFile(file string, fset *token.FileSet) (*File, error) {
	f, err := parser.ParseFile(fset, file, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return New(f).(*File), nil
}
