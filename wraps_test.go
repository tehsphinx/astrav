package astrav

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectorExpr_PackageName(t *testing.T) {
	n := getTree(t)

	selExpr := n.FindFirstByNodeType(NodeTypeSelectorExpr).(*SelectorExpr)
	pkgIdent := selExpr.PackageIdent()

	assert.Equal(t, "strings", pkgIdent.Name)
	assert.Equal(t, selExpr, pkgIdent.Parent())
}

func TestFuncType_Params(t *testing.T) {
	n := getTree(t)

	f := n.FindFirstByName("Score").ChildByNodeType(NodeTypeFuncType)
	params := f.(*FuncType).Params()
	assert.NotNil(t, params)
	assert.Equal(t, 1, len(params.List))
}

func TestFuncType_Results(t *testing.T) {
	n := getTree(t)

	f := n.FindFirstByName("Score").ChildByNodeType(NodeTypeFuncType)
	params := f.(*FuncType).Results()
	assert.NotNil(t, params)
	assert.Equal(t, 1, len(params.List))
}
