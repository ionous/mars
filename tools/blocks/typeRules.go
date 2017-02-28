package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

type TypeRules struct {
	Rules  Rules
	parsed map[string]bool
}

func (tr *TypeRules) FindBestRule(src MatchSource) (ret *Rule, okay bool) {
	// first make sure the rules for this command are parsed
	// we can rely on the stack to ensure the container rules are already parsed.
	if cmd := src.Command; cmd != nil {
		tr.addCommand(cmd)
	}
	if r, ok := tr.Rules.FindBestRule(src); ok {
		ret, okay = r, true
	}
	return
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

func (tr *TypeRules) addSingleton(target string, p *string) {
	var text string
	if p != nil {
		text = *p
	} else {
		text = PascalSpaces(target)
	}
	tr.Rules = append(tr.Rules, Prepend(target, text))
}

func (tr *TypeRules) addPhrase(target string, p *string) {
	var pre, post, token string
	if p != nil {
		pre, post, token = TokenizePhrase(*p)
	}
	if len(token) == 0 {
		token = MakeToken(PascalSpaces(target))
	}
	if len(pre) > 0 {
		tr.Rules = append(tr.Rules, Prepend(target, pre))
	}
	if len(token) > 0 {
		tr.Rules = append(tr.Rules, Token(target, token))
	}
	if len(post) > 0 {
		tr.Rules = append(tr.Rules, Append(target, post))
	}
}
