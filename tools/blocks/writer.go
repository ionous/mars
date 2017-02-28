package blocks

import (
	"io"
	"strings"
	"unicode"
)

type Words interface {
	WriteWord(string)
}

type WordWriter struct {
	out   io.Writer
	space bool
}

func NewWordWriter(out io.Writer) *WordWriter {
	return &WordWriter{out, false}
}

func (wm *WordWriter) WriteWord(s string) {
	if wm.space {
		i := strings.IndexFunc(s, func(r rune) bool {
			limitedPunct := r != '[' && r != ']' && r != ')' && r != ')' && unicode.IsPunct(r)
			return limitedPunct
		})
		if i != 0 {
			io.WriteString(wm.out, " ")
		}
	}
	io.WriteString(wm.out, s)
	wm.space = true
}
