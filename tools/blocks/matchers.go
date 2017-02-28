package blocks

import (
	// "github.com/ionous/mars/tools/inspect"
	"strings"
)

func ParamMatch(src MatchSource, name string) (okay bool) {
	return src.Param != nil && src.Param.Name == name
}

func ContainerMatch(src MatchSource, name string) (okay bool) {
	return src.Parent != nil && src.Parent.Command.Name == name
}

func CommandMatch(src MatchSource, name string) (okay bool) {
	return src.Command != nil && src.Command.Name == name
}

func (src MatchSource) IsEmpty() (okay bool) {
	// switch src.Type {
	// case CommandNode:
	// 	okay = src.Command == nil
	// case ArrayNode:
	// 	okay = cap(src.Children) == 0
	// case ValueNode:
	// 	okay = src.Data == nil
	// }
	// FIX: im not convinced about cap,
	// we could do child by internal index
	// and length here, and elsewhere?
	return cap(src.Children) == 0 && src.Data == nil
}

// matches []Matcher:Matches(MatchSource) bool
// 	what    func(*DocNode) string
func AddText(when ApplyWhen, target, text string) *Rule {
	var m MatcherFunc
	if tp := strings.SplitN(target, ".", 2); len(tp) > 1 {
		container, field := tp[0], tp[1]
		m = func(src MatchSource) bool {
			return src.ApplyWhen == when &&
				ContainerMatch(src, container) &&
				ParamMatch(src, field)
		}
	} else {
		typeName := tp[0]
		m = func(src MatchSource) bool {
			return src.ApplyWhen == when &&
				CommandMatch(src, typeName)
		}
	}
	return TextRule(text, MatcherFunc(m))
}

func Prepend(target, text string) *Rule {
	return AddText(ApplyBefore, target, text)
}

func Append(target, text string) *Rule {
	return AddText(ApplyAfter, target, text)
}

func TextRule(text string, m ...Matcher) *Rule {
	return &Rule{m, func(n *DocNode) (string, error) {
		return text, nil
	}}
}

type FormatNode func(data interface{}) (string, error)

func WriteType(name string, fn FormatNode) *Rule {
	var m MatcherFunc = func(src MatchSource) bool {
		// FIX: we could store type, etc. expanded into the DocNode.
		return src.ApplyWhen == ApplyOn && src.Parent != nil && src.Param.Type() == name
	}
	return &Rule{[]Matcher{m},
		func(n *DocNode) (string, error) {
			return fn(n.Data)
		}}
}

// IM NOT GETTING A MATCH FOR THE DIRECTIVE ITSELF
func Token(target, text string) *Rule {
	// FIX: possibly type scoping?
	var typeName, fieldName string
	tp := strings.SplitN(target, ".", 2)
	if len(tp) > 0 {
		typeName = tp[0]
	}
	if len(tp) > 1 {
		fieldName = tp[1]
	}

	// FIX: possibly rule stacking?
	m := func(src MatchSource) bool {
		return src.ApplyWhen == ApplyOn &&
			ContainerMatch(src, typeName) &&
			ParamMatch(src, fieldName) &&
			src.IsEmpty()
	}
	return TextRule(text, MatcherFunc(m))
}

type MatcherFunc func(MatchSource) bool

func (f MatcherFunc) Matches(src MatchSource) bool {
	return f(src)
}
