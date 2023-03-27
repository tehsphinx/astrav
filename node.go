package astrav

import (
	"bytes"
	"go/ast"
	"go/token"
	"go/types"
	"math"
	"reflect"
	"regexp"
	"strings"
)

// NewNode creates a new node
func NewNode(node ast.Node) Node {
	return creator(baseNode{node: node})
}

// NewFileNode creates a new file node including raw content for regex searches. Use NewNode to create
// a file without regex capabilities.
func NewFileNode(node *ast.File, rawFile *RawFile) *File {
	file := creator(baseNode{
		node: node,
	}).(*File)
	file.rawFile = rawFile
	return file
}

func newChild(node ast.Node, parent Node, pkg *Package, parentLevel int) Node {
	rawFile := parent.getRawFile(node)

	return creator(baseNode{
		node:    node,
		parent:  parent,
		pkg:     pkg,
		level:   parentLevel + 1,
		rawFile: rawFile,
	})
}

// Node wraps a ast.Node with helpful traversal functions
type Node interface {
	ast.Node

	Level() int
	AstNode() ast.Node
	IsNodeType(nodeType NodeType) bool
	NodeType() NodeType
	IsValueType(valType string) bool
	ValueType() types.Type
	Object() types.Object
	Pkg() *Package
	Info() *types.Info
	Walk(f func(node Node) bool)

	GetScope() (Node, *types.Scope)
	FindByPos(pos token.Pos) (Node, bool)
	Parent() Node
	Parents() []Node
	NextParentByType(nodeType NodeType) Node
	IsContainedByType(nodeType NodeType) bool
	Siblings() []Node
	Children() []Node
	Contains(node Node) bool
	ChildByName(name string) Node
	ChildrenByNodeType(nodeType NodeType) []Node
	ChildByNodeType(nodeType NodeType) Node

	FindByName(name string) []Node
	FindFirstByName(name string) Node
	FindIdentByName(name string) []*Ident
	FindFirstIdentByName(name string) *Ident
	FindNameInCallTree(name string) []Node
	FindByNodeType(nodeType NodeType) []Node
	FindFirstByNodeType(nodeType NodeType) Node
	FindNodeTypeInCallTree(nodeType NodeType) []Node
	FindByValueType(valType string) []Node
	FindByToken(t token.Token) []Node
	FindMaps() []Node
	FindDeclarations() []*Ident
	FindDeclarationsByType(nodeType NodeType) []*Ident
	FindVarDeclarations() []*Ident
	FindDeclaration(usage *Ident) *Ident
	FindUsages(*Ident) []*Ident

	ChildNodes(cond func(n Node) bool) []Node
	ChildNode(cond func(n Node) bool) Node
	TreeNodes(cond func(n Node) bool) []Node
	TreeNode(cond func(n Node) bool) Node
	CallTreeNodes(cond func(n Node) bool) []Node
	callTreeNodes(cond func(n Node) bool, visited map[Node]bool) []Node
	CallTreeNode(cond func(n Node) bool) Node

	GetSource() []byte
	GetSourceString() string

	setRealMe(node Node)
	getRawFile(node ast.Node) *RawFile
}

type baseNode struct {
	node ast.Node

	realMe   Node
	parent   Node
	pkg      *Package
	level    int
	built    bool
	children []Node

	nodeType NodeType
	rawFile  *RawFile

	usageCache map[*Ident][]*Ident
	declCache  map[*Ident]*Ident
}

// Parent return the parent node
func (s *baseNode) Parent() Node {
	return s.parent
}

// Children returns all child nodes
func (s *baseNode) Children() []Node {
	s.build()
	return s.children
}

// Siblings returns all sibling nodes
func (s *baseNode) Siblings() []Node {
	var nodes []Node
	for _, node := range s.parent.Children() {
		if node != s.realMe {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

// Level returns the level counted from instantiated node = 0
func (s *baseNode) Level() int {
	return s.level
}

// AstNode returns the original ast.Node
func (s *baseNode) AstNode() ast.Node {
	return s.node
}

// Walk traverses the tree and its children.
// return false to skip children of the current element
func (s *baseNode) Walk(f func(node Node) bool) {
	if !f(s.realMe) {
		return
	}

	for _, child := range s.Children() {
		child.Walk(f)
	}
}

// Parents returns the parent path of nodes
func (s *baseNode) Parents() []Node {
	if s.parent == nil {
		return nil
	}
	return append([]Node{s.parent}, s.parent.Parents()...)
}

// NextParentByType returns the next parent of given type
func (s *baseNode) NextParentByType(nodeType NodeType) Node {
	if s.parent == nil {
		return nil
	}
	if s.parent.IsNodeType(nodeType) {
		return s.parent
	}
	return s.parent.NextParentByType(nodeType)
}

// IsContainedByType checks if node is contained by a node of given node type
func (s *baseNode) IsContainedByType(nodeType NodeType) bool {
	if s.parent == nil {
		return false
	}
	if s.parent.IsNodeType(nodeType) {
		return true
	}
	return s.parent.IsContainedByType(nodeType)
}

// Contains checks if a node contains another node
func (s *baseNode) Contains(node Node) bool {
	for _, p := range node.Parents() {
		if p == s.realMe {
			return true
		}
	}
	return false
}

// IsNodeType checks if node is of given node type
func (s *baseNode) IsNodeType(nodeType NodeType) bool {
	return s.nodeType == nodeType
}

// NodeType returns the NodeType of the node
func (s *baseNode) NodeType() NodeType {
	return s.nodeType
}

// FindByName looks for a name in the entire sub tree
func (s *baseNode) FindByName(name string) []Node {
	return s.TreeNodes(func(n Node) bool {
		f, ok := n.(Named)
		if !ok {
			return false
		}

		return f.NodeName() == name
	})
}

// FindNameInCallTree returns all nodes in call tree with given name
func (s *baseNode) FindNameInCallTree(name string) []Node {
	return s.CallTreeNodes(func(n Node) bool {
		f, ok := n.(Named)
		if !ok {
			return false
		}

		return f.NodeName() == name
	})
}

// FindFirstByName looks for a name in the entire sub tree. First node is returned if there are multiple.
func (s *baseNode) FindFirstByName(name string) Node {
	return s.TreeNode(func(n Node) bool {
		f, ok := n.(Named)
		if !ok {
			return false
		}

		return f.NodeName() == name
	})
}

// FindIdentByName looks for Ident nodes in the entire sub tree with given name
func (s *baseNode) FindIdentByName(name string) []*Ident {
	nodes := s.TreeNodes(func(n Node) bool {
		id, ok := n.(*Ident)
		if !ok {
			return false
		}

		return id.Name == name
	})

	var idents []*Ident
	for _, node := range nodes {
		idents = append(idents, node.(*Ident))
	}
	return idents
}

// FindFirstIdentByName looks for the first Ident node in subtree with given name
func (s *baseNode) FindFirstIdentByName(name string) *Ident {
	ident := s.TreeNode(func(n Node) bool {
		id, ok := n.(*Ident)
		if !ok {
			return false
		}

		return id.Name == name
	})
	if ident == nil {
		return nil
	}

	return ident.(*Ident)
}

// ChildByName retrieves a node among the direkt children by name (only nodes that have a name)
func (s *baseNode) ChildByName(name string) Node {
	return s.ChildNode(func(n Node) bool {
		f, ok := n.(Named)
		if !ok {
			return false
		}

		return f.NodeName() == name
	})
}

// ChildrenByNodeType returns the first child of a certain type.
func (s *baseNode) ChildrenByNodeType(nodeType NodeType) []Node {
	return s.ChildNodes(func(n Node) bool {
		return n.IsNodeType(nodeType)
	})
}

// ChildByNodeType returns the first child of a certain type.
func (s *baseNode) ChildByNodeType(nodeType NodeType) Node {
	return s.ChildNode(func(n Node) bool {
		return n.IsNodeType(nodeType)
	})
}

// FindByNodeType returns all sub nodes of a certain type
func (s *baseNode) FindByNodeType(nodeType NodeType) []Node {
	return s.TreeNodes(func(n Node) bool {
		return n.IsNodeType(nodeType)
	})
}

// FindNodeTypeInCallTree returns all nodes in call tree of a certain type
func (s *baseNode) FindNodeTypeInCallTree(nodeType NodeType) []Node {
	return s.CallTreeNodes(func(n Node) bool {
		return n.IsNodeType(nodeType)
	})
}

// FindFirstByNodeType returns the first sub node of a certain type
func (s *baseNode) FindFirstByNodeType(nodeType NodeType) Node {
	return s.TreeNode(func(n Node) bool {
		return n.IsNodeType(nodeType)
	})
}

// FindByValueType find all nodes with given value type
func (s *baseNode) FindByValueType(valType string) []Node {
	return s.TreeNodes(func(n Node) bool {
		return n.IsValueType(valType)
	})
}

// FindByToken finds a node with a token.Token attached and of given token type
func (s *baseNode) FindByToken(t token.Token) []Node {
	return s.TreeNodes(func(n Node) bool {
		tok, ok := n.(Token)
		return ok && tok.Token() == t
	})
}

// FindMaps find all nodes with given value type
func (s *baseNode) FindMaps() []Node {
	return s.TreeNodes(func(n Node) bool {
		valueType := n.ValueType()
		return valueType != nil && strings.HasPrefix(valueType.String(), "map")
	})
}

// FindDeclarations finds all declarations
func (s *baseNode) FindDeclarations() []*Ident {
	return s.findIdents(s.Info().Defs, func(node *Ident) bool {
		return true
	})
}

// FindDeclarationsByType finds all declarations by (parent) type
func (s *baseNode) FindDeclarationsByType(nodeType NodeType) []*Ident {
	return s.findIdents(s.Info().Defs, func(node *Ident) bool {
		return node.Parent().IsNodeType(nodeType)
	})
}

// FindVarDeclarations finds all assignStmt (:=), valueSpec (var, const) declarations
func (s *baseNode) FindVarDeclarations() []*Ident {
	return s.findIdents(s.Info().Defs, func(node *Ident) bool {
		return node.Parent().IsNodeType(NodeTypeAssignStmt) ||
			node.Parent().IsNodeType(NodeTypeValueSpec)
	})
}

// FindDeclaration finds the declaration for a usage
func (s *baseNode) FindDeclaration(usage *Ident) *Ident {
	if decl, ok := s.cachedDeclaration(usage); ok {
		return decl
	}

	decls := s.findIdents(s.Info().Defs, func(node *Ident) bool {
		return node.Name == usage.Name
	})
	if len(decls) == 1 {
		return decls[0]
	}

	var (
		correctDecl  *Ident
		correctScope *types.Scope
	)
	for _, decl := range decls {
		scopeNode, scope := decl.GetScope()
		if !scope.Contains(usage.Pos()) {
			continue
		}
		if correctDecl != nil && correctScope != nil && !correctScope.Contains(scopeNode.Pos()) {
			continue
		}
		correctDecl = decl
		correctScope = scope
	}

	s.cacheDeclaration(usage, correctDecl)
	return correctDecl
}

func (s *baseNode) cachedDeclaration(usage *Ident) (*Ident, bool) {
	if s.declCache == nil {
		return nil, false
	}

	decl, ok := s.declCache[usage]
	return decl, ok
}

func (s *baseNode) cacheDeclaration(usage *Ident, decl *Ident) {
	if s.declCache == nil {
		s.declCache = map[*Ident]*Ident{}
	}

	s.declCache[usage] = decl
}

func (s *baseNode) findIdents(search map[*ast.Ident]types.Object, f func(node *Ident) bool) []*Ident {
	var decls []*Ident
	for astIdent := range search {
		node := s.findChildByAstNode(astIdent)
		if node == nil {
			continue
		}
		ident := node.(*Ident)
		if ident.Name == "_" || !f(ident) {
			continue
		}
		decls = append(decls, ident)
	}
	return decls
}

// FindUsages finds usages of given declaration
func (s *baseNode) FindUsages(declaration *Ident) []*Ident {
	if usages, ok := s.cachedUsages(declaration); ok {
		return usages
	}

	usgs := s.findIdents(s.Info().Uses, func(node *Ident) bool {
		return node.Name == declaration.Name
	})
	otherDecls := s.findIdents(s.Info().Defs, func(node *Ident) bool {
		_, scope := declaration.GetScope()
		return node.Name == declaration.Name && scope.Contains(node.Pos()) && node.Pos() != declaration.AstNode().Pos()
	})

	var usages []*Ident
	_, scope := declaration.GetScope()
	for _, usg := range usgs {
		if !scope.Contains(usg.Pos()) {
			continue
		}

		var found bool
		for _, otherDecl := range otherDecls {
			_, scope := otherDecl.GetScope()
			if scope.Contains(usg.Pos()) {
				found = true
				break
			}
		}
		if found {
			continue
		}

		usages = append(usages, usg)
	}

	s.cacheUsages(declaration, usages)
	return usages
}

func (s *baseNode) cachedUsages(declaration *Ident) ([]*Ident, bool) {
	if s.usageCache == nil {
		return nil, false
	}

	usages, ok := s.usageCache[declaration]
	return usages, ok
}

func (s *baseNode) cacheUsages(declaration *Ident, usages []*Ident) {
	if s.usageCache == nil {
		s.usageCache = map[*Ident][]*Ident{}
	}

	s.usageCache[declaration] = usages
}

// FindFirstUsage selects the first usage
func (s *baseNode) FindFirstUsage(declaration *Ident) *Ident {
	usgs := s.FindUsages(declaration)

	var (
		firstUsage *Ident
		minPos     token.Pos = math.MaxInt32
	)
	for _, usage := range usgs {
		if usage.NamePos < minPos {
			minPos = usage.NamePos
			firstUsage = usage
		}
	}
	return firstUsage
}

// GetScope returns the scope of the node
func (s *baseNode) GetScope() (Node, *types.Scope) {
	var (
		maxScopePos token.Pos
		maxNode     ast.Node
		maxScope    *types.Scope
	)
	for node, scope := range s.Info().Scopes {
		if scope.Contains(s.realMe.Pos()) && maxScopePos < scope.Pos() {
			maxScopePos = scope.Pos()
			maxScope = scope
			maxNode = node
		}
	}
	return s.Pkg().findChildByAstNode(maxNode), maxScope
}

// FindByPos finds a node by the given position. If there is no exact match, the closest node
// before the position is returned.
func (s *baseNode) FindByPos(pos token.Pos) (Node, bool) {
	_, scope := s.realMe.GetScope()
	if scope == nil {
		return nil, false
	}
	if !scope.Contains(pos) {
		return nil, false
	}

	const maxDiff = 50
	var (
		minDiff = maxDiff
		retNode Node
	)

	s.Walk(func(node Node) bool {
		if node == nil {
			return true
		}

		if pos == node.Pos() {
			retNode = node
			return false
		}

		diff := int(pos - node.Pos())
		if 0 <= diff && diff < maxDiff {
			if diff < minDiff {
				retNode = node
			}
		}
		return true
	})

	if retNode == nil {
		return nil, false
	}
	return retNode, true
}

// IsValueType checks if value type is of given type
func (s *baseNode) IsValueType(valType string) bool {
	if expr, ok := s.node.(ast.Expr); ok {
		if t, ok := s.Info().Types[expr]; ok {
			if t.Type.String() == valType {
				return true
			}
		}
		if t := s.Info().TypeOf(expr); t != nil {
			if t.String() == valType {
				return true
			}
		}
	}
	return false
}

// ValueType returns value type information of an expression, nil otherwise
func (s *baseNode) ValueType() types.Type {
	if expr, ok := s.node.(ast.Expr); ok {
		return s.Info().TypeOf(expr)
	}
	return nil
}

// Object returns the object of an identifier, nil otherwise
func (s *baseNode) Object() types.Object {
	if expr, ok := s.node.(*ast.Ident); ok {
		return s.Info().ObjectOf(expr)
	}
	return nil
}

// Info returns the types.info node.
func (s *baseNode) Info() *types.Info {
	return s.Pkg().info
}

// Pkg returns the package this node belongs to.
func (s *baseNode) Pkg() *Package {
	if s.nodeType == NodeTypePackage {
		return s.realMe.(*Package)
	}
	if s.pkg == nil {
		return nil
	}
	return s.pkg
}

// GetSource returns the source code of the current node
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

	return s.getSourceRange(s.node.Pos(), s.node.End())
}

func (s *baseNode) getSourceRange(pos, end token.Pos) []byte {
	base := token.Pos(s.rawFile.Base())
	return s.rawFile.source[pos-base : end-base]
}

// GetSourceString is a convenience function to GetSource as string
func (s *baseNode) GetSourceString() string {
	return string(s.GetSource())
}

// Match matches the source code of current node and content with given regex
func (s *baseNode) Match(regex regexp.Regexp) bool {
	return regex.Match(s.rawFile.source)
}

func (s *baseNode) setRealMe(node Node) {
	s.realMe = node
	s.nodeType = NodeType(reflect.TypeOf(node).String())
}

func (s *baseNode) getRawFile(node ast.Node) *RawFile {
	pkg, ok := s.realMe.(*Package)
	if ok {
		if n, ok := node.(*ast.File); ok {
			for _, rf := range pkg.rawFiles {
				if rf.ContainsPos(n.Pos()) {
					return rf
				}
			}
		}
	}
	return s.rawFile
}

func (s *baseNode) findChildByAstNode(astNode ast.Node) Node {
	return s.TreeNode(func(n Node) bool {
		return n.AstNode() == astNode
	})
}

// ChildNodes walks only child nodes collecting all nodes that meet the condition
func (s *baseNode) ChildNodes(cond func(n Node) bool) []Node {
	var nodes []Node
	for _, child := range s.Children() {
		if cond(child) {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// ChildNode walks only child nodes returning the first node that meets the condition
func (s *baseNode) ChildNode(cond func(n Node) bool) Node {
	for _, child := range s.Children() {
		if cond(child) {
			return child
		}
	}
	return nil
}

// TreeNodes walks the child tree collecting all nodes that meet the condition
func (s *baseNode) TreeNodes(cond func(n Node) bool) []Node {
	var nodes []Node
	for _, child := range s.Children() {
		if cond(child) {
			nodes = append(nodes, child)
		}
		nodes = append(nodes, child.TreeNodes(cond)...)
	}
	return nodes
}

// TreeNode walks the child tree returning the first node that meets the condition
func (s *baseNode) TreeNode(cond func(n Node) bool) Node {
	for _, child := range s.Children() {
		if cond(child) {
			return child
		}
		if n := child.TreeNode(cond); n != nil {
			return n
		}
	}
	return nil
}

// CallTreeNodes walks the call tree collecting all nodes that meet the condition
func (s baseNode) CallTreeNodes(cond func(n Node) bool) []Node {
	return s.callTreeNodes(cond, map[Node]bool{})
}

func (s baseNode) callTreeNodes(cond func(n Node) bool, visited map[Node]bool) []Node {
	if s.Pkg() == nil {
		return nil
	}

	var nodes []Node
	for _, child := range s.Children() {
		if cond(child) {
			nodes = append(nodes, child)
		}

		if n := s.callNode(child); n != nil {
			func() {
				if visited[n] {
					return
				}
				visited[n] = true
				if cond(n) {
					nodes = append(nodes, n)
				}
				// TODO: circular call trees
				nodes = append(nodes, n.callTreeNodes(cond, visited)...)
			}()
		}

		nodes = append(nodes, child.callTreeNodes(cond, visited)...)
	}
	return nodes
}

// CallTreeNode walks the call tree returning the first node that meets the condition
func (s baseNode) CallTreeNode(cond func(n Node) bool) Node {
	if s.Pkg() == nil {
		return nil
	}

	for _, child := range s.Children() {
		if cond(child) {
			return child
		}

		if n := s.callNode(child); n != nil {
			if cond(n) {
				return n
			}
			// TODO: circular call trees
			if node := n.CallTreeNode(cond); node != nil {
				return node
			}
		}

		if node := child.CallTreeNode(cond); node != nil {
			return node
		}
	}
	return nil
}

func (s *baseNode) callNode(n Node) Node {
	if n.NodeType() == NodeTypeCallExpr {
		if node := s.Pkg().FuncDeclbyCallExpr(n.(*CallExpr)); node != nil {
			return node
		}
	}
	return nil
}

func (s *baseNode) build() {
	if s.built {
		return
	}
	s.built = true

	ast.Walk(s, s.node)
}

// Visit implements the ast.Visitor interface and is used to walk the underlying ast.Node tree
func (s *baseNode) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	if node == s.node {
		return s
	}

	pkg := s.Pkg()
	switch n := node.(type) {
	case *ast.Field:
		if len(n.Names) < 2 {
			child := newChild(node, s.realMe, pkg, s.level)
			s.children = append(s.children, child)
			break
		}
		// splitting one field with multiple names into multiple fields
		for _, name := range n.Names {
			newNode := &ast.Field{
				Comment: n.Comment,
				Type:    n.Type,
				Doc:     n.Doc,
				Tag:     n.Tag,
				Names:   []*ast.Ident{name},
			}
			child := newChild(newNode, s.realMe, pkg, s.level)
			s.children = append(s.children, child)
		}
	default:
		child := newChild(node, s.realMe, pkg, s.level)
		s.children = append(s.children, child)
	}
	return nil
}
