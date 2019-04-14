package astrav

// Named provides an interface for nodes with a name
type Named interface {
	// NodeName returns the name of the node
	NodeName() string
}

func getIdentName(s Identifier) string {
	ident := s.GetIdent()
	if ident == nil {
		return ""
	}
	return ident.Name
}

// NodeName returns the name of the node
func (s *FuncDecl) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *LabeledStmt) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *BranchStmt) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *ImportSpec) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *File) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *SelectorExpr) NodeName() string {
	name := getIdentName(s)
	if sel := s.findChildByAstNode(s.X); sel != nil && sel.IsNodeType(NodeTypeIdent) {
		return sel.(*Ident).NodeName() + "." + name
	}
	return name
}

// NodeName returns the name of the node
func (s *TypeSpec) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *CallExpr) NodeName() string {
	if s.Fun == nil {
		return ""
	}

	node := s.findChildByAstNode(s.Fun)
	switch t := node.(type) {
	case *ArrayType:
		return t.NodeName()
	}

	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *Ident) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *Field) NodeName() string {
	return getIdentName(s)
}

// NodeName returns the name of the node
func (s *ArrayType) NodeName() string {
	return "[]" + getIdentName(s)
}
