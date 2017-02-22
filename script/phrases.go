package script

import (
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/internal"
	"github.com/ionous/sashimi/util/errutil"
)

type FragmentCollection interface {
	GetFragments() []backend.Fragment
}

// where is a good place for things like this?
// the commands, ideally, would be in one place -- and probably *not* internal
// the wrapper would be somewhere else
// maybe instead of "script" as a package -- we should have "nouns" -- and script is solely the generator of "nouns" inside of go.
type KnownAsList struct {
	fragments []backend.Fragment
}

// Add additional aliases for the current subject.
func (fc KnownAsList) And(name string) KnownAsList {
	fc.fragments = append(fc.fragments, internal.KnownAs{name})
	return fc
}

func (fc KnownAsList) GetFragments() []backend.Fragment {
	return fc.fragments
}

type ChoiceList struct {
	fragments []backend.Fragment
}

func (fc ChoiceList) And(choice string) ChoiceList {
	fc.fragments = append(fc.fragments, internal.Choice{choice})
	return fc
}

func (fc ChoiceList) GetFragments() []backend.Fragment {
	return fc.fragments
}

// The targets a noun for new assertions.
func The(target string, fragments ...interface{}) backend.Directive {
	flat := []backend.Fragment{}
	for i, src := range fragments {
		switch val := src.(type) {
		case backend.Fragment:
			flat = append(flat, val)
		case FragmentCollection:
			for _, f := range val.GetFragments() {
				flat = append(flat, f)
			}
		default:
			panic(errutil.New("script noun phrase expects fragments of lists of fragments. at", i, "got", val))
		}
	}
	return internal.NounDirective{target, flat}
}

// Understand builds statements for parsing player input.
func Understand(s string) internal.ParserPartial {
	return internal.ParserPartial{}.And(s)
}

// Our is an alias for The
var Our = The

// Called asserts the existence of a class or instance.
// For example, The("room", Called("home"))
func Called(subject string) internal.ScriptSubject {
	return internal.ScriptSubject{subject}
}

func HasSingularName(subject string) internal.ScriptSingular {
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
func Is(choice string, choices ...string) (ret ChoiceList) {
	ret = ret.And(choice)
	for _, choice := range choices {
		ret = ret.And(choice)
	}
	return
}

// IsKnownAs declares an alias for the current subject.
// ex. The("cabinet", IsKnownAs("armoire").And("..."))
func IsKnownAs(name string, names ...string) (ret KnownAsList) {
	ret = ret.And(name)
	for _, name := range names {
		ret = ret.And(name)
	}
	return
}

// Have asserts the existance of a property for all instances of a given class.
// For relations, use HaveOne or HaveMany.
func Have(name string, class string) backend.Fragment {
	return internal.ClassProperty{name, class}
}

// HaveOne establishes a one-to-one, or one-to-many relation.
func HaveOne(name string, class string) internal.PartialRelation {
	return internal.NewHaveOne(name, class)
}

// HaveMany establishs a many-to-one relation.
func HaveMany(name string, class string) internal.PartialRelation {
	return internal.NewHaveMany(name, class)
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
func Can(verb string) internal.CanDoPhrase {
	return internal.CanDoPhrase{ActionName: verb}
}

// To assigns runtime statements to a default action handler.
func To(action string, call rt.Execute, calls ...rt.Execute) backend.Fragment {
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
