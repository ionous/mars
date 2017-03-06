package blocks

import (
	"strconv"
	"strings"
)

// IsNot: reverses the results of its matcher.
type IsNot struct {
	Negate Matcher
}

func (m IsNot) String() string {
	return Spaces("IsNot", m.Negate.String())
}

func (m IsNot) Matches(src *DocNode) (okay bool) {
	return !m.Negate.Matches(src)
}

// IsParent: asks matching questions of a node's parent.
type IsParent struct {
	Parent Matcher
}

func (m IsParent) String() string {
	return Spaces("is parent", m.Parent.String())
}

func (m IsParent) Matches(src *DocNode) (okay bool) {
	return src.Parent != nil && m.Parent.Matches(src.Parent)
}

// IsCommand: matches a command of the exact named type.
type IsCommand struct {
	Command string
}

func (m IsCommand) String() string {
	return Spaces("IsCommand", m.Command)
}

func (m IsCommand) Matches(src *DocNode) bool {
	return src.Command != nil && src.Command.Name == m.Command
}

// IsField,
type IsField struct {
	Field string
}

func (m IsField) String() string {
	return Spaces("IsField", m.Field)
}

func (m IsField) Matches(src *DocNode) bool {
	return src.Param != nil && src.Param.Name == m.Field
}

// IsCommandField, matches fields of the passed name inside commands of the passed name.
func IsCommandField(command, field string) Matcher {
	return Matchers{IsParent{IsCommand{command}}, IsField{field}}
}

// IsTarget creates a command/field matcher for targets of the form "Command.Field"
func IsTarget(target string) (ret Matcher) {
	if tp := strings.SplitN(target, ".", 2); len(tp) > 1 {
		ret = IsCommandField(tp[0], tp[1])
	} else {
		ret = IsCommand{tp[0]}
	}
	return
}

// IsImplementor: matches a command implementing the named type.
type IsImplementor struct {
	BaseType string
}

func (m IsImplementor) String() string {
	return Spaces("IsImplementor", m.BaseType)
}

func (m IsImplementor) Matches(src *DocNode) bool {
	// we implement commands, and our base type is the passed name.
	return src.Command != nil && src.BaseType != nil && src.BaseType.Name == m.BaseType
}

// IsElement: matches any array element.
type IsElement struct{}

func (_ IsElement) String() string {
	return "IsElement"
}

func (_ IsElement) Matches(src *DocNode) bool {
	return src.Parent != nil && src.Parent.Command == nil
}

// IsValue: matches any primitive value.
type IsValue struct{}

func (_ IsValue) String() string {
	return "IsValue"
}

func (_ IsValue) Matches(src *DocNode) bool {
	return cap(src.Children) == 0
}

// IsThisLast: matches the nth to last element of an array;
// when zero, the default value, it matches the last element.
type IsThisLast struct {
	TerminalDist int
}

func (m IsThisLast) String() string {
	return Spaces("IsThisLast", strconv.Itoa(m.TerminalDist))
}

func (m IsThisLast) Matches(src *DocNode) (okay bool) {
	if src.Parent != nil {
		if cnt := src.Parent.NumChildren(); cnt > m.TerminalDist {
			test := src.Parent.Children[cnt-m.TerminalDist-1]
			okay = test == src
		}
	}
	return
}

// IsNextLast: matches the second to last element of an array.
func IsNextLast() Matcher {
	return IsThisLast{1}
}

// IsEmpty: matches any nil command, array, or primitive value.
type IsEmpty struct{}

func (_ IsEmpty) String() string {
	return "IsEmpty"
}

func (_ IsEmpty) Matches(src *DocNode) (okay bool) {
	// FIX: im not convinced about cap,
	// we could do child by internal index, len here, and elsewhere?
	return cap(src.Children) == 0 && src.Data == nil
}

// IsParamType: matches command fields of the named primitive type.
type IsParamType struct {
	Name string
}

func (m IsParamType) String() string {
	return Spaces("IsParamType", m.Name)
}

func (m IsParamType) Matches(src *DocNode) (okay bool) {
	// FIX: we could store param type, etc. expanded into the DocNode.
	return src.Param != nil && src.Param.Type() == m.Name
}

// IsPropertyValue matches declarative user properties in mars
// ex. author, or headline.
type IsPropertyValue struct {
	Type     string
	Property string
}

func (m IsPropertyValue) String() string {
	return Spaces("is property value", m.Type, m.Property)
}

func (m IsPropertyValue) Matches(src *DocNode) (okay bool) {
	if (IsParent{IsCommand{m.Type}}).Matches(src) {
		prop := src.Parent.Children[0]
		couldBe := prop != src && prop.Data != nil
		if couldBe {
			if s, ok := prop.Data.(string); ok && s == m.Property {
				okay = true
			}
		}
	}
	return
}
