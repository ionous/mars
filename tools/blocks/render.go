package blocks

import (
	"github.com/ionous/sashimi/util/lang"
	"io"
	// "strings"
	"text/tabwriter"
	"unicode"
)

type Renderer struct {
	w     *tabwriter.Writer
	scope Scope
	// sep           Separator
	openq, closeq QuoteStack
	transform     TextTransform
	lineDepth     int
	spaces        bool
	afterFirst    bool
}
type TextTransform func(string) string

func NewRenderer(w io.Writer) *Renderer {
	const minwidth, tabwidth, padding, padchar = 1, 0, 1, '-'
	out := tabwriter.NewWriter(w, minwidth, tabwidth, padding, padchar, tabwriter.Debug)
	return &Renderer{w: out,
		scope: Scope{w: out},
		// sep:    Separator{w: out},
		openq:  QuoteStack{w: out},
		closeq: QuoteStack{w: out},
	}
}

func (r *Renderer) Render(p *DocNode, rules GenerateTerms) (err error) {
	if e := r.render(p, rules); e != nil {
		err = e
	} else {
		r.closeq.flushQuote()
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
			//
			if v[TransformTerm] == "capitalize" {
				r.transform = lang.Capitalize
			}

			if prefix := v[PrefixTerm]; len(prefix) > 0 || quote {
				// hack for leading colons
				if len(prefix) == 1 && unicode.IsPunct(rune(prefix[0])) {
					r.spaces = false
				}
				r.writeWord(prefix)
			}

			// when we start a new scope, we want to start a new line.
			// this happens b/t pre and content; where sep happens fully after postfix.
			if scope {
				r.scope.changeIndent(true)
				r.lineDepth = 0
				r.spaces = false
			}

			if quote {
				r.openq.push(openQuote)
			}

			// trial: block children if we have explicit content
			if content := v[ContentTerm]; len(content) > 0 {
				r.writeWord(content)
			} else if len(n.Children) > 0 {
				if e := r.render(n, rules); e != nil {
					err = e
					break
				}
			}
			if postfix := v[PostfixTerm]; len(postfix) > 0 {
				r.writeWord(postfix)
			}

			// sep exists between things --
			// the last sep in a chain wins.
			if sep := v[SepTerm]; len(sep) > 0 {
				r.spaces = false
				if r.lineDepth > 0 {
					r.writeWord(sep)
					r.spaces = false
				}
			}

			// you could eval the sep here, and if it were terminal put it in the quotes, otherwise put it outside of the quotes -- and perhaps some other thoughtful magic.
			if quote {
				r.closeq.push(closeQuote)
			}

			if scope {
				// FIX: always write into the current scope so it can figure lineDepth and spaces
				r.scope.changeIndent(false)
				r.lineDepth = 0
				r.spaces = false
			}
		}
	}
	return
}

func (r *Renderer) writeWord(s string) {
	if cnt := len(s); cnt > 0 {
		r.forceWord(s)
	}
}

func (r *Renderer) forceWord(s string) {
	r.closeq.flushQuote()

	// fix: simplify?
	if r.spaces {
		io.WriteString(r.w, " ")
	}
	r.spaces = true

	// hack to suppress separators at the start of a scoped/line.
	if r.afterFirst && r.lineDepth == 0 {
		io.WriteString(r.w, NewLineString)
		r.lineDepth = r.scope.writeIndent() - 1
	}
	//
	r.openq.flushQuote()

	// for title-case might need transform to yield next transform
	if r.transform != nil {
		s = r.transform(s)
		r.transform = nil
	}
	io.WriteString(r.w, s)
	r.afterFirst = true
	r.lineDepth += len(s)
}
