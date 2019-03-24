package astrav

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"net/http"
	"os"
	"strings"
)

// NewFolder creates a new folder with given path. Use ParseFolder to parse ast from go files in path.
// The pkgPath is the import path of the package to be used by types.ParseInfo.
func NewFolder(pkgPath string, dir http.FileSystem) *Folder {
	return &Folder{
		path: pkgPath,
		dir:  dir,
		FSet: token.NewFileSet(),
		Pkgs: map[string]*Package{},
	}
}

// Folder represents a go package folder
type Folder struct {
	path string
	dir  http.FileSystem

	Info     *types.Info
	FSet     *token.FileSet
	Pkgs     map[string]*Package
	Pkg      *types.Package
	RawFiles map[string]*RawFile
}

// ParseFolder will parse all to files in folder. It skips test files.
func (s *Folder) ParseFolder() (map[string]*Package, error) {
	pkgs, fileSources, err := Parse(s.FSet, s.dir, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.AllErrors+parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if err := s.fillRawFiles(fileSources); err != nil {
		return nil, err
	}

	for name, pkg := range pkgs {
		s.Pkgs[name] = creator(baseNode{
			node: pkg,
		}).(*Package)
		s.Pkgs[name].rawFiles = s.RawFiles
	}

	if s.Pkg, err = s.ParseInfo(s.path, s.FSet, s.getFiles()); err != nil {
		return nil, err
	}

	return s.Pkgs, nil
}

// GetRawFiles a map of file contents
func (s *Folder) GetRawFiles() map[string][]byte {
	files := map[string][]byte{}
	for name, file := range s.RawFiles {
		files[name] = file.Source()
	}
	return files
}

// GetPath returns the pkg import path.
func (s *Folder) GetPath() string {
	return s.path
}

func (s *Folder) fillRawFiles(fileSources map[string][]byte) error {
	var err error
	s.RawFiles = map[string]*RawFile{}

	s.FSet.Iterate(func(file *token.File) bool {
		fileSrc, ok := fileSources[file.Name()]
		if !ok {
			return true
		}

		s.RawFiles[file.Name()] = &RawFile{
			File:   file,
			source: fileSrc,
		}
		return true
	})

	return err
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

// Package returns a package by name
func (s *Folder) Package(name string) *Package {
	return s.Pkgs[name]
}
