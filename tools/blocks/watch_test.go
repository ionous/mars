package blocks

import (
	"bytes"
	"strings"
	"testing"
)

type WatchTerms struct {
	t   *testing.T
	src GenerateTerms
}

func (w WatchTerms) GenerateTerms(n *DocNode) TermSet {
	ts := w.src.GenerateTerms(n)
	w.t.Log(n.Path, "generated", len(ts), "terms:")
	for k, old := range ts {
		k, old := k, old // pin these to generate unique variables
		fn := func(data interface{}) string {
			res := old.Filter(data)
			log := res
			if log == NewLineString {
				log = "NewLine!"
			}
			w.t.Log(" ", "`"+log+"`", "from", k.String(), old.Src.String())
			return res
		}
		ts[k] = TermResult{old.Src, fn}
	}
	return ts
}

type WatchWriter struct {
	t       *testing.T
	buf     bytes.Buffer
	lines   []string
	pending string
}

func (w *WatchWriter) String() string {
	return w.buf.String()
}

func (w *WatchWriter) Lines() []string {
	l := w.lines
	if w.pending != "" {
		l = append(l, w.pending)
	}
	return l
}

func (w *WatchWriter) Write(p []byte) (int, error) {
	s := string(p)
	s = strings.Trim(s, "-")
	if cnt := len(s); cnt > 0 {
		if s == NewLineString {
			s = "NewLine!"
			w.lines, w.pending = append(w.lines, w.pending), ""
		} else if s == "|" {
			s = strings.Repeat("Tab!", cnt)
			w.pending += strings.Repeat("-", cnt)
		} else {
			w.pending += s
		}
		w.t.Log("wrote:", "'"+s+"'")
	}
	return w.buf.Write(p)
}
