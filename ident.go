package astrav

import (
	"go/ast"
)

// Identifier defines a node element with an ident attached to it or ident itself
type Identifier interface {
	GetIdent() *Ident
}

// GetIdent returns the name of the node
func (s *FuncDecl) GetIdent() *Ident {
	if s.Name == nil {
		return nil
	}

	return s.findChildByAstNode(s.Name).(*Ident)
}

// GetIdent returns the name of the node
func (s *LabeledStmt) GetIdent() *Ident {
	if s.Label == nil {
		return nil
	}
	return s.findChildByAstNode(s.Label).(*Ident)
}

// GetIdent returns the name of the node
func (s *BranchStmt) GetIdent() *Ident {
	if s.Label == nil {
		return nil
	}
	return s.findChildByAstNode(s.Label).(*Ident)
}

// GetIdent returns the name of the node
func (s *ImportSpec) GetIdent() *Ident {
	if s.Name == nil {
		return nil
	}
	return s.findChildByAstNode(s.Name).(*Ident)
}

// GetIdent returns the name of the node
func (s *File) GetIdent() *Ident {
	if s.Name == nil {
		return nil
	}
	return s.findChildByAstNode(s.Name).(*Ident)
}

// GetIdent returns the name of the node
func (s *SelectorExpr) GetIdent() *Ident {
	if s.Sel == nil {
		return nil
	}
	return s.findChildByAstNode(s.Sel).(*Ident)
}

// GetIdent returns the name of the node
func (s *TypeSpec) GetIdent() *Ident {
	if s.Name == nil {
		return nil
	}
	return s.findChildByAstNode(s.Name).(*Ident)
}

// GetIdent returns the name of the node
func (s *CallExpr) GetIdent() *Ident {
	if s.Fun == nil {
		return nil
	}
	switch t := s.Fun.(type) {
	case *ast.Ident:
		return s.findChildByAstNode(t).(*Ident)
	case *ast.ArrayType:
		arrType := s.findChildByAstNode(t).(*ArrayType)
		return arrType.GetIdent()
	}
	return nil
}

// GetIdent returns the name of the node
func (s *Ident) GetIdent() *Ident {
	return s
}

// GetIdent returns the name of the node
func (s *Field) GetIdent() *Ident {
	if len(s.Names) == 0 {
		return nil
	}
	return s.findChildByAstNode(s.Names[0]).(*Ident)
}

// GetIdent returns the name of the node
func (s *ArrayType) GetIdent() *Ident {
	if s.Elt == nil {
		return nil
	}
	ident, ok := s.Elt.(*ast.Ident)
	if !ok {
		return nil
	}
	return s.findChildByAstNode(ident).(*Ident)
}
