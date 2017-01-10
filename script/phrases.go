package script

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	"github.com/ionous/sashimi/source/types"
)

// The targets a noun for new assertions.
func The(target string, fragments ...backend.Fragment) backend.Declaration {
	return internal.NounPhrase{target, fragments}
}

// Understand builds statements for parsing player input.
func Understand(s string) internal.ParserPartial {
	return internal.ParserPartial{s}
}

// Our is an alias for The
var Our = The

// Called asserts the existence of a class or instance.
// For example, The("room", Called("home"))
func Called(subject string) internal.ScriptSubject {
	return internal.ScriptSubject{subject}
}

func HasSingularName(subject types.NamedClass) internal.ScriptSingular {
	return internal.ScriptSingular{subject}
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
	return internal.Choices{append(choices, choice)}
}

// IsKnownAs declares an alias for the current subject.
// ex. The("cabinet", IsKnownAs("armoire").And("..."))
func IsKnownAs(name string, names ...string) internal.KnownAs {
	return internal.KnownAs{append(names, name)}
}

// Have asserts the existance of a property for all instances of a given class.
// For relations, use HaveOne or HaveMany.
func Have(name string, class types.NamedClass) backend.Fragment {
	return internal.ClassProperty{types.NamedProperty(name), class}
}

// HaveOne establishes a one-to-one, or one-to-many relation.
func HaveOne(name string, class types.NamedClass) internal.PartialRelation {
	return internal.NewHaveOne(types.NamedProperty(name), class)
}

// HaveMany establishs a many-to-one relation.
func HaveMany(name string, class types.NamedClass) internal.PartialRelation {
	return internal.NewHaveMany(types.NamedProperty(name), class)
}

func HasNumber(property string, value rt.NumberEval) (ret backend.Fragment) {
	return internal.NumberValue{property, value}
}

func HasText(property string, value rt.TextEval) (ret backend.Fragment) {
	return internal.TextValue{property, value}
}

// fix? ref as string because property builder
func HasRef(property string, value string) (ret backend.Fragment) {
	return internal.RefValue{property, value}
}

// fix? there is no ref list, only a series of single refs
// func HasRefs(property string, values ...string) (ret backend.Fragment) {
// 	return internal.RefValues{property, values}
// }

func HasNumberList(property string, value rt.NumListEval) (ret backend.Fragment) {
	return internal.NumberValues{property, value}
}

func HasTextList(property string, value rt.TextListEval) (ret backend.Fragment) {
	return internal.TextValues{property, value}
}

// Can asserts a new verb for the targeted noun.
func Can(verb types.NamedAction) internal.CanDoPhrase {
	return internal.CanDoPhrase{ActionName: verb}
}

// To assigns runtime statements to a default action handler.
func To(action types.NamedAction, call rt.Execute, calls ...rt.Execute) backend.Fragment {
	return internal.DefaultAction{action, internal.JoinCallbacks(call, calls)}
}

// Before actions are implemented as capturing event listeners which allow them to run prior to the default actions of the passed event.
func Before(event string) internal.EventPartial {
	return internal.NewEvent(event, internal.BeforeEvent{})
}

// After actions are queued to run after the default actions for the passed event have completed.
func After(event string) internal.EventPartial {
	return internal.NewEvent(event, internal.AfterEvent{})
}

// When actions are implemented as bubbling event handlers. This allows them to run sandwiched between the "before actions" and the default actions of the passed event.
func When(event string) internal.EventPartial {
	return internal.NewEvent(event, internal.WhenEvent{})
}
