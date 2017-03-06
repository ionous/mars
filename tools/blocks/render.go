package blocks

import (
	"io"
	"strings"
	"unicode"
)

type Renderer struct {
	w                      io.Writer
	scopeCount, quoteCount int
	lastSep                string
	openQuote, closeQuote  QuoteStack
}

func NewRenderer(w io.Writer) *Renderer {
	return &Renderer{w: w}
}

func (r *Renderer) Render(p *DocNode, rules GenerateTerms) (err error) {
	if e := r.render(p, rules); e != nil {
		err = e
	} else {
		r.closeQuote.flush(r.w)
		r.lastSep = strings.TrimRightFunc(r.lastSep, unicode.IsSpace)
		r.flushSep("")
	}
	return
}

var openQuote, closeQuote = "'", "'"

func (r *Renderer) render(p *DocNode, rules GenerateTerms) (err error) {
	for _, n := range p.Children {
		terms := rules.GenerateTerms(n)
		//
		if v, e := terms.Produce(n.Data); e != nil {
			err = e
			break
		} else {
			scope, quote := v[ScopeTerm] == "true", v[QuotesTerm] == "true"
			if scope {
				panic("not implemented")
				r.scope(true)
			}

			if prefix := v[PreTerm]; len(prefix) > 0 || quote {
				r.write(prefix)
			}

			if quote {
				r.openQuote.push(openQuote)
			}

			// trial: block children if we have content
			if content := v[ContentTerm]; len(content) > 0 {
				r.write(content)
			} else if len(n.Children) > 0 {
				if e := r.render(n, rules); e != nil {
					err = e
					break
				}
			}
			if postfix := v[PostTerm]; len(postfix) > 0 {
				r.write(postfix)
			}

			// sep exists between things --
			// the last sep in a chain wins.
			if sep, ok := v[SepTerm]; ok {
				r.lastSep = sep
			}

			// you could eval the sep here, and if it were terminal put it in the quotes, otherwise put it outside of the quotes -- and perhaps some other thoughtful magic.
			if quote {
				r.closeQuote.push(closeQuote)
			}

			if scope {
				r.scope(false)
			}
		}
	}
	return
}

func (r *Renderer) scope(start bool) {
	if start {
		r.scopeCount++
	} else {
		r.scopeCount--
	}
}

type QuoteStack struct {
	quotes []string
}

func (r *QuoteStack) push(s string) {
	r.quotes = append(r.quotes, s)
}

func (r *QuoteStack) flush(w io.Writer) {
	if cnt := len(r.quotes); cnt > 0 {
		var q string
		q, r.quotes = r.quotes[cnt-1], r.quotes[:cnt-1]
		io.WriteString(w, q)
	}
}

func (r *Renderer) flushSep(next string) {
	// we treat spaces in sep as soft-spaces --
	// collapsing our default separator down down down
	if len(r.lastSep) > 0 {
		io.WriteString(r.w, r.lastSep)
	}
	r.lastSep = next
}

func (r *Renderer) write(s string) {
	if len(s) > 0 {
		r.closeQuote.flush(r.w)
		r.flushSep(" ")
		r.openQuote.flush(r.w)
		io.WriteString(r.w, s)
	}
}

// var _Defaults = Productions{
// 	SepTerm: " ",
// }
