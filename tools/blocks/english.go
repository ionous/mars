package blocks

import (
	"fmt"
	"github.com/ionous/mars/tools/inspect"
	// "github.com/ionous/sashimi/util/errutil"
	"strconv"
)

type EnglishRules struct {
	TypeRules *TypeRules
	UserRules Rules
}

func NewEnglishRules(types inspect.Types) *EnglishRules {
	return &EnglishRules{
		TypeRules: &TypeRules{
			Rules: Rules{
				// fix: move matchers into a separate package?
				ValueFormatter(),
				CommaSep(),
				CommaAndSep(),
				Quote("Text.Value"),
				Quote("ScriptSubject.Subject"),
				SpaceSep("ScriptSubject"),
				TextRule(QuotesTerm, "false", IsParent{IsPropertyValue{"TextValue", "author"}}),
			},
			parsed: make(map[string]bool),
		},
	}
}

func (en *EnglishRules) GenerateTerms(src *DocNode) TermSet {
	return Merge(
		en.UserRules.GenerateTerms(src),
		en.TypeRules.GenerateTerms(src),
	)
}

func CommaSep() *Rule {
	// have to exlude the last element for one element lists.
	return SepRule("element-comma-sep", ", ", IsElement{}, IsNot{IsThisLast{}})
}

func CommaAndSep() *Rule {
	return SepRule("element-comma-and", ", and ", IsElement{}, IsNextLast())
}

func SpaceSep(target string) *Rule {
	return SepRule("space after "+target, " ", IsElement{}, IsCommand{target})
}

func Quote(target string) *Rule {
	return TextRule(QuotesTerm, "true", IsTarget(target))
}

func Scope(target string) *Rule {
	return TextRule(ScopeTerm, "true", IsTarget(target))
}

// ValueFormatter generates Content for non-empty values
func ValueFormatter() *Rule {
	return &Rule{"global-value-formatter",
		Matchers{IsValue{}, IsNot{IsEmpty{}}},
		TermSet{ContentTerm: func(data interface{}) (ret string) {
			// FIX: arrays of these???
			switch val := data.(type) {
			case string:
				ret = val
			case float64:
				ret = strconv.FormatFloat(val, 'g', -1, 64)
			case bool:
				ret = strconv.FormatBool(val)
			default:
				ret = fmt.Sprint(val)
			}
			return
		}},
	}
}
