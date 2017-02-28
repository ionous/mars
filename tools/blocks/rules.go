package blocks

//go:generate stringer -type=ApplyWhen
type ApplyWhen int

const (
	ApplyBefore ApplyWhen = iota
	ApplyOn
	ApplyAfter
)

type MatchSource struct {
	*DocNode
	ApplyWhen
}

// func NewRule(desc string, fn func(*DocNode) (string, error), ms ...Matcher) *Rule {
// 	return &Rule{desc, ms, fn}
// }

type Matcher interface {
	Matches(MatchSource) bool
}

type RuleFinder interface {
	FindBestRule(MatchSource) (*Rule, bool)
}

type Rule struct {
	desc    string
	matches []Matcher
	what    FilterNode
}

func (c Rule) String() string {
	return c.desc
}

func (c Rule) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

type FilterNode func(*DocNode) (string, error)

func (r *Rule) Write(out Words, node *DocNode) (err error) {
	if s, e := r.what(node); e != nil {
		err = e
	} else {
		out.WriteWord(s)
	}
	return
}

func (r *Rule) AllMatch(src MatchSource) (okay bool) {
	if len(r.matches) > 0 {
		failed := false
		for _, m := range r.matches {
			if !m.Matches(src) {
				failed = true
				break
			}
		}
		okay = !failed
	}
	return
}

type Rules []*Rule

func (rs Rules) FindBestRule(src MatchSource) (ret *Rule, okay bool) {
	for i, cnt := 0, len(rs); i < cnt; i++ {
		if r := rs[cnt-i-1]; r.AllMatch(src) {
			ret, okay = r, true
			break
		}
	}
	return
}
