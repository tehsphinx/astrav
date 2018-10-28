package astrav

import "go/ast"

//SelectorExpr wraps ast.SelectorExpr
type SelectorExpr struct {
	*ast.SelectorExpr
	baseNode
}

//NodeName returns the name of the node
func (s *SelectorExpr) NodeName() *Ident {
	if s.Sel == nil {
		return nil
	}
	return newChild(s.Sel, s.parent, s.level).(*Ident)
}

//PackageName returns the package name
func (s *SelectorExpr) PackageName() *Ident {
	if s.X == nil {
		return nil
	}
	if _, ok := s.X.(*ast.Ident); !ok {
		return nil
	}
	return newChild(s.X, s.parent, s.level).(*Ident)
}
