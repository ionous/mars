package blocks

import (
	// "github.com/ionous/sashimi/util/errutil"
	"strings"
)

type Matcher interface {
	Matches(*DocNode) bool
	String() string
}

type Matchers []Matcher

func (ms Matchers) String() (ret string) {
	str := make([]string, len(ms))
	for i, v := range ms {
		str[i] = v.String()
	}
	return strings.Join(str, " && ")
}

// Matchers implements Matcher.
func (ms Matchers) Matches(src *DocNode) bool {
	test := true // empty is okay
	for _, m := range ms {
		if !m.Matches(src) {
			test = false
			break
		}
	}
	return test
}

type GenerateTerms interface {
	GenerateTerms(*DocNode) TermSet
}

// a rule unconditionally applies its terms if all its matches are true
// the terms have functions which produce data. an alternative implementation might be a single function which produces multiple terms based on data.
type Rule struct {
	desc    string
	matcher Matcher
	terms   TermSet
}

func (c Rule) String() string {
	return Spaces(c.desc, c.matcher.String())
}

func (r *Rule) Merge(dst TermSet) TermSet {
	for k, v := range r.terms {
		if len(dst) == 0 {
			dst = TermSet{k: TermResult{r, v.Filter}}
			// dst[k]~
		} else {
			if _, alreadySet := dst[k]; !alreadySet {
				dst[k] = TermResult{r, v.Filter}
			}
		}
	}
	return dst
}

// MarshalText, helper for debugging rule generation via json output.
func (c Rule) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

type Rules []*Rule

func (rs Rules) GenerateTerms(src *DocNode) (ret TermSet) {
	// FIX: test dst terms len for max and break early out?
	for i, cnt := 0, len(rs); i < cnt; i++ {
		if r := rs[cnt-i-1]; r.matcher.Matches(src) {
			ret = r.Merge(ret)
		}
	}
	return
}
