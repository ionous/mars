package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"strings"
	"unicode"
)

type TypeRules struct {
	Rules  Rules
	parsed map[string]bool
}

func (tr *TypeRules) GenerateTerms(src *DocNode) (ret TermSet) {
	// first make sure the rules for this command are parsed
	// we can rely on the stack to ensure the container rules are already parsed.
	if cmd := src.Command; cmd != nil {
		tr.addCommand(cmd)
	}
	return tr.Rules.GenerateTerms(src)
}

// if we dont have a rule for this command, then parse its "mars" tags to create some.
func (tr *TypeRules) addCommand(cmd *inspect.CommandInfo) {
	if !tr.parsed[cmd.Name] {
		tr.parsed[cmd.Name] = true
		// if there are no parameters, then we simply want to print the name
		if cnt := len(cmd.Parameters); cnt == 0 {
			tr.addSingleton(cmd.Name, cmd.Phrase)
		} else {
			tr.addPhrase(cmd.Name, cmd.Phrase)
			for _, param := range cmd.Parameters {
				// FIX: the rules are just going to split these again.
				tr.addPhrase(cmd.Name+"."+param.Name, param.Phrase)
			}
		}
	}
}

func (tr *TypeRules) addSingleton(
	target string,
	p *string,
) {
	var text string
	if p != nil {
		text = *p
	} else {
		text = PascalSpaces(target)
	}
	tr.Rules = append(tr.Rules, TermTextWhen(ContentTerm, text, IsTarget(target)))
}

func (tr *TypeRules) addPhrase(target string, p *string) {
	var pre, post, token string
	if p != nil {
		pre, post, token = TokenizePhrase(*p)
	}

	if len(token) == 0 {
		token = MakeToken(PascalSpaces(target))
	}
	r := Token("mars", target, token)
	tr.Rules = append(tr.Rules, r)

	terms := make(TermSet)
	if len(pre) > 0 {
		terms[PreTerm] = TermText(pre)
	}
	if s := post; len(s) > 0 {
		last := strings.LastIndexFunc(s, func(r rune) bool {
			// treat closing parens as text so we get spaces before them.
			return r == ')' || !unicode.IsPunct(r)
		})

		post, sep := s[0:last+1], s[last+1:len(s)]
		if len(post) > 0 {
			terms[PostTerm] = TermText(post)
		}
		if len(sep) > 0 {
			terms[SepTerm] = TermText(sep)
		}
	}

	if len(terms) > 0 {
		r := &Rule{
			"mars",
			IsTarget(target), terms}
		tr.Rules = append(tr.Rules, r)
	}
}
