package astrav

import (
	"go/ast"
	"go/types"
)

// ValueTyper provides an interface for nodes with a value type.
type ValueTyper interface {
	ValueType() types.Type
}

func (s *baseNode) getType(node ast.Node) types.Type {
	if node == nil {
		return nil
	}
	ident, ok := node.(*ast.Ident)
	if !ok {
		return nil
	}
	obj := s.Info().ObjectOf(ident)
	if obj == nil {
		return nil
	}
	return s.Info().ObjectOf(ident).Type()
}

// ValueType returns the value type of the node.
func (s *Field) ValueType() types.Type {
	return s.getType(s.Type)
}

// ValueType returns the value type of the node.
func (s *CompositeLit) ValueType() types.Type {
	return s.getType(s.Type)
}

// ValueType returns the value type of the node.
func (s *TypeAssertExpr) ValueType() types.Type {
	return s.getType(s.Type)

}

// ValueType returns the value type of the node.
func (s *ValueSpec) ValueType() types.Type {
	return s.getType(s.Type)

}

// ValueType returns the value type of the node.
func (s *TypeSpec) ValueType() types.Type {
	return s.getType(s.Type)

}

// ValueType returns the value type of the node.
func (s *Ident) ValueType() types.Type {
	if t := s.Info().TypeOf(s); t != nil {
		return t
	}

	return s.getType(s.node)
}
