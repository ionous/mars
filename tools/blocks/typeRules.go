package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"strings"
	"unicode"
)

type TypeRules struct {
	Rules  Rules
	parsed map[string]Rules
}

func (tr *TypeRules) GenerateTerms(src *DocNode) (ret TermSet) {
	// first make sure the rules for this command are parsed
	// we can rely on the stack to ensure the container rules are already parsed.
	if cmd := src.Command; cmd != nil {
		if _, ok := tr.parsed[cmd.Name]; !ok {
			rules := genCmd(cmd)
			tr.parsed[cmd.Name] = rules
			tr.Rules = append(tr.Rules, rules...)
		}
	}
	return tr.Rules.GenerateTerms(src)
}

// if we dont have a rule for this command, then parse its "mars" tags to create some.
func genCmd(cmd *inspect.CommandInfo) (ret Rules) {
	// if there are no parameters, then we simply want to print the name
	if cnt := len(cmd.Parameters); cnt == 0 {
		ret = append(ret, genSingleton(cmd.Name, cmd.Phrase))
	} else {
		if cmd.Phrase != nil {
			if terms := genPhrase(*cmd.Phrase); len(terms) > 0 {
				ret = append(ret, &Rule{"mars", IsCommand{cmd.Name}, terms})
			}
		}
		for _, p := range cmd.Parameters {
			ret = append(ret, genToken(cmd.Name, p.Name, p.Phrase))
			if p.Phrase != nil {
				if terms := genPhrase(*p.Phrase); len(terms) > 0 {
					ret = append(ret, &Rule{"mars", IsCommandField(cmd.Name, p.Name), terms})
				}
			}
		}
	}
	return
}

func genSingleton(cmd string, p *string) *Rule {
	var text string
	if p != nil {
		text = *p
	} else {
		text = PascalSpaces(cmd)
	}
	return TermTextWhen(ContentTerm, text, IsCommand{cmd})
}

// makes a token when empty
func genToken(cmd, field string, p *string) *Rule {
	var token string
	if p != nil {
		_, _, token = TokenizePhrase(*p)
	}
	if len(token) == 0 {
		token = MakeToken(PascalSpaces(field))
	}
	return TermTextWhen(ContentTerm, token, IsCommandField(cmd, field), IsEmpty{})
}

// makes a token when empty
func genPhrase(p string) TermSet {
	terms := make(TermSet)
	pre, post, _ := TokenizePhrase(p)
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
	return terms
}
