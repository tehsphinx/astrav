package astrav

import (
	"go/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInfo(t *testing.T) {
	pkg := getPackageFromPath(t, "../go-analyzer/tests/two-fer/17")

	var nameType types.Type
	for expr, def := range pkg.Info().Defs {
		if expr.Name == "name" {
			nameType = def.Type()
		}
	}
	assert.NotNil(t, nameType)

	pkg.Walk(func(n Node) bool {
		if !n.IsNodeType(NodeTypeIdent) {
			return true
		}
		node := n.(*Ident)
		if node.NodeName() != "name" {
			return true
		}
		assert.Equal(t, nameType, node.Object().Type())
		return true
	})
}
