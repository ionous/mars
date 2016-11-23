package script

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	S "github.com/ionous/sashimi/source"
	"github.com/ionous/sashimi/source/types"
)

// The targets a noun for new assertions.
func The(target string, fragments ...backend.Fragment) backend.Spec {
	return internal.NounPhrase{
		types.Subject(target), fragments,
	}
}

// Understand builds statements for parsing player input.
func Understand(s string) internal.ParserPartial {
	return internal.ParserPartial([]string{s})
}

// Our is an alias for The
var Our = The

// Called asserts the existence of a class or instance.
// For example, The("room", Called("home"))
func Called(subject string) internal.ScriptSubject {
	return internal.SetSubject(subject)
}

// Exists is an optional fragment for making otherwise empty statements more readable.
// For example, The("room", Called("parlor of despair"), Exists())
func Exists() backend.Fragment {
	return internal.Exists{}
}

// Exist is an alias of Exists used for classes.
// For example, The("kinds", Called("rooms"), Exist())
var Exist = Exists

// AreOneOf defines a enumerated choices for all instances of the class.
func AreOneOf(name string, rest ...string) internal.ClassEnum {
	names := append([]string{name}, rest...)
	return internal.ClassEnum{Choices: names}
}

// AreEither defines one of two states for all instances of the class.
// ex. AreEither("this").Or("that")
func AreEither(firstChoice string) internal.EitherChoice {
	return internal.EitherChoice(firstChoice)
}

// Is asserts one or more states of one or more enumerations.
// The enumerations must (eventually) be declared for the target's class. ( For example, via AreEither, or AreOneOf, )
func Is(choice string, choices ...string) internal.Choices {
	return internal.HasChoices(append(choices, choice)...)
}

// IsKnownAs declares an alias for the current subject.
// ex. The("cabinet", IsKnownAs("armoire").And("..."))
func IsKnownAs(name string, names ...string) internal.KnownAs {
	return internal.KnownAs{append(names, name)}
}

// Have asserts the existance of a property for all instances of a given class.
// For relations, use HaveOne or HaveMany.
func Have(name string, class string) backend.Fragment {
	return internal.NewClassProperty(name, class)
}

// HaveOne establishes a one-to-one, or one-to-many relation.
func HaveOne(name string, class string) internal.ClassRelation {
	return internal.ClassRelation{Src: internal.ClassRelative{name, class, S.RelativeOne}}
}

// HaveMany establishs a many-to-one relation.
func HaveMany(name string, class string) internal.ClassRelation {
	return internal.ClassRelation{Src: internal.ClassRelative{name, class, S.RelativeMany}}
}

// Has asserts the value of an object's property. The property must (eventually) be declared for the class. ( For example, via Have. )
func Has(property string, values ...interface{}) (ret backend.Fragment) {
	// we let the compiler checks ( and marshals ) types via a property "Builder".
	// (ex. enumBuilder, numBuilder, textBuilder, pointerBuilder, and relativeBuilder.)
	// this is necessary because we use string for both text, obj, and relation values.
	switch len(values) {
	case 0:
		ret = internal.HasChoices(property)
	case 1:
		ret = internal.HasPropertyValue(property, values[0])
	default:
		// used for table, list definitions
		// MARS: should tables be reworked? even lists should probably use something more like the rt section uses
		// for example: HasList{} -- dont be afraid to be specific,
		ret = internal.HasPropertyValue(property, values)
	}
	return ret
}

// Can asserts a new verb for the targeted noun.
func Can(verb types.ActionName) internal.CanDoPhrase {
	return internal.NewCanDo(verb)
}

// To assigns runtime statements to a default action handler.
func To(action string, call rt.Execute, calls ...rt.Execute) backend.Fragment {
	return internal.NewDefaultAction(action, internal.JoinCallbacks(call, calls))
}

// Before actions are implemented as capturing event listeners which allow them to run prior to the default actions of the passed event.
func Before(event string) internal.EventPartial {
	return internal.NewEvent(S.ListenTargetOnly|S.ListenCapture, event)
}

// After actions are queued to run after the default actions for the passed event have completed.
func After(event string) internal.EventPartial {
	return internal.NewEvent(S.ListenTargetOnly|S.ListenCapture|S.ListenRunAfter, event)
}

// When actions are implemented as bubbling event handlers. This allows them to run sandwiched between the "before actions" and the default actions of the passed event.
func When(event string) internal.EventPartial {
	return internal.NewEvent(S.ListenTargetOnly, event)
}
