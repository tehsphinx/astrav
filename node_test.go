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
	f := getFile(t, 1)
	n := NewNode(f)

	file, ok := n.(*File)

	assert.True(t, ok)
	assert.Equal(t, f, file.node)
}

func TestBaseNode_AstNode(t *testing.T) {
	f := getFile(t, 1)
	n := NewNode(f)

	assert.Equal(t, f, n.AstNode())
}

func TestBaseNode_Children(t *testing.T) {
	n := getTree(t, 1)

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
	n := getTree(t, 1)

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
	n := getTree(t, 1)

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
	n := getTree(t, 1)

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
	n := getTree(t, 1)

	child1 := n.Children()[2]
	child2 := child1.Children()[2]
	child3 := child2.Children()[1]
	child4 := child3.Children()[2]
	assert.Equal(t, []Node{child3, child2, child1, n}, child4.Parents())
}

func TestBaseNode_Walk(t *testing.T) {
	n := getTree(t, 1)

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
	n := getTree(t, 1)

	f := n.ChildByName("Score")
	assert.Equal(t, "Score", f.(*FuncDecl).Name.Name)
}

func TestBaseNode_ChildByNodeType(t *testing.T) {
	n := getTree(t, 3)

	f := n.ChildByNodeType(NodeTypeFuncDecl)
	assert.Equal(t, "IsIsogram", f.(*FuncDecl).Name.Name)
}

func TestBaseNode_ChildrenByNodeType(t *testing.T) {
	n := getTree(t, 3)

	f := n.FindFirstByNodeType(NodeTypeRangeStmt)
	nodes := f.ChildrenByNodeType(NodeTypeIdent)
	assert.Equal(t, 3, len(nodes))
}

func TestBaseNode_FindFirstByName(t *testing.T) {
	n := getTree(t, 1)

	f := n.FindFirstByName("Score")
	assert.Equal(t, "Score", f.(*FuncDecl).Name.Name)

	f = n.FindFirstByName("ToLower")
	assert.Equal(t, "ToLower", f.(*SelectorExpr).Sel.Name)
}

func TestBaseNode_IsType(t *testing.T) {
	n := getTree(t, 1)

	assert.True(t, n.IsNodeType(NodeType(reflect.TypeOf(n).String())))

	for _, child := range n.Children() {
		tName := reflect.TypeOf(child).String()
		assert.False(t, n.IsNodeType(NodeType(tName)))
		assert.True(t, child.IsNodeType(NodeType(tName)))
	}
}

func TestBaseNode_FindByType(t *testing.T) {
	n := getTree(t, 1)

	sn := n.FindByNodeType(NodeTypeSwitchStmt)
	assert.Equal(t, 1, len(sn))
	if _, ok := sn[0].(*SwitchStmt); !ok {
		t.Fail()
	}
}

func TestBaseNode_FindByName(t *testing.T) {
	n := getTree(t, 1)

	fns := n.FindByName("Score")
	assert.Equal(t, 2, len(fns))
	assert.Equal(t, "Score", fns[0].(*FuncDecl).Name.Name)

	fns = n.FindByName("ToLower")
	assert.Equal(t, 2, len(fns))
	assert.Equal(t, "ToLower", fns[0].(*SelectorExpr).Sel.Name)

	fns = n.FindFirstByName("Score").FindByName("c")
	assert.Equal(t, 2, len(fns))
	assert.Equal(t, "c", fns[0].(*Ident).Name)
}

func TestBaseNode_FindIdentByName(t *testing.T) {
	n := getTree(t, 1)

	idents := n.FindIdentByName("Score")
	assert.Equal(t, 1, len(idents))
	assert.Equal(t, "Score", idents[0].Name)

	idents = n.FindIdentByName("ToLower")
	assert.Equal(t, 1, len(idents))
	assert.Equal(t, "ToLower", idents[0].Name)

	idents = n.FindFirstByName("Score").FindIdentByName("c")
	assert.Equal(t, 2, len(idents))
	assert.Equal(t, "c", idents[0].Name)
}

func TestBaseNode_FindFirstIdentByName(t *testing.T) {
	n := getTree(t, 1)

	ident := n.FindFirstIdentByName("Score")
	assert.Equal(t, "Score", ident.Name)

	ident = n.FindFirstIdentByName("ToLower")
	assert.Equal(t, "ToLower", ident.Name)

	ident = n.FindFirstByName("Score").FindFirstIdentByName("c")
	assert.Equal(t, "c", ident.Name)
}

func TestBaseNode_FindByValueType(t *testing.T) {
	n := getTree(t, 1)

	nodes := n.FindByValueType("byte")
	assert.Equal(t, 4, len(nodes))
	for _, v := range nodes {
		assert.Equal(t, "byte", v.ValueType().String())
	}
}

func TestBaseNode_FindByToken(t *testing.T) {
	n := getTree(t, 3)

	nodes := n.FindByToken(token.BREAK)
	assert.Equal(t, 1, len(nodes))
	for _, v := range nodes {
		assert.Equal(t, token.BREAK, v.(Token).Token())
	}
}

func TestBaseNode_ValueType(t *testing.T) {
	n := getTree(t, 1)

	v := n.FindFirstIdentByName("Score")
	assert.Equal(t, "func(word string) int", v.ValueType().String())

	v = n.FindFirstIdentByName("c")
	assert.Equal(t, "byte", v.ValueType().String())
}

func TestBaseNode_IsValueType(t *testing.T) {
	n := getTree(t, 1)

	var isTp bool
	for _, v := range n.FindFirstByName("Score").FindByName("c") {
		if v.IsValueType("byte") {
			isTp = true
		}
	}
	assert.True(t, true, isTp)
}

func TestBaseNode_IsNodeType(t *testing.T) {
	n := getTree(t, 1)

	assert.True(t, n.IsNodeType(NodeTypeFile))
}

func TestBaseNode_FindByNodeType(t *testing.T) {
	n := getTree(t, 1)

	nodes := n.FindByNodeType(NodeTypeFuncDecl)
	for _, node := range nodes {
		assert.True(t, node.IsNodeType(NodeTypeFuncDecl))
	}
}

func TestBaseNode_IsContainedByType(t *testing.T) {
	n := getTree(t, 1)

	child := n.FindFirstByName("c")

	for _, node := range child.Parents() {
		assert.True(t, child.IsContainedByType(NodeType(reflect.TypeOf(node).String())))
	}
}

func getTree(t *testing.T, example int) Node {
	f := getFile(t, example)
	return NewNode(f)
}

func getFile(t *testing.T, example int) ast.Node {
	fileName := fmt.Sprintf("example/%d/example.go", example)

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, fileName, nil, parser.AllErrors)
	if err != nil {
		t.Fatal(err)
	}
	ParseInfo(".", fs, []*ast.File{f})
	return f
}

func getPackage(t *testing.T, example int) Node {
	path := fmt.Sprintf("example/%d", example)
	folder := NewFolder(path)
	pkgs, err := folder.ParseFolder()
	if err != nil {
		t.Fatal(err)
	}
	for _, pkg := range pkgs {
		return pkg
	}
	return nil
}
