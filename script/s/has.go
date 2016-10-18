package s

import (
	"github.com/ionous/mars/script/frag"
)

// Has sets the value of an instance property.
// The property must (eventually) be declared for the class.
// ( For example, via Have. )
func Has(property string, values ...interface{}) (ret frag.Fragment) {
	switch len(values) {
	case 0:
		ret = frag.Select{Choices: []string{property}}
	case 1:
		ret = frag.SetKeyValue{property, values[0]}
	default:
		// used for table, list definitions
		// FIX: tables should be reworked
		// lists should use something more like the rt section uses
		// for example: HasList{} -- dont be afraid to be specific,
		ret = frag.SetKeyValue{property, values}
	}
	return ret
}
