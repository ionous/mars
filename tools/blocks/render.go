package blocks

import (
	"io"
	"text/tabwriter"
)

type Renderer struct {
	w             *tabwriter.Writer
	scope         Scope
	sep           Separator
	openq, closeq QuoteStack
}

func NewRenderer(w io.Writer) *Renderer {
	const minwidth, tabwidth, padding, padchar = 1, 0, 1, '-'
	out := tabwriter.NewWriter(w, minwidth, tabwidth, padding, padchar, tabwriter.Debug)
	return &Renderer{w: out,
		scope:  Scope{w: out},
		sep:    Separator{w: out},
		openq:  QuoteStack{w: out},
		closeq: QuoteStack{w: out},
	}
}

func (r *Renderer) Render(p *DocNode, rules GenerateTerms) (err error) {
	if e := r.render(p, rules); e != nil {
		err = e
	} else {
		r.closeq.flushQuote()
		r.sep.writeEnd()
		err = r.w.Flush()
	}
	return
}

const openQuote, closeQuote = "\"", "\""

func (r *Renderer) render(p *DocNode, rules GenerateTerms) (err error) {
	for _, n := range p.Children {
		terms := rules.GenerateTerms(n)
		//
		if v, e := terms.Produce(n.Data); e != nil {
			err = e
			break
		} else {
			scope, quote := v[ScopeTerm] == "true", v[QuotesTerm] == "true"

			if prefix := v[PreTerm]; len(prefix) > 0 || quote {
				r.write(prefix)
			}
			if scope {
				r.sep.separate(NewLineString)
				r.scope.changeIndent(true)
				r.flush()
				r.sep.chars = ""
			}

			if quote {
				r.openq.push(openQuote)
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
				r.sep.separate(sep)
			}

			// you could eval the sep here, and if it were terminal put it in the quotes, otherwise put it outside of the quotes -- and perhaps some other thoughtful magic.
			if quote {
				r.closeq.push(closeQuote)
			}

			if scope {
				r.scope.changeIndent(false)
				r.sep.separate("")
				r.flush()
			}
		}
	}
	return
}

func (r *Renderer) flush() {
	r.closeq.flushQuote()
	newLine := r.sep.chars == NewLineString
	if newLine {
		r.sep.flushSep()
		r.scope.writeIndent()
	} else {
		r.sep.flushSep()
	}
	r.openq.flushQuote()
}

func (r *Renderer) write(s string) {
	if len(s) > 0 {
		r.flush()
		io.WriteString(r.w, s)
	}
}
