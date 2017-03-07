package blocks

import (
	"io"
)

//
type Scope struct {
	w     io.Writer
	depth []byte
}

func (r *Scope) changeIndent(okay bool) {
	if okay {
		r.depth = append(r.depth, '\t')
	} else if cnt := len(r.depth); cnt == 0 {
		panic("indent out of depth")
	} else {
		r.depth = r.depth[:cnt-1]
	}
}

func (r *Scope) writeIndent() {
	if len(r.depth) > 0 {
		r.w.Write(r.depth)
	}
}
