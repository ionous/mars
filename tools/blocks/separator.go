package blocks

import (
	"io"
	"strings"
	"unicode"
)

//
type Separator struct {
	w     io.Writer
	chars string
}

// separator
func (s *Separator) writeEnd() {
	// remove any final spaces
	if chars := strings.TrimRightFunc(s.chars, unicode.IsSpace); len(chars) > 0 {
		io.WriteString(s.w, chars)
	}
	s.chars = ""
}

func (s *Separator) separate(chars string) {
	s.chars = chars
}

func (s *Separator) flushSep() {
	if len(s.chars) > 0 {
		io.WriteString(s.w, s.chars)
	}
	s.chars = " "
	return
}
