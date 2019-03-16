package astrav

import (
	"go/ast"
)

// Named provides an interface for nodes with a name
type Named interface {
	NodeName() *Ident
}

// NodeName returns the name of the node
func (s *FuncDecl) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}

	return newChild(s.Name, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *LabeledStmt) NodeName() *Ident {
	if s.Label == nil {
		return nil
	}
	return newChild(s.Label, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *BranchStmt) NodeName() *Ident {
	if s.Label == nil {
		return nil
	}
	return newChild(s.Label, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *ImportSpec) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return newChild(s.Name, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *File) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return newChild(s.Name, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *SelectorExpr) NodeName() *Ident {
	if s.Sel == nil {
		return nil
	}
	return newChild(s.Sel, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *TypeSpec) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return newChild(s.Name, s.realMe, s.pkg, s.level).(*Ident)
}

// NodeName returns the name of the node
func (s *CallExpr) NodeName() *Ident {
	if s.Fun == nil {
		return nil
	}
	switch t := s.Fun.(type) {
	case *ast.Ident:
		return newChild(t, s.realMe, s.pkg, s.level).(*Ident)
	case *ast.ArrayType:
		return newChild(t.Elt, s.realMe, s.pkg, s.level).(*Ident)
	}
	return nil
}

// NodeName returns the name of the node
func (s *Ident) NodeName() *Ident {
	return s
}
