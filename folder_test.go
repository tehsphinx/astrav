package astrav

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFolder_ParseFolder(t *testing.T) {
	f := NewFolder("exercism/solution", http.Dir("./example/1"))
	pkgs, err := f.ParseFolder()
	if err != nil {
		t.Error(err)
	}
	for _, pkg := range pkgs {
		assert.Equal(t, "scrabble", pkg.Name)
	}

	for fileName, value := range f.RawFiles {
		assert.Equal(t, "example.go", fileName)
		assert.Equal(t, 549, len(value.source))
	}
}
