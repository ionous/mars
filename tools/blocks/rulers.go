package blocks

// TermTextWhen, create a reult which produces the passed text for the passed term.
func TermTextWhen(when Term, text string, matchers ...Matcher) *Rule {
	return &Rule{"when", Matchers(matchers), TermSet{when: TermText(text)}}
}

// Token, produces the passed text if the passed target is empty.
func Token(desc, target, text string) *Rule {
	m := Matchers{
		IsTarget(target),
		IsEmpty{},
	}
	return &Rule{desc, m, TermSet{ContentTerm: TermText(text)}}
}
