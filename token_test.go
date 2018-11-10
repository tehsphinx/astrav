package astrav

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBranchStmt_Token(t *testing.T) {
	n := getTree(t, 3)

	var found bool
	n.Walk(func(node Node) bool {
		if !node.IsNodeType(NodeTypeBranchStmt) {
			return true
		}

		tok, ok := node.(Token)
		assert.Equal(t, true, ok)

		assert.NotNil(t, tok.Token())
		found = true
		return true
	})

	assert.True(t, found)
}
