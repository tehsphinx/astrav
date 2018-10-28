package astrav

import (
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
)

var (
	info types.Info
)

//ParseInfo parses all files for type information which is then available
// from the Nodes. When using Folder.ParseFolder, this is done automatically.
func ParseInfo(path string, fSet *token.FileSet, files []*ast.File) (*types.Package, error) {
	info = types.Info{
		Types:      map[ast.Expr]types.TypeAndValue{},
		Defs:       map[*ast.Ident]types.Object{},
		Uses:       map[*ast.Ident]types.Object{},
		Scopes:     map[ast.Node]*types.Scope{},
		Implicits:  map[ast.Node]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
	var conf = types.Config{
		Importer: importer.Default(),
	}

	pkg, err := conf.Check(path, fSet, files, &info)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}

//NodeType defines a node type string to search for type
type NodeType string

// Nodetype contants
const (
	NodeTypeComment        NodeType = "*astrav.Comment"
	NodeTypeCommentGroup   NodeType = "*astrav.CommentGroup"
	NodeTypeField          NodeType = "*astrav.Field"
	NodeTypeFieldList      NodeType = "*astrav.FieldList"
	NodeTypeBadExpr        NodeType = "*astrav.BadExpr"
	NodeTypeIdent          NodeType = "*astrav.Ident"
	NodeTypeEllipsis       NodeType = "*astrav.Ellipsis"
	NodeTypeBasicLit       NodeType = "*astrav.BasicLit"
	NodeTypeFuncLit        NodeType = "*astrav.FuncLit"
	NodeTypeCompositeLit   NodeType = "*astrav.CompositeLit"
	NodeTypeParenExpr      NodeType = "*astrav.ParenExpr"
	NodeTypeSelectorExpr   NodeType = "*astrav.SelectorExpr"
	NodeTypeIndexExpr      NodeType = "*astrav.IndexExpr"
	NodeTypeSliceExpr      NodeType = "*astrav.SliceExpr"
	NodeTypeTypeAssertExpr NodeType = "*astrav.TypeAssertExpr"
	NodeTypeCallExpr       NodeType = "*astrav.CallExpr"
	NodeTypeStarExpr       NodeType = "*astrav.StarExpr"
	NodeTypeUnaryExpr      NodeType = "*astrav.UnaryExpr"
	NodeTypeBinaryExpr     NodeType = "*astrav.BinaryExpr"
	NodeTypeKeyValueExpr   NodeType = "*astrav.KeyValueExpr"
	NodeTypeArrayType      NodeType = "*astrav.ArrayType"
	NodeTypeStructType     NodeType = "*astrav.StructType"
	NodeTypeFuncType       NodeType = "*astrav.FuncType"
	NodeTypeInterfaceType  NodeType = "*astrav.InterfaceType"
	NodeTypeMapType        NodeType = "*astrav.MapType"
	NodeTypeChanType       NodeType = "*astrav.ChanType"
	NodeTypeBadStmt        NodeType = "*astrav.BadStmt"
	NodeTypeDeclStmt       NodeType = "*astrav.DeclStmt"
	NodeTypeEmptyStmt      NodeType = "*astrav.EmptyStmt"
	NodeTypeLabeledStmt    NodeType = "*astrav.LabeledStmt"
	NodeTypeExprStmt       NodeType = "*astrav.ExprStmt"
	NodeTypeSendStmt       NodeType = "*astrav.SendStmt"
	NodeTypeIncDecStmt     NodeType = "*astrav.IncDecStmt"
	NodeTypeAssignStmt     NodeType = "*astrav.AssignStmt"
	NodeTypeGoStmt         NodeType = "*astrav.GoStmt"
	NodeTypeDeferStmt      NodeType = "*astrav.DeferStmt"
	NodeTypeReturnStmt     NodeType = "*astrav.ReturnStmt"
	NodeTypeBranchStmt     NodeType = "*astrav.BranchStmt"
	NodeTypeBlockStmt      NodeType = "*astrav.BlockStmt"
	NodeTypeIfStmt         NodeType = "*astrav.IfStmt"
	NodeTypeCaseClause     NodeType = "*astrav.CaseClause"
	NodeTypeSwitchStmt     NodeType = "*astrav.SwitchStmt"
	NodeTypeTypeSwitchStmt NodeType = "*astrav.TypeSwitchStmt"
	NodeTypeCommClause     NodeType = "*astrav.CommClause"
	NodeTypeSelectStmt     NodeType = "*astrav.SelectStmt"
	NodeTypeForStmt        NodeType = "*astrav.ForStmt"
	NodeTypeRangeStmt      NodeType = "*astrav.RangeStmt"
	NodeTypeImportSpec     NodeType = "*astrav.ImportSpec"
	NodeTypeValueSpec      NodeType = "*astrav.ValueSpec"
	NodeTypeTypeSpec       NodeType = "*astrav.TypeSpec"
	NodeTypeBadDecl        NodeType = "*astrav.BadDecl"
	NodeTypeGenDecl        NodeType = "*astrav.GenDecl"
	NodeTypeFuncDecl       NodeType = "*astrav.FuncDecl"
	NodeTypeFile           NodeType = "*astrav.File"
	NodeTypePackage        NodeType = "*astrav.Package"
)
