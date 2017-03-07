package blocks

import (
	"fmt"
	"github.com/ionous/mars/tools/inspect"
	"strconv"
)

type StoryRules struct {
	TypeRules *TypeRules
	UserRules Rules
}

func NewStoryRules(types inspect.Types) *StoryRules {
	return &StoryRules{
		TypeRules: &TypeRules{
			Rules: Rules{
				// fix: move matchers into a separate package?
				ValueFormatter(),
				// note: exludes the last element for one element lists.
				TermTextWhen(SepTerm, ", ", IsElement{}, IsNot{IsThisLast{}}),
				TermTextWhen(SepTerm, ", and ", IsElement{}, IsNextLast()),
				TermTextWhen(QuotesTerm, "true", IsCommandField("Text", "Value")),
				TermTextWhen(QuotesTerm, "true", IsCommandField("ScriptSubject", "Subject")),
				// override previous separators when using "called something".
				TermTextWhen(SepTerm, " ", IsElement{}, IsCommand{"ScriptSubject"}),
				TermTextWhen(QuotesTerm, "false", IsParent{IsPropertyValue{"TextValue", "author"}}),
				TermTextWhen(ScopeTerm, "true", IsArrayOf{"Execute"}),
				//
				TermTextWhen(SepTerm, NewLineString, IsElement{}, IsCommandOf{"Execute"}),
			},
			parsed: make(map[string]bool),
		},
	}
}

func (en *StoryRules) GenerateTerms(src *DocNode) TermSet {
	return Merge(
		en.UserRules.GenerateTerms(src),
		en.TypeRules.GenerateTerms(src),
	)
}

// ValueFormatter generates Content for non-empty values
func ValueFormatter() *Rule {
	return &Rule{"value",
		Matchers{IsValue{}, IsNot{IsEmpty{}}},
		TermSet{ContentTerm: TermFunction(DataToString)},
	}
}

func DataToString(data interface{}) (ret string) {
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
}
