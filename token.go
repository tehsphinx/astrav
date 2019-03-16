package astrav

import (
	"go/token"
)

// Token provides an interface for nodes with a token.Token
type Token interface {
	Token() token.Token
}

// Token returns the token of the node
func (s *AssignStmt) Token() token.Token {
	return s.Tok
}

// Token returns the token of the node
func (s *BasicLit) Token() token.Token {
	return s.Kind
}

// Token returns the token of the node
func (s *BinaryExpr) Token() token.Token {
	return s.Op
}

// Token returns the token of the node
func (s *BranchStmt) Token() token.Token {
	return s.Tok
}

// Token returns the token of the node
func (s *GenDecl) Token() token.Token {
	return s.Tok
}

// Token returns the token of the node
func (s *IncDecStmt) Token() token.Token {
	return s.Tok
}

// Token returns the token of the node
func (s *RangeStmt) Token() token.Token {
	return s.Tok
}

// Token returns the token of the node
func (s *UnaryExpr) Token() token.Token {
	return s.Op
}
