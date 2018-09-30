package astrav

import "go/ast"

//Package wraps ast.Package
type Package struct {
	*ast.Package
	baseNode

	rawFiles map[string]*RawFile
}
