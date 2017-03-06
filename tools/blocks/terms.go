package blocks

//go:generate stringer -type=Term
type Term int

const (
	PreTerm Term = iota
	QuotesTerm
	ContentTerm
	PostTerm
	SepTerm
	ScopeTerm
	NumTerms
)

// for now, we generate string for all terms
// bools are represented by the strings "true" and "false".
type TermSet map[Term]TermFilter

type TermFilter func(interface{}) string

func FixedText(c string) TermFilter {
	return func(interface{}) string {
		return c
	}
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
	if fn, ok := ts[term]; ok {
		ret, okay = fn(data), true
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

// func (ts TermSet) ProduceDefaults(data interface{}, defaults Productions) (ret Productions, err error) {
// 	ret = make(Productions)
// 	for k := Term(0); k < NumTerms; k++ {
// 		if v, ok := ts.Filter(k, data); ok {
// 			ret[k] = v
// 		} else {
// 			ret[k] = defaults[k]
// 		}
// 	}
// 	return
// }
