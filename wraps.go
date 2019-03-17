package astrav

import "go/ast"

// Comment wraps ast.Comment
type Comment struct {
	*ast.Comment
	baseNode
}

// CommentGroup wraps ast.CommentGroup
type CommentGroup struct {
	*ast.CommentGroup
	baseNode
}

// Field wraps ast.Field
type Field struct {
	*ast.Field
	baseNode
}

// FieldList wraps ast.FieldList
type FieldList struct {
	*ast.FieldList
	baseNode
}

// BadExpr wraps ast.BadExpr
type BadExpr struct {
	*ast.BadExpr
	baseNode
}

// Ident wraps ast.Ident
type Ident struct {
	*ast.Ident
	baseNode
}

// Ellipsis wraps ast.Ellipsis
type Ellipsis struct {
	*ast.Ellipsis
	baseNode
}

// BasicLit wraps ast.BasicLit
type BasicLit struct {
	*ast.BasicLit
	baseNode
}

// FuncLit wraps ast.FuncLit
type FuncLit struct {
	*ast.FuncLit
	baseNode
}

// CompositeLit wraps ast.CompositeLit
type CompositeLit struct {
	*ast.CompositeLit
	baseNode
}

// ParenExpr wraps ast.ParenExpr
type ParenExpr struct {
	*ast.ParenExpr
	baseNode
}

// SelectorExpr wraps ast.SelectorExpr
type SelectorExpr struct {
	*ast.SelectorExpr
	baseNode
}

// PackageName returns the package name
func (s *SelectorExpr) PackageName() *Ident {
	if s.X == nil {
		return nil
	}
	if _, ok := s.X.(*ast.Ident); !ok {
		return nil
	}

	return s.findChildByAstNode(s.X).(*Ident)
}

// IndexExpr wraps ast.IndexExpr
type IndexExpr struct {
	*ast.IndexExpr
	baseNode
}

// SliceExpr wraps ast.SliceExpr
type SliceExpr struct {
	*ast.SliceExpr
	baseNode
}

// TypeAssertExpr wraps ast.TypeAssertExpr
type TypeAssertExpr struct {
	*ast.TypeAssertExpr
	baseNode
}

// CallExpr wraps ast.CallExpr
type CallExpr struct {
	*ast.CallExpr
	baseNode
}

// StarExpr wraps ast.StarExpr
type StarExpr struct {
	*ast.StarExpr
	baseNode
}

// UnaryExpr wraps ast.UnaryExpr
type UnaryExpr struct {
	*ast.UnaryExpr
	baseNode
}

// BinaryExpr wraps ast.BinaryExpr
type BinaryExpr struct {
	*ast.BinaryExpr
	baseNode
}

// X returns left hand side nodes
func (s *BinaryExpr) X() Node {
	return s.findChildByAstNode(s.node.(*ast.BinaryExpr).X)
}

// Y returns left hand side nodes
func (s *BinaryExpr) Y() Node {
	return s.findChildByAstNode(s.node.(*ast.BinaryExpr).Y)
}

// KeyValueExpr wraps ast.KeyValueExpr
type KeyValueExpr struct {
	*ast.KeyValueExpr
	baseNode
}

// ArrayType wraps ast.ArrayType
type ArrayType struct {
	*ast.ArrayType
	baseNode
}

// StructType wraps ast.StructType
type StructType struct {
	*ast.StructType
	baseNode
}

// FuncType wraps ast.FuncType
type FuncType struct {
	*ast.FuncType
	baseNode
}

// Params returns the parameter FieldList
func (s *FuncType) Params() *FieldList {
	if s.FuncType.Params == nil {
		return nil
	}
	return s.findChildByAstNode(s.FuncType.Params).(*FieldList)
}

// Results returns the return parameter FieldList
func (s *FuncType) Results() *FieldList {
	if s.FuncType.Results == nil {
		return nil
	}
	return s.findChildByAstNode(s.FuncType.Results).(*FieldList)
}

// InterfaceType wraps ast.InterfaceType
type InterfaceType struct {
	*ast.InterfaceType
	baseNode
}

// MapType wraps ast.MapType
type MapType struct {
	*ast.MapType
	baseNode
}

// ChanType wraps ast.ChanType
type ChanType struct {
	*ast.ChanType
	baseNode
}

// BadStmt wraps ast.BadStmt
type BadStmt struct {
	*ast.BadStmt
	baseNode
}

// DeclStmt wraps ast.DeclStmt
type DeclStmt struct {
	*ast.DeclStmt
	baseNode
}

// EmptyStmt wraps ast.EmptyStmt
type EmptyStmt struct {
	*ast.EmptyStmt
	baseNode
}

// LabeledStmt wraps ast.LabeledStmt
type LabeledStmt struct {
	*ast.LabeledStmt
	baseNode
}

// ExprStmt wraps ast.ExprStmt
type ExprStmt struct {
	*ast.ExprStmt
	baseNode
}

// SendStmt wraps ast.SendStmt
type SendStmt struct {
	*ast.SendStmt
	baseNode
}

// IncDecStmt wraps ast.IncDecStmt
type IncDecStmt struct {
	*ast.IncDecStmt
	baseNode
}

// AssignStmt wraps ast.AssignStmt
type AssignStmt struct {
	*ast.AssignStmt
	baseNode
}

// RHS returns right hand side nodes
func (s *AssignStmt) RHS() []Node {
	var nodes []Node
	for _, node := range s.node.(*ast.AssignStmt).Rhs {
		nodes = append(nodes, s.findChildByAstNode(node))
	}
	return nodes
}

// LHS returns left hand side nodes
func (s *AssignStmt) LHS() []Node {
	var nodes []Node
	for _, node := range s.node.(*ast.AssignStmt).Lhs {
		nodes = append(nodes, s.findChildByAstNode(node))
	}
	return nodes
}

// GoStmt wraps ast.GoStmt
type GoStmt struct {
	*ast.GoStmt
	baseNode
}

// DeferStmt wraps ast.DeferStmt
type DeferStmt struct {
	*ast.DeferStmt
	baseNode
}

// ReturnStmt wraps ast.ReturnStmt
type ReturnStmt struct {
	*ast.ReturnStmt
	baseNode
}

// BranchStmt wraps ast.BranchStmt
type BranchStmt struct {
	*ast.BranchStmt
	baseNode
}

// BlockStmt wraps ast.BlockStmt
type BlockStmt struct {
	*ast.BlockStmt
	baseNode
}

// IfStmt wraps ast.IfStmt
type IfStmt struct {
	*ast.IfStmt
	baseNode
}

// CaseClause wraps ast.CaseClause
type CaseClause struct {
	*ast.CaseClause
	baseNode
}

// SwitchStmt wraps ast.SwitchStmt
type SwitchStmt struct {
	*ast.SwitchStmt
	baseNode
}

// TypeSwitchStmt wraps ast.TypeSwitchStmt
type TypeSwitchStmt struct {
	*ast.TypeSwitchStmt
	baseNode
}

// CommClause wraps ast.CommClause
type CommClause struct {
	*ast.CommClause
	baseNode
}

// SelectStmt wraps ast.SelectStmt
type SelectStmt struct {
	*ast.SelectStmt
	baseNode
}

// ForStmt wraps ast.ForStmt
type ForStmt struct {
	*ast.ForStmt
	baseNode
}

// Init returns the init node of the for loop
func (s *ForStmt) Init() Node {
	if s.ForStmt.Init == nil {
		return nil
	}
	return s.findChildByAstNode(s.ForStmt.Init)
}

// Cond returns the condition node of the for loop
func (s *ForStmt) Cond() Node {
	if s.ForStmt.Cond == nil {
		return nil
	}
	return s.findChildByAstNode(s.ForStmt.Cond)
}

// Post returns the post iteration node of the for loop
func (s *ForStmt) Post() Node {
	if s.ForStmt.Post == nil {
		return nil
	}
	return s.findChildByAstNode(s.ForStmt.Post)
}

// RangeStmt wraps ast.RangeStmt
type RangeStmt struct {
	*ast.RangeStmt
	baseNode
}

// Key returns the key of the range statment
func (s *RangeStmt) Key() Node {
	if s.RangeStmt.Key == nil {
		return nil
	}
	return s.findChildByAstNode(s.RangeStmt.Key)
}

// Value returns the value of the range statment
func (s *RangeStmt) Value() Node {
	if s.RangeStmt.Value == nil {
		return nil
	}
	return s.findChildByAstNode(s.RangeStmt.Value)
}

// X returns the variable to range over
func (s *RangeStmt) X() Node {
	if s.RangeStmt.X == nil {
		return nil
	}
	return s.findChildByAstNode(s.RangeStmt.X)
}

// ImportSpec wraps ast.ImportSpec
type ImportSpec struct {
	*ast.ImportSpec
	baseNode
}

// ValueSpec wraps ast.ValueSpec
type ValueSpec struct {
	*ast.ValueSpec
	baseNode
}

// TypeSpec wraps ast.TypeSpec
type TypeSpec struct {
	*ast.TypeSpec
	baseNode
}

// BadDecl wraps ast.BadDecl
type BadDecl struct {
	*ast.BadDecl
	baseNode
}

// GenDecl wraps ast.GenDecl
type GenDecl struct {
	*ast.GenDecl
	baseNode
}

// FuncDecl wraps ast.FuncDecl
type FuncDecl struct {
	*ast.FuncDecl
	baseNode
}

// Params returns the parameter FieldList
func (s *FuncDecl) Params() *FieldList {
	return s.ChildByNodeType(NodeTypeFuncType).(*FuncType).Params()
}

// Results returns the return parameter FieldList
func (s *FuncDecl) Results() *FieldList {
	return s.ChildByNodeType(NodeTypeFuncType).(*FuncType).Results()
}

// File wraps ast.File
type File struct {
	*ast.File
	baseNode
}
