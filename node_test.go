package astrav

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	f := getFile(t)
	n := New(f)

	file, ok := n.(*File)

	assert.True(t, ok)
	assert.Equal(t, f, file.node)
}

func TestBaseNode_AstNode(t *testing.T) {
	f := getFile(t)
	n := New(f)

	assert.Equal(t, f, n.AstNode())
}

func TestBaseNode_Children(t *testing.T) {
	n := getTree(t)

	for i, child := range n.Children() {
		switch i {
		case 0:
			assert.Equal(t, "*astrav.Ident", reflect.TypeOf(child).String())
		case 1:
			assert.Equal(t, "*astrav.GenDecl", reflect.TypeOf(child).String())
		case 2, 3:
			assert.Equal(t, "*astrav.FuncDecl", reflect.TypeOf(child).String())
		}
	}
}

func TestBaseNode_Contains(t *testing.T) {
	n := getTree(t)

	child := n.Children()[2].Children()[2].Children()[1].Children()[2]
	assert.Equal(t, true, n.Contains(child))
	assert.Equal(t, false, child.Contains(n))

	child2 := child.Children()[0].Children()[0]
	assert.Equal(t, true, child.Contains(child2))
	assert.Equal(t, true, n.Contains(child2))
	assert.Equal(t, false, child2.Contains(n))
	assert.Equal(t, false, child2.Contains(child))
}

func TestBaseNode_Level(t *testing.T) {
	n := getTree(t)

	assert.Equal(t, 0, n.Level())
	child1 := n.Children()[2]
	assert.Equal(t, 1, child1.Level())
	child2 := child1.Children()[2]
	assert.Equal(t, 2, child2.Level())
	child3 := child2.Children()[1]
	assert.Equal(t, 3, child3.Level())
	child4 := child3.Children()[2]
	assert.Equal(t, 4, child4.Level())
}

func TestBaseNode_Parent(t *testing.T) {
	n := getTree(t)

	child1 := n.Children()[2]
	assert.Equal(t, n, child1.Parent())
	child2 := child1.Children()[2]
	assert.Equal(t, child1, child2.Parent())
	child3 := child2.Children()[1]
	assert.Equal(t, child2, child3.Parent())
	child4 := child3.Children()[2]
	assert.Equal(t, child3, child4.Parent())
}

func TestBaseNode_Parents(t *testing.T) {
	n := getTree(t)

	child1 := n.Children()[2]
	child2 := child1.Children()[2]
	child3 := child2.Children()[1]
	child4 := child3.Children()[2]
	assert.Equal(t, []Node{child3, child2, child1, n}, child4.Parents())
}

func TestBaseNode_Walk(t *testing.T) {
	n := getTree(t)

	var (
		ints    = []int{0, 2, 2, 1}
		counter = make([]int, len(ints))
	)

	child1 := n.Children()[ints[1]]
	child2 := child1.Children()[ints[2]]
	child3 := child2.Children()[ints[3]]
	nodes := []Node{n, child1, child2, child3}

	n.Walk(func(node Node) bool {
		level := node.Level()
		if len(ints) <= level {
			return false
		}
		corrIndex := counter[level] == ints[level]
		counter[level]++
		if corrIndex {
			assert.Equal(t, nodes[level], node, fmt.Sprintf("failed at level %d", level))
		}
		return corrIndex
	})
}

func TestBaseNode_ChildByName(t *testing.T) {
	n := getTree(t)

	f := n.ChildByName("Score")
	assert.Equal(t, "Score", f.(*FuncDecl).Name.Name)
}

func TestBaseNode_SubNodesByType(t *testing.T) {
	t.Fail()
}

func TestBaseNode_IsType(t *testing.T) {
	t.Fail()
}

func getTree(t *testing.T) Node {
	f := getFile(t)
	return New(f)
}

func getFile(t *testing.T) ast.Node {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "example/file_to_parse.go", nil, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	return f
}
