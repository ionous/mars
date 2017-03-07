package blocks

import (
	"io"
)

//
type QuoteStack struct {
	w      io.Writer
	quotes []string
}

func (r *QuoteStack) push(s string) {
	r.quotes = append(r.quotes, s)
}

func (r *QuoteStack) flushQuote() {
	if cnt := len(r.quotes); cnt > 0 {
		var q string
		q, r.quotes = r.quotes[cnt-1], r.quotes[:cnt-1]
		io.WriteString(r.w, q)
	}
}
