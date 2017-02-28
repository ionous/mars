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
				CommaSep(),
				CommaAndSep(),
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
func CommaSep() *Rule {
	return TextRule("comma sep", ",", MatcherFunc(func(src MatchSource) bool {
		return src.ApplyWhen == ApplyAfter &&
			IsElement(src) &&
			!IsThisLast(src)
	}))
}
func CommaAndSep() *Rule {
	return TextRule("comma and", ", and", MatcherFunc(func(src MatchSource) bool {
		return src.ApplyWhen == ApplyAfter &&
			IsElement(src) &&
			IsNextLast(src)
	}))
}
