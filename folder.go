package astrav

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

	studentName string

	FSet     *token.FileSet
	Pkgs     map[string]*Package
	Pkg      *types.Package
	RawFiles map[string]*RawFile
}

//ParseFolder will parse all to files in folder. It skips test files.
func (s *Folder) ParseFolder() (map[string]*Package, error) {
	s.getStudentName()

	pkgs, err := parser.ParseDir(s.FSet, s.path, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.AllErrors+parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if err := s.fillRawFiles(); err != nil {
		return nil, err
	}

	for name, pkg := range pkgs {
		s.Pkgs[name] = creator(baseNode{
			node: pkg,
		}).(*Package)
		s.Pkgs[name].rawFiles = s.RawFiles
	}

	if s.Pkg, err = ParseInfo(s.path, s.FSet, s.getFiles()); err != nil {
		return nil, err
	}

	return s.Pkgs, nil
}

//GetRawFiles a map of file contents
func (s *Folder) GetRawFiles() map[string][]byte {
	files := map[string][]byte{}
	for name, file := range s.RawFiles {
		files[name] = file.Source()
	}
	return files
}

var student = regexp.MustCompile("users/([^/]*)/go/")

func (s *Folder) getStudentName() {
	path, err := filepath.Abs(s.path)
	if err != nil {
		return
	}

	submatch := student.FindStringSubmatch(path)
	if 1 < len(submatch) {
		s.studentName = submatch[1]
	}
}

func (s *Folder) fillRawFiles() error {
	var err error
	s.RawFiles = map[string]*RawFile{}

	s.FSet.Iterate(func(file *token.File) bool {
		if !strings.HasSuffix(file.Name(), "_test.go") {
			b, r := ioutil.ReadFile(file.Name())
			if r != nil {
				err = r
				return false
			}

			s.RawFiles[file.Name()] = &RawFile{
				File:   file,
				source: b,
			}
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

//Package returns a package by name
func (s *Folder) Package(name string) *Package {
	return s.Pkgs[name]
}

//StudentName returns the students name. ParseFolder has to be run before to parse students name from path.
func (s *Folder) StudentName() string {
	return s.studentName
}
