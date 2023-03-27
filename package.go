package astrav

import (
	"go/ast"
	"go/token"
	"go/types"
	"math"

	"golang.org/x/tools/go/packages"
)

// Package wraps ast.Package
type Package struct {
	*ast.Package
	baseNode

	rawFiles map[string]*RawFile
	defs     map[*Ident]Node
	info     *types.Info
	typesPkg *types.Package
	pack     *packages.Package

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

// // Pos returns the position of the package.
// func (s *Package) Pos() token.Pos {
// 	s.node.(*ast.Package).Pos()
// 	return s.pack.Types.Scope().Pos()
// }

// GetScope returns custom scope.
func (s *Package) GetScope() (Node, *types.Scope) {
	minPos, maxPos := 0, 0
	if len(s.rawFiles) != 0 {
		minPos, maxPos = math.MaxInt, 0
	}

	for _, file := range s.rawFiles {
		pos := file.Base()
		end := pos + file.Size()

		if pos < minPos {
			minPos = pos
		}
		if maxPos < end {
			maxPos = end
		}
	}
	scope := types.NewScope(nil, token.Pos(minPos), token.Pos(maxPos), "")
	return s, scope
}
