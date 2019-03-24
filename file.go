package astrav

import (
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/pkg/errors"
)

// NewFile creates a new file from file path
func NewFile(file string, fset *token.FileSet) (*File, error) {
	f, err := parser.ParseFile(fset, file, nil, parser.AllErrors+parser.ParseComments)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rawFile := NewRawFile(fset.File(f.Pos()), b)

	return NewFileNode(f, rawFile), nil
}
