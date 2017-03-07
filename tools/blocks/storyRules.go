package blocks

import (
	"github.com/ionous/mars/tools/inspect"
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
				// TermTextWhen(ScopeTerm, "true", IsArrayOf{"Execute"}),
				// &Rule{"exec-block",
				// 	IsArrayOf{"Execute"},
				// 	TermSet{
				// 		ScopeTerm: TermText("true"),
				// 		// SepTerm:   TermText(NewLineString)},
				// },
				//
				ListFormatter("HandleEvent.Events", DataItemOrItem),
				//
				&Rule{"exec-item",
					IsCommandOf{"Execute"},
					TermSet{
						TransformTerm: TermText("capitalize"),
						// ScopeTerm:     TermText(NewLineString)},
						ScopeTerm: TermText("true")},
				},
				TermTextWhen(SepTerm, ".", IsParent{IsArrayOf{"Execute"}}, IsThisLast{}),
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

func ListFormatter(target string, c TermFilter) *Rule {
	return &Rule{"value",
		IsTarget(target),
		TermSet{ContentTerm: TermFunction(c)},
	}
}
