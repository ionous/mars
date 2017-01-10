package rt

import (
	"github.com/ionous/sashimi/util/ident"
	"strconv"
)

// Bool provides a dl boolean primitive.
type Bool struct {
	Value bool
}

// GetBool implements BoolEval; providing the dl with a literal boolean type.
func (b Bool) GetBool(Runtime) (Bool, error) {
	return b, nil
}

// String uses strconv.FormatBool.
func (b Bool) String() string {
	return strconv.FormatBool(b.Value)
}

// Number provides a dl float primitive.
type Number struct {
	Value float64
}

// GetNumber implements NumberEval providing the dl with a literal number type.
func (n Number) GetNumber(Runtime) (Number, error) {
	return n, nil
}

// Int converts to native int.
func (n Number) Int() int {
	return int(n.Value)
}

// Float converts to native float.
func (n Number) Float() float64 {
	return n.Value
}

// String returns a nicely formatted float, with no decimal point when possible.
func (n Number) String() string {
	return strconv.FormatFloat(n.Value, 'g', -1, 64)
}

// Text provides a dl string primitive.
type Text struct {
	Value string `mars:";text"`
}

// GetText implements interface TextEval.
func (s Text) GetText(Runtime) (Text, error) {
	return s, nil
}

// String returns the text.
func (s Text) String() string {
	return s.Value
}

// State provides a dl enumerated value primitive.
type State struct {
	Value string
}

// GetState implements StateEval; providing the dl with a literal enum type.
func (s State) GetState(Runtime) (State, error) {
	return s, nil
}

func (s State) GetId() ident.Id {
	// MARS: a non-serializing cached value?
	return ident.MakeId(s.Value)
}

// String returns the underlying ident.Id string.
func (s State) String() string {
	return s.Value
}
