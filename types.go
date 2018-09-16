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
	NodeTypeComment        = "*astrav.Comment"
	NodeTypeCommentGroup   = "*astrav.CommentGroup"
	NodeTypeField          = "*astrav.Field"
	NodeTypeFieldList      = "*astrav.FieldList"
	NodeTypeBadExpr        = "*astrav.BadExpr"
	NodeTypeIdent          = "*astrav.Ident"
	NodeTypeEllipsis       = "*astrav.Ellipsis"
	NodeTypeBasicLit       = "*astrav.BasicLit"
	NodeTypeFuncLit        = "*astrav.FuncLit"
	NodeTypeCompositeLit   = "*astrav.CompositeLit"
	NodeTypeParenExpr      = "*astrav.ParenExpr"
	NodeTypeSelectorExpr   = "*astrav.SelectorExpr"
	NodeTypeIndexExpr      = "*astrav.IndexExpr"
	NodeTypeSliceExpr      = "*astrav.SliceExpr"
	NodeTypeTypeAssertExpr = "*astrav.TypeAssertExpr"
	NodeTypeCallExpr       = "*astrav.CallExpr"
	NodeTypeStarExpr       = "*astrav.StarExpr"
	NodeTypeUnaryExpr      = "*astrav.UnaryExpr"
	NodeTypeBinaryExpr     = "*astrav.BinaryExpr"
	NodeTypeKeyValueExpr   = "*astrav.KeyValueExpr"
	NodeTypeArrayType      = "*astrav.ArrayType"
	NodeTypeStructType     = "*astrav.StructType"
	NodeTypeFuncType       = "*astrav.FuncType"
	NodeTypeInterfaceType  = "*astrav.InterfaceType"
	NodeTypeMapType        = "*astrav.MapType"
	NodeTypeChanType       = "*astrav.ChanType"
	NodeTypeBadStmt        = "*astrav.BadStmt"
	NodeTypeDeclStmt       = "*astrav.DeclStmt"
	NodeTypeEmptyStmt      = "*astrav.EmptyStmt"
	NodeTypeLabeledStmt    = "*astrav.LabeledStmt"
	NodeTypeExprStmt       = "*astrav.ExprStmt"
	NodeTypeSendStmt       = "*astrav.SendStmt"
	NodeTypeIncDecStmt     = "*astrav.IncDecStmt"
	NodeTypeAssignStmt     = "*astrav.AssignStmt"
	NodeTypeGoStmt         = "*astrav.GoStmt"
	NodeTypeDeferStmt      = "*astrav.DeferStmt"
	NodeTypeReturnStmt     = "*astrav.ReturnStmt"
	NodeTypeBranchStmt     = "*astrav.BranchStmt"
	NodeTypeBlockStmt      = "*astrav.BlockStmt"
	NodeTypeIfStmt         = "*astrav.IfStmt"
	NodeTypeCaseClause     = "*astrav.CaseClause"
	NodeTypeSwitchStmt     = "*astrav.SwitchStmt"
	NodeTypeTypeSwitchStmt = "*astrav.TypeSwitchStmt"
	NodeTypeCommClause     = "*astrav.CommClause"
	NodeTypeSelectStmt     = "*astrav.SelectStmt"
	NodeTypeForStmt        = "*astrav.ForStmt"
	NodeTypeRangeStmt      = "*astrav.RangeStmt"
	NodeTypeImportSpec     = "*astrav.ImportSpec"
	NodeTypeValueSpec      = "*astrav.ValueSpec"
	NodeTypeTypeSpec       = "*astrav.TypeSpec"
	NodeTypeBadDecl        = "*astrav.BadDecl"
	NodeTypeGenDecl        = "*astrav.GenDecl"
	NodeTypeFuncDecl       = "*astrav.FuncDecl"
	NodeTypeFile           = "*astrav.File"
	NodeTypePackage        = "*astrav.Package"
)
