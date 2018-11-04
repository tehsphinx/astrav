package astrav

import "go/ast"

//Package wraps ast.Package
type Package struct {
	*ast.Package
	baseNode

	rawFiles map[string]*RawFile
	defs     map[*Ident]Node

	filled bool
}

//FuncDeclByName returns a func declaration by name
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

//FuncDeclbyCallExpr returns a function declaration from its usage
func (s *Package) FuncDeclbyCallExpr(node *CallExpr) *FuncDecl {
	ident := node.NodeName()
	if ident == nil {
		return nil
	}

	return s.FuncDeclByName(ident.Name)
}

func (s *Package) fill() {
	if s.filled {
		return
	}
	s.filled = true

	s.defs = map[*Ident]Node{}

	s.Walk(func(node Node) bool {
		switch n := node.(type) {
		case *FuncDecl:
			ident := n.NodeName()
			if ident == nil {
				return true
			}
			s.defs[ident] = n
		}

		return true
	})
}
