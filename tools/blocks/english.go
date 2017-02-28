package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

type EnglishRules struct {
	TypeRules *TypeRules
	UserRules Rules
}

func NewEnglishRules(types inspect.Types) *EnglishRules {
	return &EnglishRules{
		TypeRules: &TypeRules{
			Rules: Rules{
				FormatType("string", FormatString),
			},
			parsed: make(map[string]bool),
		},
	}
}

func (en *EnglishRules) FindBestRule(src MatchSource) (ret *Rule, okay bool) {
	if r, ok := en.UserRules.FindBestRule(src); ok {
		ret, okay = r, true
	} else if r, ok := en.TypeRules.FindBestRule(src); ok {
		ret, okay = r, true
	}
	return
}
