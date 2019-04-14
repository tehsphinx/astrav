package astrav

import (
	"go/ast"
	"go/types"
)

// Package wraps ast.Package
type Package struct {
	*ast.Package
	baseNode

	rawFiles map[string]*RawFile
	defs     map[*Ident]Node
	info     *types.Info

	filled bool
}

// FuncDeclByName returns a func declaration by name
func (s *Package) FuncDeclByName(name string) *FuncDecl {
	s.fill()

	for k, v := range s.defs {
		if k.Name == name {
			if f, ok := v.(*FuncDecl); ok {
				return f
			}
		}
	}

	return nil
}

// FuncDeclbyCallExpr returns a function declaration from its usage
func (s *Package) FuncDeclbyCallExpr(node *CallExpr) *FuncDecl {
	ident := node.GetIdent()
	if ident == nil {
		return nil
	}

	return s.FuncDeclByName(ident.Name)
}

// GetRawFiles returns the raw files from the package.
func (s *Package) GetRawFiles() map[string][]byte {
	var files = map[string][]byte{}
	for key, file := range s.rawFiles {
		files[key] = file.Source()
	}
	return files
}

func (s *Package) fill() {
	if s.filled {
		return
	}
	s.filled = true

	s.defs = map[*Ident]Node{}

	s.Walk(func(node Node) bool {
		if n, ok := node.(*FuncDecl); ok {
			ident := n.GetIdent()
			if ident == nil {
				return true
			}
			s.defs[ident] = n
		}

		return true
	})
}
