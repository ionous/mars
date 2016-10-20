package rt

import (
	"github.com/ionous/sashimi/util/ident"
	"strconv"
)

// Bool provides a dl boolean primitive.
type Bool bool

// GetBool implements BoolEval; providing the dl with a literal boolean type.
func (b Bool) GetBool(Runtime) (Bool, error) {
	return b, nil
}

// String returns a nicely formatted float, with no decimal point when possible.
func (b Bool) String() (ret string) {
	if b {
		ret = "true"
	} else {
		ret = "false"
	}
	return
}

// Number provides a dl float primitive.
type Number float64

// GetNumber implements NumEval providing the dl with a literal number type.
func (n Number) GetNumber(Runtime) (Number, error) {
	return n, nil
}

// Int converts to native int.
func (n Number) Int() int {
	return int(n)
}

// Float converts to native float.
func (n Number) Float() float64 {
	return float64(n)
}

// String returns a nicely formatted float, with no decimal point when possible.
func (n Number) String() string {
	return strconv.FormatFloat(n.Float(), 'g', -1, 64)
}

// Text provides a dl string primitive.
type Text string

// GetNumber implements TextEval providing the dl with a literal text type.
func (s Text) GetText(Runtime) (Text, error) {
	return s, nil
}

func (s Text) String() string {
	return string(s)
}

// Reference provides an object pointer primitive.
type Reference ident.Id

// GetReference implements RefEval allowing the NullRef, R, The to be used as literals
func (xr Reference) GetReference(Runtime) (Reference, error) {
	return xr, nil
}

func (xr Reference) Id() ident.Id {
	return ident.Id(xr)
}

func (xr Reference) String() string {
	return ident.Id(xr).String()
}
