package s

import (
	"github.com/ionous/mars/script/frag"
)

// Has sets the value of an instance property.
// The property must (eventually) be declared for the class. ( For example, via Have. )
//
// NOTE: its the compiler checks ( and marshals ) types.
// this is necessary because we dont know what's valid yet:
// string is used, for instance, for both text, ref, and relation.
//
// under the current compiler, each property type has its own "Builder". they are:
// enumBuilder, numBuilder, textBuilder, pointerBuilder, and relativeBuilder.
func Has(property string, values ...interface{}) (ret frag.Fragment) {
	switch len(values) {
	case 0:
		ret = frag.Select{Choices: []string{property}}
	case 1:
		ret = frag.SetKeyValue{property, values[0]}
	default:
		// used for table, list definitions
		// FIX: tables should be reworked
		// lists should probably use something more like the rt section uses
		// for example: HasList{} -- dont be afraid to be specific,
		ret = frag.SetKeyValue{property, values}
	}
	return ret
}

// return one of the valid model storage types:
// NumEval, TextEval, RefEval,...
// func marshal(value interface{}) (ret interface{}) {
// 	switch val := value.(type) {
// 	case Execute, BoolEval, NumEval, TextEval, RefEval, NumListEval, TextListEval, RefListEval:
// 		ret = value
// 	case bool:
// 		ret = rt.Bool(val)
// 	case int:
// 		ret = rt.I(val)
// 	case float64:
// 		ret = rt.N(val)
// 	case string:
// 		ret = rt.T(val)
// 	case []string:
// 		ret= rt.Ts(val)
// 	case []
// 	// case int:
// 	// 	sayWhat = append(sayWhat, PrintNum{I(val)})
// 	// case rt.NumEval:
// 	// 	sayWhat = append(sayWhat, PrintNum{val})
// 	// case string:
// 	// 	sayWhat = append(sayWhat, PrintText{T(val)})
// 	// case rt.TextEval:
// 	// 	sayWhat = append(sayWhat, PrintText{val})
// 	// case rt.Execute:
// 	// 	// FIX: could buffer operations have a specialized interface implementation?
// 	// 	sayWhat = append(sayWhat, val)
// 	default:
// 		panic("say what?")
// 	}
// }
