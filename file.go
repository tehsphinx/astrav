package astrav

import (
	"go/parser"
	"go/token"
	"io/ioutil"
)

// NewFile creates a new file from file path
func NewFile(file string, fset *token.FileSet) (*File, error) {
	f, err := parser.ParseFile(fset, file, nil, parser.AllErrors+parser.ParseComments)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	rawFile := NewRawFile(fset.File(f.Pos()), b)

	return NewFileNode(f, rawFile), nil
}
