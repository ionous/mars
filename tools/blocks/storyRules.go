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
			},
			parsed: make(map[string]Rules),
		},
		UserRules: Rules{
			TermTextWhen(QuotesTerm, "true", IsCommandField("Text", "Value")),
			TermTextWhen(QuotesTerm, "true", IsCommandField("ScriptSubject", "Subject")),
			// override previous separators when using "called something".
			TermTextWhen(SepTerm, " ", IsElement{}, IsCommand{"ScriptSubject"}),
			TermTextWhen(QuotesTerm, "false", IsParent{IsPropertyValue{"TextValue", "author"}}),
			ListFormatter("HandleEvent.Events", DataItemOrItem),
			ListFormatter("ContainsContents.Contents", DataItemAndItem),
			//
			&Rule{"exec-item",
				IsCommandOf{"Execute"},
				TermSet{
					TransformTerm: TermText("capitalize"),
					ScopeTerm:     TermText("true"),
				},
			},
			// each item generates a block, but: when our array is empty, we also want the placeholder content to generate a block: this isnt the most satisfying:
			// maybe each direct child of a scope should wind up on a new line automatically:
			// then, we simply introduce a scope on the array of execute.
			TermTextWhen(ScopeTerm, "true", IsArrayOf{"Execute"}, IsEmpty{}),
			// the default sep would be comma, and comma-and:
			TermTextWhen(SepTerm, "", IsParent{IsArrayOf{"Execute"}}, IsElement{}),
			TermTextWhen(SepTerm, ".", IsParent{IsArrayOf{"Execute"}}, IsThisLast{}),
			// generate a block for choose.false when empty to hide the empty children
			// TermTextWhen(ContentTerm, "xxxxx", IsTarget("Choose.False"), IsEmpty{}),
			&Rule{"elide false",
				Matchers{IsTarget("Choose.False"), IsEmpty{}},
				TermSet{
					// FIX: might want display false instead of this.
					PrefixTerm:  TermText(""),
					ContentTerm: TermText(""),
					PostfixTerm: TermText(""),
					ScopeTerm:   TermText("false"),
				},
			},
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
		Matchers{IsTarget(target), IsNot{IsEmpty{}}},
		TermSet{ContentTerm: TermFunction(c)},
	}
}
