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
	return stringToType(s.Type.(*ast.Ident).Name)
}

// ValueType returns the value type of the node.
func (s *CompositeLit) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return stringToType(s.Type.(*ast.Ident).Name)
}

// ValueType returns the value type of the node.
func (s *TypeAssertExpr) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return stringToType(s.Type.(*ast.Ident).Name)
}

// ValueType returns the value type of the node.
func (s *ValueSpec) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return stringToType(s.Type.(*ast.Ident).Name)
}

// ValueType returns the value type of the node.
func (s *TypeSpec) ValueType() types.Type {
	if s.Type == nil {
		return nil
	}
	return stringToType(s.Type.(*ast.Ident).Name)
}

func stringToType(typeName string) types.Type {
	for _, t := range types.Typ {
		if t.Name() == typeName {
			return t
		}
	}
	return nil
}
