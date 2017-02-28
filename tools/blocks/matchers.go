package blocks

import (
	// "log"
	"strings"
)

func ParamMatch(src MatchSource, name string) (okay bool) {
	return src.Param != nil && src.Param.Name == name
}

// container command match; but not array containers
func ContainerMatch(src MatchSource, name string) (okay bool) {
	return src.Parent != nil && src.Parent.Command != nil && src.Parent.Command.Name == name
}

func CommandMatch(src MatchSource, name string) (okay bool) {
	return src.Command != nil && src.Command.Name == name
}

func (src MatchSource) IsEmpty() (okay bool) {
	// FIX: im not convinced about cap,
	// we could do child by internal index, len here, and elsewhere?
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
	// ironic, isnt it?
	desc := target + " " + when.String() + " add text " + text
	return TextRule(desc, text, MatcherFunc(m))
}

func Prepend(target, text string) *Rule {
	return AddText(ApplyBefore, target, text)
}

func Append(target, text string) *Rule {
	return AddText(ApplyAfter, target, text)
}

func TextRule(desc, text string, m ...Matcher) *Rule {
	return &Rule{desc, m, func(n *DocNode) (string, error) {
		return text, nil
	}}
}

type FormatNode func(data interface{}) (string, error)

func FormatType(name string, fn FormatNode) *Rule {
	var m MatcherFunc = func(src MatchSource) bool {
		// FIX: we could store type, etc. expanded into the DocNode.
		return src.ApplyWhen == ApplyOn && src.Param != nil && src.Param.Type() == name
	}
	desc := "write type " + name
	return &Rule{desc, []Matcher{m},
		func(n *DocNode) (string, error) {
			return fn(n.Data)
		}}
}

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
	desc := target + " token " + text
	return TextRule(desc, text, MatcherFunc(m))
}

type MatcherFunc func(MatchSource) bool

func (f MatcherFunc) Matches(src MatchSource) bool {
	return f(src)
}
