package astrav

import (
	"go/types"
	"testing"

	"github.com/tehsphinx/dbg"

	"github.com/stretchr/testify/assert"
)

func TestParseInfo(t *testing.T) {
	pkg := getPackageFromPath(t, "../go-analyzer/tests/two-fer/17")

	var nameType types.Type
	for expr, def := range pkg.Info().Defs {
		if expr.Name == "name" {
			nameType = def.Type()
		}
		dbg.Blue(expr, expr.NamePos)
	}
	assert.NotNil(t, nameType)

	pkg.Walk(func(n Node) bool {
		if !n.IsNodeType(NodeTypeIdent) {
			return true
		}
		node := n.(*Ident)
		if node.NodeName().String() != "name" {
			return true
		}
		dbg.Magenta(node, node.NamePos, node.Pos(), node.Object())
		assert.Equal(t, nameType, node.Object().Type())
		return true
	})
}
