package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/lang"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/script/test"
	"github.com/ionous/mars/std/compat"
)

var Desc = Script(
	The("objects",
		Have("description", "text"),
		Have("brief", "text")),

	The("objects",
		Can("print description").And("describing").RequiresNothing(),
		To("print description",
			Describe(g.The("object")),
		)),

	// hrmmm.... are actors really scenery? handled?
	The("objects",
		AreEither("unhandled").Or("handled"),
		AreEither("scenery").Or("not scenery").Usually("not scenery"),
	),

	The("objects",
		Can("print contents").And("printing contents").RequiresNothing()),

	// MARS: re-evaluate. something like: objects have a default property name script which evals to text
	// print name may ( or may not ) still exist as a way to print that name:
	// more likely people would just use the property, perhaps the default is uncapitalized, and then upper name adds capitalization rules.
	// this only differens from "actions" ( can/do ) in that it returns a string
	The("objects",
		Can("print name").And("printing name text").RequiresNothing(),
		To("print name", g.Say(AnUpper{g.The("object")}, ".")),
	),

	// FIX: When() puts the contents after the object
	// look at some default actions of the DOM
	// maybe it'd be better to put the print, not in the action,
	// but in a target handler: then this could be after by being in the capture.

	// FIX: After() isnt working well, it goes into the default action
	// but not all objects are containers, so it errors
)

// Describe shortcut
func Describe(obj compat.ScriptRef) rt.Execute {
	return Choose{
		If: All(obj.Exists(), IsNot{obj.Is("scenery")}),
		True: g.Go(
			Choose{
				If:   obj.Is("handled"),
				True: obj.Go("print name"),
				False:
				// MARS: first of / reduce instead of choose?
				// couldnt print name would have to get name
				// but thats what we want anyway.
				Choose{
					If:    IsEmpty{obj.Text("brief")},
					False: g.Say(obj.Text("brief")),
					True:  obj.Go("print name"),
				},
			},
			obj.Go("print contents"),
		),
	}
}

var DescTest = test.NewSuite("Desc",
	test.Setup(
		The("object", Called("plain ring"), Exists(),
			Has("brief", "Cast aside, as if worthless, is a plain brass ring.")),
		//"No better than the loops of metal the old women use for fastening curtains."
	).Try(
		test.Execute(
			Describe(g.The("plain ring"))).
			Match("Cast aside, as if worthless, is a plain brass ring."),
		test.Execute(g.The("plain ring").IsNow("handled")),
		test.Execute(Try("should be handled", IsState{g.The("plain ring"), "handled"})),
		test.Execute(
			Describe(g.The("plain ring"))).
			Match("A plain ring."),
	),
)
