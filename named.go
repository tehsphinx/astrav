package astrav

//Named provides an interface for nodes with a name
type Named interface {
	Ident() *Ident
}

//Ident returns the name of the node
func (s *FuncDecl) Ident() *Ident {
	return New(s.Name).(*Ident)
}

//Ident returns the name of the node
func (s *LabeledStmt) Ident() *Ident {
	return New(s.Label).(*Ident)
}

//Ident returns the name of the node
func (s *BranchStmt) Ident() *Ident {
	return New(s.Label).(*Ident)
}

//Ident returns the name of the node
func (s *ImportSpec) Ident() *Ident {
	return New(s.Name).(*Ident)
}

//Ident returns the name of the node
func (s *File) Ident() *Ident {
	return New(s.Name).(*Ident)
}

//Ident returns the name of the node
func (s *SelectorExpr) Ident() *Ident {
	return New(s.Sel).(*Ident)
}

//Ident returns the name of the node
func (s *TypeSpec) Ident() *Ident {
	return New(s.Name).(*Ident)
}
