package astrav

import (
	"go/ast"
)

// Named provides an interface for nodes with a name
type Named interface {
	// NodeName returns the name of the node
	NodeName() *Ident
}

// NodeName returns the name of the node
func (s *FuncDecl) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}

	return s.findChildByAstNode(s.Name).(*Ident)
}

// NodeName returns the name of the node
func (s *LabeledStmt) NodeName() *Ident {
	if s.Label == nil {
		return nil
	}
	return s.findChildByAstNode(s.Label).(*Ident)
}

// NodeName returns the name of the node
func (s *BranchStmt) NodeName() *Ident {
	if s.Label == nil {
		return nil
	}
	return s.findChildByAstNode(s.Label).(*Ident)
}

// NodeName returns the name of the node
func (s *ImportSpec) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return s.findChildByAstNode(s.Name).(*Ident)
}

// NodeName returns the name of the node
func (s *File) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return s.findChildByAstNode(s.Name).(*Ident)
}

// NodeName returns the name of the node
func (s *SelectorExpr) NodeName() *Ident {
	if s.Sel == nil {
		return nil
	}
	return s.findChildByAstNode(s.Sel).(*Ident)
}

// NodeName returns the name of the node
func (s *TypeSpec) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return s.findChildByAstNode(s.Name).(*Ident)
}

// NodeName returns the name of the node
func (s *CallExpr) NodeName() *Ident {
	if s.Fun == nil {
		return nil
	}
	switch t := s.Fun.(type) {
	case *ast.Ident:
		return s.findChildByAstNode(t).(*Ident)
	case *ast.ArrayType:
		// node does not exist yet
		return newChild(t.Elt, s.realMe, s.pkg, s.level).(*Ident)
	}
	return nil
}

// NodeName returns the name of the node
func (s *Ident) NodeName() *Ident {
	return s
}

// NodeName returns the name of the node
func (s *Field) NodeName() *Ident {
	if len(s.Names) == 0 {
		return nil
	}
	return s.findChildByAstNode(s.Names[0]).(*Ident)
}
