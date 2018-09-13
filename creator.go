package astrav

import (
	"go/ast"
	"log"
)

func creator(bNode baseNode) Node {
	n := sw(bNode)
	n.setRealMe(n)
	return n
}

func sw(bNode baseNode) Node {
	if bNode.node == nil {
		log.Println("astrav creator: given node cannot be nil")
		return nil
	}

	switch n := bNode.node.(type) {
	case *ast.Comment:
		return &Comment{Comment: *n, baseNode: bNode}
	case *ast.CommentGroup:
		return &CommentGroup{CommentGroup: *n, baseNode: bNode}
	case *ast.Field:
		return &Field{Field: *n, baseNode: bNode}
	case *ast.FieldList:
		return &FieldList{FieldList: *n, baseNode: bNode}
	case *ast.BadExpr:
		return &BadExpr{BadExpr: *n, baseNode: bNode}
	case *ast.Ident:
		return &Ident{Ident: *n, baseNode: bNode}
	case *ast.Ellipsis:
		return &Ellipsis{Ellipsis: *n, baseNode: bNode}
	case *ast.BasicLit:
		return &BasicLit{BasicLit: *n, baseNode: bNode}
	case *ast.FuncLit:
		return &FuncLit{FuncLit: *n, baseNode: bNode}
	case *ast.CompositeLit:
		return &CompositeLit{CompositeLit: *n, baseNode: bNode}
	case *ast.ParenExpr:
		return &ParenExpr{ParenExpr: *n, baseNode: bNode}
	case *ast.SelectorExpr:
		return &SelectorExpr{SelectorExpr: *n, baseNode: bNode}
	case *ast.IndexExpr:
		return &IndexExpr{IndexExpr: *n, baseNode: bNode}
	case *ast.SliceExpr:
		return &SliceExpr{SliceExpr: *n, baseNode: bNode}
	case *ast.TypeAssertExpr:
		return &TypeAssertExpr{TypeAssertExpr: *n, baseNode: bNode}
	case *ast.CallExpr:
		return &CallExpr{CallExpr: *n, baseNode: bNode}
	case *ast.StarExpr:
		return &StarExpr{StarExpr: *n, baseNode: bNode}
	case *ast.UnaryExpr:
		return &UnaryExpr{UnaryExpr: *n, baseNode: bNode}
	case *ast.BinaryExpr:
		return &BinaryExpr{BinaryExpr: *n, baseNode: bNode}
	case *ast.KeyValueExpr:
		return &KeyValueExpr{KeyValueExpr: *n, baseNode: bNode}
	case *ast.ArrayType:
		return &ArrayType{ArrayType: *n, baseNode: bNode}
	case *ast.StructType:
		return &StructType{StructType: *n, baseNode: bNode}
	case *ast.FuncType:
		return &FuncType{FuncType: *n, baseNode: bNode}
	case *ast.InterfaceType:
		return &InterfaceType{InterfaceType: *n, baseNode: bNode}
	case *ast.MapType:
		return &MapType{MapType: *n, baseNode: bNode}
	case *ast.ChanType:
		return &ChanType{ChanType: *n, baseNode: bNode}
	case *ast.BadStmt:
		return &BadStmt{BadStmt: *n, baseNode: bNode}
	case *ast.DeclStmt:
		return &DeclStmt{DeclStmt: *n, baseNode: bNode}
	case *ast.EmptyStmt:
		return &EmptyStmt{EmptyStmt: *n, baseNode: bNode}
	case *ast.LabeledStmt:
		return &LabeledStmt{LabeledStmt: *n, baseNode: bNode}
	case *ast.ExprStmt:
		return &ExprStmt{ExprStmt: *n, baseNode: bNode}
	case *ast.SendStmt:
		return &SendStmt{SendStmt: *n, baseNode: bNode}
	case *ast.IncDecStmt:
		return &IncDecStmt{IncDecStmt: *n, baseNode: bNode}
	case *ast.AssignStmt:
		return &AssignStmt{AssignStmt: *n, baseNode: bNode}
	case *ast.GoStmt:
		return &GoStmt{GoStmt: *n, baseNode: bNode}
	case *ast.DeferStmt:
		return &DeferStmt{DeferStmt: *n, baseNode: bNode}
	case *ast.ReturnStmt:
		return &ReturnStmt{ReturnStmt: *n, baseNode: bNode}
	case *ast.BranchStmt:
		return &BranchStmt{BranchStmt: *n, baseNode: bNode}
	case *ast.BlockStmt:
		return &BlockStmt{BlockStmt: *n, baseNode: bNode}
	case *ast.IfStmt:
		return &IfStmt{IfStmt: *n, baseNode: bNode}
	case *ast.CaseClause:
		return &CaseClause{CaseClause: *n, baseNode: bNode}
	case *ast.SwitchStmt:
		return &SwitchStmt{SwitchStmt: *n, baseNode: bNode}
	case *ast.TypeSwitchStmt:
		return &TypeSwitchStmt{TypeSwitchStmt: *n, baseNode: bNode}
	case *ast.CommClause:
		return &CommClause{CommClause: *n, baseNode: bNode}
	case *ast.SelectStmt:
		return &SelectStmt{SelectStmt: *n, baseNode: bNode}
	case *ast.ForStmt:
		return &ForStmt{ForStmt: *n, baseNode: bNode}
	case *ast.RangeStmt:
		return &RangeStmt{RangeStmt: *n, baseNode: bNode}
	case *ast.ImportSpec:
		return &ImportSpec{ImportSpec: *n, baseNode: bNode}
	case *ast.ValueSpec:
		return &ValueSpec{ValueSpec: *n, baseNode: bNode}
	case *ast.TypeSpec:
		return &TypeSpec{TypeSpec: *n, baseNode: bNode}
	case *ast.BadDecl:
		return &BadDecl{BadDecl: *n, baseNode: bNode}
	case *ast.GenDecl:
		return &GenDecl{GenDecl: *n, baseNode: bNode}
	case *ast.FuncDecl:
		return &FuncDecl{FuncDecl: *n, baseNode: bNode}
	case *ast.File:
		return &File{File: *n, baseNode: bNode}
	case *ast.Package:
		return &Package{Package: *n, baseNode: bNode}
	default:
		log.Printf("astrav: not implemented ast.Node type found: %T\n", n)
	}
	return nil
}
