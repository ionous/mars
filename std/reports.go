package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
	"github.com/ionous/mars/script/g"
)

var Reports = []backend.Spec{
	// MARS: re-evaluate. something like: objects have a default property name script which evals to text
	// print name may ( or may not ) still exist as a way to print that name:
	// more likely people would just use the property, perhaps the default is uncapitalized, and then upper name adds capitalization rules.
	// this only differens from "actions" ( can/do ) in that it returns a string
	// The("objects",
	// 	Can("print name").And("printing name text").RequiresNothing(),
	// 	To("print name", g.Say(g.The("object").Upper(), ".")),
	// ),

	// name (status)
	The("containers",
		When("printing name text").
			Always(g.Say(
				g.The("object").Upper(), ChooseText{
					If: g.The("object").Is("closed"),
					True: ChooseText{
						If:   g.The("object").Is("hinged"),
						True: T("( closed )"),
					},
					False: ChooseText{
						If: g.The("object").ObjectList("contents").Empty(),
						True: ChooseText{
							If:   Any(g.The("object").Is("transparent")).Or(g.The("object").Is("open")),
							True: T("( empty )"),
						},
					},
				}),
				g.StopHere(),
			)),

	// MARS: fix!
	// The("doors",
	// 	When("printing name text").
	// 		Always(func(g G.Play) {
	// 			text := DefiniteName(g, "door", func(obj G.IObject) (status string) {
	// 				if obj.Is("hinged") {
	// 					if obj.Is("open") {
	// 						status = "open"
	// 					} else {
	// 						status = "closed"
	// 					}
	// 				}
	// 				return status
	// 			})
	// 			text = lang.Capitalize(text)
	// 			g.Say(text)
	// 			g.StopHere()
	// 		})),

	// The("rooms",
	// 	Can("report the view").And("reporting the view").RequiresNothing(),
	// 	When("reporting the view").Always(func(g G.Play) {
	// 		room := g.The("room")
	// 		g.The("status bar").SetText("left", lang.Titleize(room.Text("Name")))
	// 	}),
	// 	After("reporting the view").Always(func(g G.Play) {
	// 		g.Go(Change("room").To("visited"))
	// 	}),
	// 	To("report the view", func(g G.Play) {
	// 		g.Go(View("room"))
	// 	})),
}

// test container names....
