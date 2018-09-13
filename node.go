package astrav

import (
	"go/ast"
	"reflect"
)

//New creates a new node
func New(node ast.Node) Node {
	return creator(baseNode{node: node})
}

//Node wraps a ast.Node with helpful traversal functions
type Node interface {
	ast.Node

	Parent() Node
	Children() []Node
	Level() int
	AstNode() ast.Node
	Walk(f func(node Node) bool)
	Parents() []Node
	Contains(node Node) bool
	ChildByName(name string) Node
	SubNodesByType(nodeType NodeType) []Node
	IsType(nodeType NodeType) bool

	setRealMe(node Node)
}

type baseNode struct {
	node ast.Node

	realMe   Node
	parent   Node
	level    int
	built    bool
	children []Node
}

//Parent return the parent node
func (s *baseNode) Parent() Node {
	return s.parent
}

//Children returns all child nodes
func (s *baseNode) Children() []Node {
	s.build()
	return s.children
}

//Level returns the level counted from instantiated node = 0
func (s *baseNode) Level() int {
	return s.level
}

//AstNode returns the original ast.Node
func (s *baseNode) AstNode() ast.Node {
	return s.node
}

//Walk traverses the tree and its children.
// return false to skip children of the current element
func (s *baseNode) Walk(f func(node Node) bool) {
	if !f(s.realMe) {
		return
	}

	for _, child := range s.Children() {
		child.Walk(f)
	}
}

//Parents returns the parent path of nodes
func (s *baseNode) Parents() []Node {
	if s.parent == nil {
		return nil
	}
	return append([]Node{s.parent}, s.parent.Parents()...)
}

//Contains checks if a node contains another node
func (s *baseNode) Contains(node Node) bool {
	for _, p := range node.Parents() {
		if p == s.realMe {
			return true
		}
	}
	return false
}

//ChildByName retrieves a node among the direkt children by name (only nodes that have a name)
func (s baseNode) ChildByName(name string) Node {
	for _, child := range s.Children() {
		if f, ok := child.(Named); ok {
			ident := f.Ident()
			if ident != nil && ident.Name == name {
				return child
			}
		}
	}
	return nil
}

//SubNodesByType returns all sub nodes of a certain types. sub nodes are all nodes in the current node.
func (s baseNode) SubNodesByType(nodeType NodeType) []Node {
	var nodes []Node
	for _, child := range s.Children() {
		if child.IsType(nodeType) {
			nodes = append(nodes, child)
		}
		nodes = append(nodes, child.SubNodesByType(nodeType)...)
	}
	return nodes
}

//IsType checks if node is of given node type
func (s baseNode) IsType(nodeType NodeType) bool {
	return reflect.TypeOf(s.realMe).Name() == string(nodeType)
}

func (s *baseNode) setRealMe(node Node) {
	s.realMe = node
}

func (s *baseNode) build() {
	if s.built {
		return
	}
	s.built = true

	ast.Walk(s, s.node)
}

//Visit implements the ast.Visitor interface and is used to walk the underlying ast.Node tree
func (s *baseNode) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	if node == s.node {
		return s
	}

	n := creator(baseNode{
		node:   node,
		parent: s.realMe,
		level:  s.level + 1,
	})
	s.children = append(s.children, n)
	return nil
}
