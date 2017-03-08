package blocks

//go:generate stringer -type=Term
type Term int

const (
	PrefixTerm Term = iota
	QuotesTerm
	ContentTerm
	PostfixTerm
	SepTerm
	ScopeTerm
	TransformTerm
	NumTerms
)

// for now, we generate string for all terms
// bools are represented by the strings "true" and "false".
type TermResult struct {
	Src    *Rule
	Filter TermFilter
}

type TermSet map[Term]TermResult

type TermFilter func(interface{}) string

func TermText(c string) TermResult {
	return TermFunction(func(interface{}) string {
		return c
	})
}
func TermFunction(c TermFilter) TermResult {
	return TermResult{nil, c}
}

// Merge, things in dst take precedence.
func Merge(dst, src TermSet) TermSet {
	for k, v := range src {
		if len(dst) == 0 {
			dst = TermSet{k: v}
			// dst[k]~
		} else {
			if _, alreadySet := dst[k]; !alreadySet {
				dst[k] = v
			}
		}
	}
	return dst
}

func (ts TermSet) Terms() (ret []Term) {
	for k, _ := range ts {
		ret = append(ret, k)
	}
	return
}

func (ts TermSet) String() (ret string) {
	str := make([]string, 0, len(ts))
	for k, _ := range ts {
		str = append(str, k.String())
	}
	return Spaces(str...)
}

func (ts TermSet) Filter(term Term, data interface{}) (ret string, okay bool) {
	if r, ok := ts[term]; ok {
		ret, okay = r.Filter(data), true
	}
	return
}

type Productions map[Term]string

func (ts TermSet) Produce(data interface{}) (ret Productions, err error) {
	ret = make(Productions)
	for k := Term(0); k < NumTerms; k++ {
		if v, ok := ts.Filter(k, data); ok {
			ret[k] = v
		}
	}
	return
}
