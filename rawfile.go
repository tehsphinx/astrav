package astrav

import "go/token"

//NewRawFile creates a new RawFile. RawFile is based on token.File and contains the source code
func NewRawFile(file *token.File, source []byte) *RawFile {
	return &RawFile{
		File:   file,
		source: source,
	}
}

//RawFile is based on token.File to add the source code of the file
type RawFile struct {
	*token.File

	source []byte
}

//Source returns the source code of the file
func (s *RawFile) Source() []byte {
	return s.source
}
