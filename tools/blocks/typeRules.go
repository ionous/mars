package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

type TypeRules struct {
	rules  Rules
	parsed map[string]bool
}

func (tr *TypeRules) FindBestRule(src MatchSource) (ret *Rule, okay bool) {
	// first make sure the rules for this command are parsed
	// we can rely on the stack to ensure the container rules are already parsed.
	if cmd := src.Command; cmd != nil {
		tr.addCommand(cmd)
	}
	if r, ok := tr.rules.FindBestRule(src); ok {
		ret, okay = r, true
	}
	return
}

// if we dont have this rule, add it.
func (tr *TypeRules) addCommand(cmd *inspect.CommandInfo) {
	if !tr.parsed[cmd.Name] {
		tr.parsed[cmd.Name] = true
		if p := cmd.Phrase; p != nil {
			tr.addPhrase(cmd.Name, *p)
		}
		//
		for _, param := range cmd.Parameters {
			if p := param.Phrase; p != nil {
				// FIX: the rules are just going to split these again.
				tr.addPhrase(cmd.Name+"."+param.Name, *p)
			}
		}
	}
}

func (tr *TypeRules) addPhrase(target string, p string) {
	// FIX: remove deref
	pre, post, token := TokenizePhrase(p)
	if len(token) == 0 {
		token = MakeToken(PascalSpaces(target))
	}
	if len(pre) > 0 {
		tr.rules = append(tr.rules, Prepend(target, pre))
	}
	if len(token) > 0 {
		tr.rules = append(tr.rules, Token(target, token))
	}
	if len(post) > 0 {
		tr.rules = append(tr.rules, Append(target, post))
	}
}
