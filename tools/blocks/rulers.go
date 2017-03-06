package blocks

// TextRule, create a reult which produces the passed text for the passed term.
func TextRule(when Term, text string, matchers ...Matcher) *Rule {
	return &Rule{text, Matchers(matchers), TermSet{when: FixedText(text)}}
}

// Prepend, create a rule which produces prefix text for the target.
func Prepend(target, text string) *Rule {
	return TextRule(PreTerm, text, IsTarget(target))
}

// Append, create a rule which produces postfix text.
func Append(target, text string) *Rule {
	return TextRule(PostTerm, text, IsTarget(target))
}

// SepRule, create a rule which produces sep text
func SepRule(desc, sep string, matchers ...Matcher) *Rule {
	return &Rule{desc, Matchers(matchers), TermSet{SepTerm: FixedText(sep)}}
}

func TypeRule(name string, fn TermFilter) *Rule {
	desc := "format type " + name
	return &Rule{desc, IsParamType{name}, TermSet{ContentTerm: fn}}
}

// Token, produces the passed text if the passed target is empty.
func Token(target, text string) *Rule {
	desc := Spaces("token for", target, "`"+text+"`")
	m := Matchers{
		IsTarget(target),
		IsEmpty{},
	}
	return &Rule{desc, m, TermSet{ContentTerm: FixedText(text)}}
}
