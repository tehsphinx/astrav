package astrav

import (
	"go/ast"
	"go/build"
	"go/token"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// NewModule creates an analyzer for an entire module or application. The dir argument is
// the folder of the module or application within root. If root only contains the module,
// dir can be left empty.
// Modules are loaded from disc. Support for fs.FS is currently not implemented
func NewModule(dir string) *Module {
	return &Module{
		dir:  dir,
		FSet: token.NewFileSet(),
		Pkgs: map[string]*Package{},
	}
}

// Module represents a Go module or an application.
type Module struct {
	dir string

	FSet     *token.FileSet
	Pkgs     map[string]*Package
	RawFiles map[string]*RawFile

	Packages []*packages.Package
	SSAPkgs  []*ssa.Package
	Graph    *callgraph.Graph
}

// Dir returns the directory of the module.
func (s *Module) Dir() string {
	return s.dir
}

// Load will load the module. DefaultFilterGoPackages is used to filter only Go packages that contain Go code.
// Packages only containing test files are excluded by default. Custom filter functions can be passed instead.
// The filter function should return false to filter a file or directory.
func (s *Module) Load(filterFns ...func(d fs.DirEntry) bool) error {
	if len(filterFns) == 0 {
		filterFns = append(filterFns, DefaultFilterGoPackages)
	}

	paths, fileSources, err := s.loadFiles(filterFns...)
	if err != nil {
		return err
	}

	if r := s.loadPackages(s.dir, paths); r != nil {
		return r
	}

	s.fillRawFiles(fileSources)

	if r := s.processPackages(); r != nil {
		return r
	}

	s.buildCallGraph()

	return nil
}

func (s *Module) buildCallGraph() {
	program, pkgs := ssautil.AllPackages(s.Packages, 0)
	s.SSAPkgs = pkgs

	program.Build()

	s.Graph = cha.CallGraph(program)
}

func (s *Module) processPackages() error {
	for _, pack := range s.Packages {
		if len(pack.CompiledGoFiles) != len(pack.Syntax) {
			return errors.Errorf("file without Go code? %v", pack.CompiledGoFiles)
		}

		files := make(map[string]*ast.File, len(pack.Syntax))
		for i, fileName := range pack.CompiledGoFiles {
			file := pack.Syntax[i]
			files[fileName] = file
		}

		pkgNode := creator(baseNode{
			node: &ast.Package{
				Name:  pack.PkgPath,
				Files: files,
			},
		}).(*Package)

		pkgNode.rawFiles = map[string]*RawFile{}
		for fileName, file := range s.RawFiles {
			dir, _ := path.Split(fileName)
			relativePath := strings.TrimSuffix(strings.TrimPrefix(dir, s.dir), "/")
			if !strings.HasSuffix(pack.PkgPath, relativePath) {
				continue
			}

			pkgNode.rawFiles[fileName] = file
		}
		pkgNode.info = pack.TypesInfo
		pkgNode.typesPkg = pack.Types
		pkgNode.pack = pack

		s.Pkgs[pack.PkgPath] = pkgNode
	}
	return nil
}

func (s *Module) loadPackages(repoRoot string, paths []string) error {
	s.FSet = token.NewFileSet()
	packs, err := packages.Load(&packages.Config{
		Dir: repoRoot,
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports |
			packages.NeedTypes | packages.NeedTypesSizes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedDeps,
		BuildFlags: build.Default.BuildTags,
		Fset:       s.FSet,
		// Overlay:    fileSources,
	}, paths...)
	if err != nil {
		return errors.WithMessagef(err, "failed loading %s", s.dir)
	}
	s.Packages = packs
	return nil
}

// DefaultFilterGoPackages returns the default filter function for Go packages.
func DefaultFilterGoPackages(d fs.DirEntry) bool {
	if d.IsDir() {
		return false
	}
	if !strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), "_test.go") {
		return false
	}
	return true
}

// GetRawFiles a map of file contents
func (s *Module) GetRawFiles() map[string][]byte {
	files := map[string][]byte{}
	for name, file := range s.RawFiles {
		files[name] = file.Source()
	}
	return files
}

func (s *Module) fillRawFiles(fileSources map[string][]byte) {
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
}

// Package returns a package by name
func (s *Module) Package(name string) *Package {
	return s.Pkgs[name]
}

// loadFiles creates a slice of paths which contain Go packages in given baseDir.
func (s *Module) loadFiles(filterFns ...func(d fs.DirEntry) bool) ([]string, map[string][]byte, error) {
	var (
		fileSRCs = map[string][]byte{}
		seen     = map[string]struct{}{}
		paths    []string
	)
	root := os.DirFS(s.dir)
	err := fs.WalkDir(root, ".", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}
		for _, isValid := range filterFns {
			if !isValid(d) {
				return nil
			}
		}

		bts, err := readFile(root, filePath)
		if err != nil {
			return err
		}
		fileSRCs[path.Join(s.dir, filePath)] = bts

		pkgPath := strings.TrimSuffix(filePath, d.Name())
		if _, ok := seen[pkgPath]; ok {
			return nil
		}
		seen[pkgPath] = struct{}{}

		paths = append(paths, path.Join(s.dir, pkgPath))
		return nil
	})
	if err != nil {
		return nil, nil, errors.WithMessage(err, "failed to gather go packages")
	}
	return paths, fileSRCs, nil
}

func readFile(root fs.FS, filePath string) ([]byte, error) {
	file, err := root.Open(filePath)
	if err != nil {
		return nil, err
	}
	bts, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bts, nil
}
