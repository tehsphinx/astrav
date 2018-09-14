package astrav

import "go/ast"

//Comment wraps ast.Comment
type Comment struct {
	*ast.Comment
	baseNode
}

//CommentGroup wraps ast.Comment
type CommentGroup struct {
	*ast.CommentGroup
	baseNode
}

//Field wraps ast.Comment
type Field struct {
	*ast.Field
	baseNode
}

//FieldList wraps ast.Comment
type FieldList struct {
	*ast.FieldList
	baseNode
}

//BadExpr wraps ast.Comment
type BadExpr struct {
	*ast.BadExpr
	baseNode
}

//Ident wraps ast.Comment
type Ident struct {
	*ast.Ident
	baseNode
}

//Ellipsis wraps ast.Comment
type Ellipsis struct {
	*ast.Ellipsis
	baseNode
}

//BasicLit wraps ast.Comment
type BasicLit struct {
	*ast.BasicLit
	baseNode
}

//FuncLit wraps ast.Comment
type FuncLit struct {
	*ast.FuncLit
	baseNode
}

//CompositeLit wraps ast.Comment
type CompositeLit struct {
	*ast.CompositeLit
	baseNode
}

//ParenExpr wraps ast.Comment
type ParenExpr struct {
	*ast.ParenExpr
	baseNode
}

//SelectorExpr wraps ast.Comment
type SelectorExpr struct {
	*ast.SelectorExpr
	baseNode
}

//IndexExpr wraps ast.Comment
type IndexExpr struct {
	*ast.IndexExpr
	baseNode
}

//SliceExpr wraps ast.Comment
type SliceExpr struct {
	*ast.SliceExpr
	baseNode
}

//TypeAssertExpr wraps ast.Comment
type TypeAssertExpr struct {
	*ast.TypeAssertExpr
	baseNode
}

//CallExpr wraps ast.Comment
type CallExpr struct {
	*ast.CallExpr
	baseNode
}

//StarExpr wraps ast.Comment
type StarExpr struct {
	*ast.StarExpr
	baseNode
}

//UnaryExpr wraps ast.Comment
type UnaryExpr struct {
	*ast.UnaryExpr
	baseNode
}

//BinaryExpr wraps ast.Comment
type BinaryExpr struct {
	*ast.BinaryExpr
	baseNode
}

//KeyValueExpr wraps ast.Comment
type KeyValueExpr struct {
	*ast.KeyValueExpr
	baseNode
}

//ArrayType wraps ast.Comment
type ArrayType struct {
	*ast.ArrayType
	baseNode
}

//StructType wraps ast.Comment
type StructType struct {
	*ast.StructType
	baseNode
}

//FuncType wraps ast.Comment
type FuncType struct {
	*ast.FuncType
	baseNode
}

//InterfaceType wraps ast.Comment
type InterfaceType struct {
	*ast.InterfaceType
	baseNode
}

//MapType wraps ast.Comment
type MapType struct {
	*ast.MapType
	baseNode
}

//ChanType wraps ast.Comment
type ChanType struct {
	*ast.ChanType
	baseNode
}

//BadStmt wraps ast.Comment
type BadStmt struct {
	*ast.BadStmt
	baseNode
}

//DeclStmt wraps ast.Comment
type DeclStmt struct {
	*ast.DeclStmt
	baseNode
}

//EmptyStmt wraps ast.Comment
type EmptyStmt struct {
	*ast.EmptyStmt
	baseNode
}

//LabeledStmt wraps ast.Comment
type LabeledStmt struct {
	*ast.LabeledStmt
	baseNode
}

//ExprStmt wraps ast.Comment
type ExprStmt struct {
	*ast.ExprStmt
	baseNode
}

//SendStmt wraps ast.Comment
type SendStmt struct {
	*ast.SendStmt
	baseNode
}

//IncDecStmt wraps ast.Comment
type IncDecStmt struct {
	*ast.IncDecStmt
	baseNode
}

//AssignStmt wraps ast.Comment
type AssignStmt struct {
	*ast.AssignStmt
	baseNode
}

//GoStmt wraps ast.Comment
type GoStmt struct {
	*ast.GoStmt
	baseNode
}

//DeferStmt wraps ast.Comment
type DeferStmt struct {
	*ast.DeferStmt
	baseNode
}

//ReturnStmt wraps ast.Comment
type ReturnStmt struct {
	*ast.ReturnStmt
	baseNode
}

//BranchStmt wraps ast.Comment
type BranchStmt struct {
	*ast.BranchStmt
	baseNode
}

//BlockStmt wraps ast.Comment
type BlockStmt struct {
	*ast.BlockStmt
	baseNode
}

//IfStmt wraps ast.Comment
type IfStmt struct {
	*ast.IfStmt
	baseNode
}

//CaseClause wraps ast.Comment
type CaseClause struct {
	*ast.CaseClause
	baseNode
}

//SwitchStmt wraps ast.Comment
type SwitchStmt struct {
	*ast.SwitchStmt
	baseNode
}

//TypeSwitchStmt wraps ast.Comment
type TypeSwitchStmt struct {
	*ast.TypeSwitchStmt
	baseNode
}

//CommClause wraps ast.Comment
type CommClause struct {
	*ast.CommClause
	baseNode
}

//SelectStmt wraps ast.Comment
type SelectStmt struct {
	*ast.SelectStmt
	baseNode
}

//ForStmt wraps ast.Comment
type ForStmt struct {
	*ast.ForStmt
	baseNode
}

//RangeStmt wraps ast.Comment
type RangeStmt struct {
	*ast.RangeStmt
	baseNode
}

//ImportSpec wraps ast.Comment
type ImportSpec struct {
	*ast.ImportSpec
	baseNode
}

//ValueSpec wraps ast.Comment
type ValueSpec struct {
	*ast.ValueSpec
	baseNode
}

//TypeSpec wraps ast.Comment
type TypeSpec struct {
	*ast.TypeSpec
	baseNode
}

//BadDecl wraps ast.Comment
type BadDecl struct {
	*ast.BadDecl
	baseNode
}

//GenDecl wraps ast.Comment
type GenDecl struct {
	*ast.GenDecl
	baseNode
}

//FuncDecl wraps ast.Comment
type FuncDecl struct {
	*ast.FuncDecl
	baseNode
}

//File wraps ast.Comment
type File struct {
	*ast.File
	baseNode
}

//Package wraps ast.Comment
type Package struct {
	*ast.Package
	baseNode
}
