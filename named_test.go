package astrav

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncDecl_NodeName(t *testing.T) {
	n := getTree(t)

	var found bool
	n.Walk(func(node Node) bool {
		if !node.IsNodeType(NodeTypeFuncDecl) {
			return true
		}

		named, ok := node.(Named)
		assert.Equal(t, true, ok)

		assert.NotNil(t, named.NodeName())
		found = true
		return true
	})

	assert.True(t, found)
}

func TestFile_NodeName(t *testing.T) {
	n := getTree(t)

	var found bool
	n.Walk(func(node Node) bool {
		if !node.IsNodeType(NodeTypeFile) {
			return true
		}

		named, ok := node.(Named)
		assert.Equal(t, true, ok)

		assert.NotNil(t, named.NodeName())
		found = true
		return true
	})

	assert.True(t, found)
}

func TestIdent_NodeName(t *testing.T) {
	n := getTree(t)

	var found bool
	n.Walk(func(node Node) bool {
		if !node.IsNodeType(NodeTypeIdent) {
			return true
		}

		named, ok := node.(Named)
		assert.Equal(t, true, ok)

		assert.NotNil(t, named.NodeName())
		found = true
		return true
	})

	assert.True(t, found)
}

func TestImportSpec_NodeName(t *testing.T) {
	n := getTree(t)

	var found bool
	n.Walk(func(node Node) bool {
		if !node.IsNodeType(NodeTypeImportSpec) {
			return true
		}

		named, ok := node.(Named)
		assert.Equal(t, true, ok)

		assert.Nil(t, named.NodeName())
		found = true
		return true
	})

	assert.True(t, found)
}

func TestSelectorExpr_NodeName(t *testing.T) {
	n := getTree(t)

	var found bool
	n.Walk(func(node Node) bool {
		if !node.IsNodeType(NodeTypeSelectorExpr) {
			return true
		}

		named, ok := node.(Named)
		assert.Equal(t, true, ok)

		assert.NotNil(t, named.NodeName())
		found = true
		return true
	})

	assert.True(t, found)
}
