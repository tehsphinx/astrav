package astrav

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
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
	Pkg  *types.Package
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

	if s.Pkg, err = ParseInfo(s.path, s.FSet, s.getFiles()); err != nil {
		return nil, err
	}

	return s.Pkgs, nil
}

func (s *Folder) getFiles() []*ast.File {
	var files []*ast.File
	for _, pkg := range s.Pkgs {
		for _, file := range pkg.Files {
			files = append(files, file)
		}
	}
	return files
}

//Package returns a package by name
func (s *Folder) Package(name string) *Package {
	return s.Pkgs[name]
}
