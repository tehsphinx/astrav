package astrav

import "go/ast"

//Comment wraps ast.Comment
type Comment struct {
	*ast.Comment
	baseNode
}

//CommentGroup wraps ast.CommentGroup
type CommentGroup struct {
	*ast.CommentGroup
	baseNode
}

//Field wraps ast.Field
type Field struct {
	*ast.Field
	baseNode
}

//FieldList wraps ast.FieldList
type FieldList struct {
	*ast.FieldList
	baseNode
}

//BadExpr wraps ast.BadExpr
type BadExpr struct {
	*ast.BadExpr
	baseNode
}

//Ident wraps ast.Ident
type Ident struct {
	*ast.Ident
	baseNode
}

//Ellipsis wraps ast.Ellipsis
type Ellipsis struct {
	*ast.Ellipsis
	baseNode
}

//BasicLit wraps ast.BasicLit
type BasicLit struct {
	*ast.BasicLit
	baseNode
}

//FuncLit wraps ast.FuncLit
type FuncLit struct {
	*ast.FuncLit
	baseNode
}

//CompositeLit wraps ast.CompositeLit
type CompositeLit struct {
	*ast.CompositeLit
	baseNode
}

//ParenExpr wraps ast.ParenExpr
type ParenExpr struct {
	*ast.ParenExpr
	baseNode
}

//SelectorExpr wraps ast.SelectorExpr
type SelectorExpr struct {
	*ast.SelectorExpr
	baseNode
}

//IndexExpr wraps ast.IndexExpr
type IndexExpr struct {
	*ast.IndexExpr
	baseNode
}

//SliceExpr wraps ast.SliceExpr
type SliceExpr struct {
	*ast.SliceExpr
	baseNode
}

//TypeAssertExpr wraps ast.TypeAssertExpr
type TypeAssertExpr struct {
	*ast.TypeAssertExpr
	baseNode
}

//CallExpr wraps ast.CallExpr
type CallExpr struct {
	*ast.CallExpr
	baseNode
}

//StarExpr wraps ast.StarExpr
type StarExpr struct {
	*ast.StarExpr
	baseNode
}

//UnaryExpr wraps ast.UnaryExpr
type UnaryExpr struct {
	*ast.UnaryExpr
	baseNode
}

//BinaryExpr wraps ast.BinaryExpr
type BinaryExpr struct {
	*ast.BinaryExpr
	baseNode
}

//KeyValueExpr wraps ast.KeyValueExpr
type KeyValueExpr struct {
	*ast.KeyValueExpr
	baseNode
}

//ArrayType wraps ast.ArrayType
type ArrayType struct {
	*ast.ArrayType
	baseNode
}

//StructType wraps ast.StructType
type StructType struct {
	*ast.StructType
	baseNode
}

//FuncType wraps ast.FuncType
type FuncType struct {
	*ast.FuncType
	baseNode
}

//InterfaceType wraps ast.InterfaceType
type InterfaceType struct {
	*ast.InterfaceType
	baseNode
}

//MapType wraps ast.MapType
type MapType struct {
	*ast.MapType
	baseNode
}

//ChanType wraps ast.ChanType
type ChanType struct {
	*ast.ChanType
	baseNode
}

//BadStmt wraps ast.BadStmt
type BadStmt struct {
	*ast.BadStmt
	baseNode
}

//DeclStmt wraps ast.DeclStmt
type DeclStmt struct {
	*ast.DeclStmt
	baseNode
}

//EmptyStmt wraps ast.EmptyStmt
type EmptyStmt struct {
	*ast.EmptyStmt
	baseNode
}

//LabeledStmt wraps ast.LabeledStmt
type LabeledStmt struct {
	*ast.LabeledStmt
	baseNode
}

//ExprStmt wraps ast.ExprStmt
type ExprStmt struct {
	*ast.ExprStmt
	baseNode
}

//SendStmt wraps ast.SendStmt
type SendStmt struct {
	*ast.SendStmt
	baseNode
}

//IncDecStmt wraps ast.IncDecStmt
type IncDecStmt struct {
	*ast.IncDecStmt
	baseNode
}

//AssignStmt wraps ast.AssignStmt
type AssignStmt struct {
	*ast.AssignStmt
	baseNode
}

//GoStmt wraps ast.GoStmt
type GoStmt struct {
	*ast.GoStmt
	baseNode
}

//DeferStmt wraps ast.DeferStmt
type DeferStmt struct {
	*ast.DeferStmt
	baseNode
}

//ReturnStmt wraps ast.ReturnStmt
type ReturnStmt struct {
	*ast.ReturnStmt
	baseNode
}

//BranchStmt wraps ast.BranchStmt
type BranchStmt struct {
	*ast.BranchStmt
	baseNode
}

//BlockStmt wraps ast.BlockStmt
type BlockStmt struct {
	*ast.BlockStmt
	baseNode
}

//IfStmt wraps ast.IfStmt
type IfStmt struct {
	*ast.IfStmt
	baseNode
}

//CaseClause wraps ast.CaseClause
type CaseClause struct {
	*ast.CaseClause
	baseNode
}

//SwitchStmt wraps ast.SwitchStmt
type SwitchStmt struct {
	*ast.SwitchStmt
	baseNode
}

//TypeSwitchStmt wraps ast.TypeSwitchStmt
type TypeSwitchStmt struct {
	*ast.TypeSwitchStmt
	baseNode
}

//CommClause wraps ast.CommClause
type CommClause struct {
	*ast.CommClause
	baseNode
}

//SelectStmt wraps ast.SelectStmt
type SelectStmt struct {
	*ast.SelectStmt
	baseNode
}

//ForStmt wraps ast.ForStmt
type ForStmt struct {
	*ast.ForStmt
	baseNode
}

//RangeStmt wraps ast.RangeStmt
type RangeStmt struct {
	*ast.RangeStmt
	baseNode
}

//ImportSpec wraps ast.ImportSpec
type ImportSpec struct {
	*ast.ImportSpec
	baseNode
}

//ValueSpec wraps ast.ValueSpec
type ValueSpec struct {
	*ast.ValueSpec
	baseNode
}

//TypeSpec wraps ast.TypeSpec
type TypeSpec struct {
	*ast.TypeSpec
	baseNode
}

//BadDecl wraps ast.BadDecl
type BadDecl struct {
	*ast.BadDecl
	baseNode
}

//GenDecl wraps ast.GenDecl
type GenDecl struct {
	*ast.GenDecl
	baseNode
}

//FuncDecl wraps ast.FuncDecl
type FuncDecl struct {
	*ast.FuncDecl
	baseNode
}

//File wraps ast.File
type File struct {
	*ast.File
	baseNode
}
