package std

// import (
// 	. "github.com/ionous/mars/core"
// 	. "github.com/ionous/mars/lang"
// 	. "github.com/ionous/mars/script"
// 	"github.com/ionous/mars/script/g"
// 	"github.com/ionous/mars/script/test"
// )

// var endTurn = g.Go(
// 			// FIX: we should have a "starting the turn" instead
// 			// currently, tp get the "frame" counter to see a different number than the last frame
// 			// due to this frame's player input -- we have to increment both on end turn, and ending the story
// 			turnCount := g.The("story").Num("turn count") + 1
// 			g.The("story").SetNum("turn count", turnCount)
// 			//
// 			Choose{
// 				If: g.The("story").Is("scored"),

// 				g.The("status bar").SetText("right", fmt.Sprintf("%d/%d", g.The("story").Num("score"), int(turnCount)))
// 			}
// 		}

// // System actions
// func init() { addScript("Stories",		// in lieu of singletons, globals.
// 		// globals are transmitted to the client in the default view.
// 		The("kinds", Called("globals"), Exist()),
// 		The("globals",
// 			Called("stories").WithSingularName("story"),
// 			Have("author", "text"),
// 			Have("headline", "text"),
// 			AreEither("scored").Or("unscored").Usually("unscored"),
// 			// Inform uses global variables, which would be much nicer.
// 			// ex. The maximum score is 1.
// 			Have("score", "num"),
// 			Have("maximum score", "num"),
// 			Have("turn count", "num"),
// 			AreOneOf("playing", "completed", "starting").Usually("starting"),
// 		),

// 		The("stories",
// 			Can("commence").And("commencing").RequiresNothing(),
// 			Can("end the story").And("ending the story").RequiresNothing(),
// 			Can("end turn").And("ending the turn").RequiresNothing(),
// 			Before("commencing").Always(g.Go(
// 				inst := g.The("status bar")
// 				title := g.The("story").Get("name").Text()
// 				author := g.The("story").Get("author").Text()

// 				tag := fmt.Sprintf(`"%s" by %s`, title, author)
// 				inst.Get("left").SetText(title)
// 				inst.Get("right").SetText(tag)
// 			}),
// 			Before("ending the turn").Always(g.Go(
// 				story := g.The("story")
// 				if g.The("story").Is("completed") {
// 					g.StopHere()
// 				}
// 			}),
// 			To("end turn", endTurn),
// 			After("ending the story").Always(endTurn),
// 		)

// 		The("stories",
// 			To("commence", g.Go(
// 				// FIX: duplication with end turn
// 				story := g.The("story")
// 				if g.The("story").Is("scored") {
// 					score := g.The("story").Num("score")
// 					status := fmt.Sprintf("%d/%d", int(score), int(0))
// 					g.The("status bar").SetText("right", status)
// 				}
// 				room := g.The("player").Object("whereabouts")
// 				if !room.Exists() {
// 					rooms := g.Query("rooms", false)
// 					if !rooms.HasNext() {
// 						panic("story has no rooms")
// 					}
// 					room = rooms.Next()
// 				}
// 				g.The("story").Go("set initial position", g.The("player"), room).Then(g.Go(
// 					g.The("story").Go("print the banner").Then(g.Go(
// 						room = g.The("player").Object("whereabouts")
// 						// FIX: Go() should handle both Name() and ref
// 						g.The("story").Go("describe the first room", room).Then(g.Go(
// 							g.The("story").IsNow("playing")
// 						})
// 					})
// 				})
// 			}))

// 		The("stories",
// 			Have("player input", "text"),
// 			Can("parse player input").And("parsing player input").RequiresNothing())

// 		The("stories",
// 			To("end the story", g.Go(
// 				story := g.The("story")
// 				g.Say("*** The End ***")
// 				g.The("story").IsNow("completed")

// 				if g.The("story").Is("scored") {
// 					score, maxScore, turnCount := g.The("story").Num("score"), g.The("story").Num("maximum score"), g.The("story").Num("turn count")
// 					g.Say(fmt.Sprintf("In that game you scored %d out of a possible %d, in %d turns.",
// 						int(score), int(maxScore), int(turnCount)+1,
// 					))
// 				}
// 			}))

// 		The("stories",
// 			Can("set initial position").
// 				And("setting initial position").
// 				RequiresOne("actor").
// 				AndOne("room"),
// 			To("set initial position", g.Go(
// 				player := g.The("action.Target")
// 				room := g.The("action.Context")
// 				player.Set("whereabouts", room) // Now("player's whereabouts is $room")
// 			}))

// 		The("stories",
// 			Can("describe the first room").
// 				And("describing the first room").RequiresOne("room"),
// 			To("describe the first room", g.Go(
// 				room := g.The("action.Target")
// 				room.Go("report the view")
// 			}),
// 		)
// 				The("stories",
// 			Can("print the banner").
// 				And("printing the banner").RequiresNothing(),

// 			To("print the banner",
// 				g.Context{g.Our("story"),
// 					g.Statements{

// 						g.Say(g.GetText{"name"}, "."),
// 						g.Say(g.ChooseText{
// 							If:    g.Empty{g.GetText{"headline"}},
// 							False: g.GetText{"headline"},
// 							// FIX: default for headline in class.
// 							True: g.MakeText("An interactive fiction"),
// 						}, "by", g.GetText{"author"}, "."),
// 						g.Say(std.VersionString),
// 					},
// 				}))
// 	})
// }

// 	})
// }
