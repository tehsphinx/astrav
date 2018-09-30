package astrav

//Named provides an interface for nodes with a name
type Named interface {
	NodeName() *Ident
}

//NodeName returns the name of the node
func (s *FuncDecl) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}

	return newChild(s.Name, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *LabeledStmt) NodeName() *Ident {
	if s.Label == nil {
		return nil
	}
	return newChild(s.Label, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *BranchStmt) NodeName() *Ident {
	if s.Label == nil {
		return nil
	}
	return newChild(s.Label, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *ImportSpec) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return newChild(s.Name, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *File) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return newChild(s.Name, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *SelectorExpr) NodeName() *Ident {
	if s.Sel == nil {
		return nil
	}
	return newChild(s.Sel, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *TypeSpec) NodeName() *Ident {
	if s.Name == nil {
		return nil
	}
	return newChild(s.Name, s.parent, s.level).(*Ident)
}

//NodeName returns the name of the node
func (s *Ident) NodeName() *Ident {
	return s
}
