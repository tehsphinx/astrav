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
	if bNode.Node == nil {
		log.Println("astrav creator: given node cannot be nil")
		return nil
	}

	switch n := bNode.Node.(type) {
	case *ast.Comment:
		return &Comment{baseNode: bNode}
	case *ast.CommentGroup:
		return &CommentGroup{baseNode: bNode}
	case *ast.Field:
		return &Field{baseNode: bNode}
	case *ast.FieldList:
		return &FieldList{baseNode: bNode}
	case *ast.BadExpr:
		return &BadExpr{baseNode: bNode}
	case *ast.Ident:
		return &Ident{baseNode: bNode}
	case *ast.Ellipsis:
		return &Ellipsis{baseNode: bNode}
	case *ast.BasicLit:
		return &BasicLit{baseNode: bNode}
	case *ast.FuncLit:
		return &FuncLit{baseNode: bNode}
	case *ast.CompositeLit:
		return &CompositeLit{baseNode: bNode}
	case *ast.ParenExpr:
		return &ParenExpr{baseNode: bNode}
	case *ast.SelectorExpr:
		return &SelectorExpr{baseNode: bNode}
	case *ast.IndexExpr:
		return &IndexExpr{baseNode: bNode}
	case *ast.SliceExpr:
		return &SliceExpr{baseNode: bNode}
	case *ast.TypeAssertExpr:
		return &TypeAssertExpr{baseNode: bNode}
	case *ast.CallExpr:
		return &CallExpr{baseNode: bNode}
	case *ast.StarExpr:
		return &StarExpr{baseNode: bNode}
	case *ast.UnaryExpr:
		return &UnaryExpr{baseNode: bNode}
	case *ast.BinaryExpr:
		return &BinaryExpr{baseNode: bNode}
	case *ast.KeyValueExpr:
		return &KeyValueExpr{baseNode: bNode}
	case *ast.ArrayType:
		return &ArrayType{baseNode: bNode}
	case *ast.StructType:
		return &StructType{baseNode: bNode}
	case *ast.FuncType:
		return &FuncType{baseNode: bNode}
	case *ast.InterfaceType:
		return &InterfaceType{baseNode: bNode}
	case *ast.MapType:
		return &MapType{baseNode: bNode}
	case *ast.ChanType:
		return &ChanType{baseNode: bNode}
	case *ast.BadStmt:
		return &BadStmt{baseNode: bNode}
	case *ast.DeclStmt:
		return &DeclStmt{baseNode: bNode}
	case *ast.EmptyStmt:
		return &EmptyStmt{baseNode: bNode}
	case *ast.LabeledStmt:
		return &LabeledStmt{baseNode: bNode}
	case *ast.ExprStmt:
		return &ExprStmt{baseNode: bNode}
	case *ast.SendStmt:
		return &SendStmt{baseNode: bNode}
	case *ast.IncDecStmt:
		return &IncDecStmt{baseNode: bNode}
	case *ast.AssignStmt:
		return &AssignStmt{baseNode: bNode}
	case *ast.GoStmt:
		return &GoStmt{baseNode: bNode}
	case *ast.DeferStmt:
		return &DeferStmt{baseNode: bNode}
	case *ast.ReturnStmt:
		return &ReturnStmt{baseNode: bNode}
	case *ast.BranchStmt:
		return &BranchStmt{baseNode: bNode}
	case *ast.BlockStmt:
		return &BlockStmt{baseNode: bNode}
	case *ast.IfStmt:
		return &IfStmt{baseNode: bNode}
	case *ast.CaseClause:
		return &CaseClause{baseNode: bNode}
	case *ast.SwitchStmt:
		return &SwitchStmt{baseNode: bNode}
	case *ast.TypeSwitchStmt:
		return &TypeSwitchStmt{baseNode: bNode}
	case *ast.CommClause:
		return &CommClause{baseNode: bNode}
	case *ast.SelectStmt:
		return &SelectStmt{baseNode: bNode}
	case *ast.ForStmt:
		return &ForStmt{baseNode: bNode}
	case *ast.RangeStmt:
		return &RangeStmt{baseNode: bNode}
	case *ast.ImportSpec:
		return &ImportSpec{baseNode: bNode}
	case *ast.ValueSpec:
		return &ValueSpec{baseNode: bNode}
	case *ast.TypeSpec:
		return &TypeSpec{baseNode: bNode}
	case *ast.BadDecl:
		return &BadDecl{baseNode: bNode}
	case *ast.GenDecl:
		return &GenDecl{baseNode: bNode}
	case *ast.FuncDecl:
		return &FuncDecl{baseNode: bNode}
	case *ast.File:
		return &File{baseNode: bNode}
	case *ast.Package:
		return &Package{baseNode: bNode}
	default:
		log.Printf("astrav: not implemented ast.Node type found: %T\n", n)
	}
	return nil
}
