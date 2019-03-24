package astrav

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

// ParseFileSystem calls ParseFile for all files with names ending in ".go" in the
// http.FileSystem specified by path and returns a map of package name -> package
// AST with all the packages found.
//
// If filter != nil, only the files with os.FileInfo entries passing through
// the filter (and ending in ".go") are considered. The mode bits are passed
// to ParseFile unchanged. Position information is recorded in fset, which
// must not be nil.
//
// If the directory couldn't be read, a nil map and the respective error are
// returned. If a parse error occurred, a non-nil but incomplete map and the
// first error encountered are returned.
//
func Parse(fset *token.FileSet, root http.FileSystem, dir string, filter func(os.FileInfo) bool,
	mode parser.Mode) (pkgs map[string]*ast.Package, fileSources map[string][]byte, first error) {
	fd, err := root.Open(dir)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer fd.Close()

	list, err := fd.Readdir(-1)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	pkgs = make(map[string]*ast.Package)
	fileSources = make(map[string][]byte)
	for _, d := range list {
		filename := d.Name()
		if !strings.HasSuffix(filename, ".go") || filter != nil && !filter(d) {
			continue
		}
		fileBytes, err := getSource(path.Join(dir, filename), root)
		if err != nil {
			if first == nil {
				first = err
			}
			continue
		}
		fileSources[filename] = fileBytes

		src, err := parser.ParseFile(fset, filename, fileBytes, mode)
		if err != nil {
			if first == nil {
				first = errors.WithStack(err)
			}
			continue
		}
		name := src.Name.Name
		pkg, found := pkgs[name]
		if !found {
			pkg = &ast.Package{
				Name:  name,
				Files: make(map[string]*ast.File),
			}
			pkgs[name] = pkg
		}
		pkg.Files[filename] = src
	}

	return pkgs, fileSources, first
}

func getSource(path string, dir http.FileSystem) ([]byte, error) {
	f, err := dir.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	bytes, err := ioutil.ReadAll(f)
	return bytes, errors.WithStack(err)
}
