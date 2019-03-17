package astrav

import (
	"go/ast"
	"go/types"
)

// ValueTyper provides an interface for nodes with a value type.
type ValueTyper interface {
	ValueType() types.Type
}

// ValueType returns the value type of the node.
func (s *Field) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return s.Info().ObjectOf(s.Type.(*ast.Ident)).Type()
}

// ValueType returns the value type of the node.
func (s *CompositeLit) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return s.Info().ObjectOf(s.Type.(*ast.Ident)).Type()
}

// ValueType returns the value type of the node.
func (s *TypeAssertExpr) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return s.Info().ObjectOf(s.Type.(*ast.Ident)).Type()
}

// ValueType returns the value type of the node.
func (s *ValueSpec) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return s.Info().ObjectOf(s.Type.(*ast.Ident)).Type()
}

// ValueType returns the value type of the node.
func (s *TypeSpec) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return s.Info().ObjectOf(s.Type.(*ast.Ident)).Type()
}

// ValueType returns the value type of the node.
func (s *Ident) ValueType() types.Type {
	if t := s.Info().TypeOf(s); t != nil {
		return t
	}

	return s.Info().ObjectOf(s.node.(*ast.Ident)).Type()
}
