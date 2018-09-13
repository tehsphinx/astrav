package astrav

import (
	"go/parser"
	"go/token"
	"os"
	"strings"
)

//NewFolder creates a new folder with given path. Use ParseFolder to parse ast from go files in path.
func NewFolder(path string) *Folder {
	return &Folder{
		path: path,
		FSet: token.NewFileSet(),
		Pkgs: map[string]*Package{},
	}
}

//Folder represents a go package folder
type Folder struct {
	path string
	FSet *token.FileSet
	Pkgs map[string]*Package
}

//ParseFolder will parse all to files in folder. It skips test files.
func (s *Folder) ParseFolder() (map[string]*Package, error) {
	pkgs, err := parser.ParseDir(s.FSet, s.path, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.AllErrors)

	if err != nil {
		return nil, err
	}

	for name, pkg := range pkgs {
		s.Pkgs[name] = New(pkg).(*Package)
	}
	return s.Pkgs, nil
}

//Package returns a package by name
func (s *Folder) Package(name string) *Package {
	return s.Pkgs[name]
}
