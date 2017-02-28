package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
)

type English struct {
	Types inspect.Types
	rules *EnglishRules
	Generator
}

type EnglishRules struct {
	TypeRules *TypeRules
	UserRules Rules
}

func NewEnglishGenerator(types inspect.Types, words Words, cursor *DocumentCursor) *English {
	rules := &EnglishRules{
		TypeRules: &TypeRules{
			rules: Rules{
				WriteType("string", FormatString),
			},
			parsed: make(map[string]bool),
		},
	}
	return &English{types, rules, Generator{words, rules, cursor}}
}

func (en *EnglishRules) FindBestRule(src MatchSource) (ret *Rule, okay bool) {
	if r, ok := en.UserRules.FindBestRule(src); ok {
		ret, okay = r, true
	} else if r, ok := en.TypeRules.FindBestRule(src); ok {
		ret, okay = r, true
	}
	return
}

func (en *EnglishRules) AddDefaultRules() {
	// en.TypeRules["string"] = WriteType("string", func(data interface{}) (ret string, err error) {
	// 	if data == nil {
	// 		ret = "<blank>"
	// 	} else if s, ok := data.(string); ok {
	// 		ret = s
	// 	} else {
	// 		err = errutil.New("not a string")
	// 	}
	// 	return
	// })
}

func (en *English) Render(root string, data interface{}) (err error) {
	return StackRender(root, en, en.Types, data)
}

func StackRender(root string, doc DocStack, types inspect.Types, data interface{}) (err error) {
	if cmd, ok := types.TypeOf(data); !ok {
		err = errutil.New("type not found", root)
	} else {
		path := inspect.NewPath(root)
		if r, e := NewRenderer(doc, path, cmd); e != nil {
			err = e
		} else if e := inspect.Inspect(types).Visit(path, r, data); e != nil {
			err = e
		} else {
			// the visitor leaves us at the innermost last child,
			// we need to finish all terminal edges.
			err = PopStack(doc)
		}
	}
	return
}
