package astrav

import (
	"bytes"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"regexp"
	"strings"
)

//NewNode creates a new node
func NewNode(node ast.Node) Node {
	return creator(baseNode{node: node})
}

//NewFileNode creates a new file node including raw content for regex searches. Use NewNode to create
// a file without regex capabilities.
func NewFileNode(node *ast.File, rawFile *RawFile) *File {
	file := creator(baseNode{
		node: node,
	}).(*File)
	file.rawFile = rawFile
	return file
}

func newChild(node ast.Node, parent Node, parentLevel int) Node {
	rawFile := parent.getRawFile(node)

	return creator(baseNode{
		node:    node,
		parent:  parent,
		level:   parentLevel + 1,
		rawFile: rawFile,
	})
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
	NextParentByType(nodeType NodeType) Node
	IsContainedByType(nodeType NodeType) bool
	Contains(node Node) bool
	FindByName(name string) []Node
	FindFirstByName(name string) Node
	ChildByName(name string) Node
	ChildByNodeType(nodeType NodeType) Node
	FindIdentByName(name string) []*Ident
	FindFirstIdentByName(name string) *Ident
	FindByNodeType(nodeType NodeType) []Node
	FindFirstByNodeType(nodeType NodeType) Node
	FindByValueType(valType string) []Node
	FindMaps() []Node
	IsNodeType(nodeType NodeType) bool
	NodeType() NodeType
	IsValueType(valType string) bool
	ValueType() types.Type
	Object() types.Object
	GetSource() []byte

	findByName(name string, identOnly, firstOnly, childOnly bool) []Node
	setRealMe(node Node)
	getRawFile(node ast.Node) *RawFile
}

type baseNode struct {
	node ast.Node

	realMe   Node
	parent   Node
	level    int
	built    bool
	children []Node

	nodeType NodeType
	rawFile  *RawFile
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

//NextParentByType returns the next parent of given type
func (s *baseNode) NextParentByType(nodeType NodeType) Node {
	if s.parent == nil {
		return nil
	}
	if s.parent.IsNodeType(nodeType) {
		return s.parent
	}
	return s.parent.NextParentByType(nodeType)
}

//IsContainedByType checks if node is contained by a node of given node type
func (s *baseNode) IsContainedByType(nodeType NodeType) bool {
	if s.parent == nil {
		return false
	}
	if s.parent.IsNodeType(nodeType) {
		return true
	}
	return s.parent.IsContainedByType(nodeType)
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

//IsNodeType checks if node is of given node type
func (s *baseNode) IsNodeType(nodeType NodeType) bool {
	return s.nodeType == nodeType
}

//NodeType returns the NodeType of the node
func (s *baseNode) NodeType() NodeType {
	return s.nodeType
}

//FindByName looks for a name in the entire sub tree
func (s *baseNode) FindByName(name string) []Node {
	return s.findByName(name, false, false, false)
}

//FindFirstByName looks for a name in the entire sub tree. First node is returned if there are multiple.
func (s *baseNode) FindFirstByName(name string) Node {
	nodes := s.findByName(name, false, true, false)
	for _, node := range nodes {
		return node
	}
	return nil
}

//FindIdentByName looks for Ident nodes in the entire sub tree with given name
func (s *baseNode) FindIdentByName(name string) []*Ident {
	nodes := s.findByName(name, true, false, false)
	var idents []*Ident
	for _, node := range nodes {
		idents = append(idents, node.(*Ident))
	}
	return idents
}

//FindFirstIdentByName looks for the first Ident node in subtree with given name
func (s *baseNode) FindFirstIdentByName(name string) *Ident {
	nodes := s.findByName(name, true, false, false)
	for _, node := range nodes {
		return node.(*Ident)
	}
	return nil
}

//ChildByName retrieves a node among the direkt children by name (only nodes that have a name)
func (s *baseNode) ChildByName(name string) Node {
	nodes := s.findByName(name, false, true, true)
	for _, node := range nodes {
		return node
	}
	return nil
}

//ChildByNodeType returns the first child of a certain type.
func (s *baseNode) ChildByNodeType(nodeType NodeType) Node {
	for _, child := range s.Children() {
		if child.IsNodeType(nodeType) {
			return child
		}
	}
	return nil
}

func (s *baseNode) findByName(name string, identOnly, firstOnly, childOnly bool) []Node {
	var nodes []Node
	for _, child := range s.Children() {
		valid := !identOnly
		if _, ok := child.(*Ident); ok {
			valid = true
		}

		if f, ok := child.(Named); ok {
			ident := f.NodeName()
			if ident != nil && ident.Name == name {
				if valid {
					nodes = append(nodes, child)
					if firstOnly {
						return nodes
					}
				}
			}
		}

		if childOnly {
			continue
		}
		if f := child.findByName(name, identOnly, firstOnly, childOnly); f != nil {
			if firstOnly {
				return f
			}
			nodes = append(nodes, f...)
		}
	}
	return nodes
}

//FindByNodeType returns all sub nodes of a certain type
func (s *baseNode) FindByNodeType(nodeType NodeType) []Node {
	var nodes []Node
	for _, child := range s.Children() {
		if child.IsNodeType(nodeType) {
			nodes = append(nodes, child)
		}
		nodes = append(nodes, child.FindByNodeType(nodeType)...)
	}
	return nodes
}

//FindFirstByNodeType returns the first sub node of a certain type
func (s *baseNode) FindFirstByNodeType(nodeType NodeType) Node {
	for _, child := range s.Children() {
		if child.IsNodeType(nodeType) {
			return child
		}
		if n := child.FindFirstByNodeType(nodeType); n != nil {
			return n
		}
	}
	return nil
}

//FindByValueType find all nodes with given value type
func (s *baseNode) FindByValueType(valType string) []Node {
	var nodes []Node
	for _, child := range s.Children() {
		if child.IsValueType(valType) {
			nodes = append(nodes, child)
		}
		nodes = append(nodes, child.FindByValueType(valType)...)
	}
	return nodes
}

//FindMaps find all nodes with given value type
func (s *baseNode) FindMaps() []Node {
	var nodes []Node
	for _, child := range s.Children() {
		valueType := child.ValueType()
		if valueType != nil && strings.HasPrefix(valueType.String(), "map") {
			nodes = append(nodes, child)
		}
		nodes = append(nodes, child.FindMaps()...)
	}
	return nodes
}

//IsValueType checks if value type is of given type
func (s *baseNode) IsValueType(valType string) bool {
	if expr, ok := s.node.(ast.Expr); ok {
		info.TypeOf(expr)
		if t, ok := info.Types[expr]; ok {
			if t.Type.String() == valType {
				return true
			}
		}
	}
	return false
}

//ValueType returns value type information of an expression, nil otherwise
func (s *baseNode) ValueType() types.Type {
	if expr, ok := s.node.(ast.Expr); ok {
		return info.TypeOf(expr)
	}
	return nil
}

//Object returns the object of an identifier, nil otherwise
func (s *baseNode) Object() types.Object {
	if expr, ok := s.node.(*ast.Ident); ok {
		return info.ObjectOf(expr)
	}
	return nil
}

//GetSource returns the source code of the current node
func (s *baseNode) GetSource() []byte {
	if s.nodeType == NodeTypePackage {
		var sources [][]byte
		for _, rawFile := range s.realMe.(*Package).rawFiles {
			sources = append(sources, rawFile.source)
		}
		return bytes.Join(sources, []byte{'\n'})
	} else if s.nodeType == NodeTypeFile {
		return s.rawFile.source
	}

	base := token.Pos(s.rawFile.Base())
	return s.rawFile.source[s.node.Pos()-base : s.node.End()-base]
}

//Match matches the source code of current node and content with given regex
func (s *baseNode) Match(regex regexp.Regexp) bool {
	return regex.Match(s.rawFile.source)
}

func (s *baseNode) setRealMe(node Node) {
	s.realMe = node
	s.nodeType = NodeType(reflect.TypeOf(node).String())
}

func (s *baseNode) getRawFile(node ast.Node) *RawFile {
	switch p := s.realMe.(type) {
	case *Package:
		if n, ok := node.(*ast.File); ok {
			for _, rf := range p.rawFiles {
				if rf.ContainsPos(n.Pos()) {
					return rf
				}
			}
		}
	}
	return s.rawFile
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

	child := newChild(node, s.realMe, s.level)
	s.children = append(s.children, child)
	return nil
}
