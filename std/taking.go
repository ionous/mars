package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	pkg.AddScript("Taking",
		The("actors",
			Can("take it").And("taking it").RequiresOnly("prop"),
			To("take it", g.ReflectToTarget("report take")),
		),
		The("props",
			Can("report take").And("reporting take").RequiresOnly("actor"),
			To("report take",
				Choose{
					If:    Enclosure(g.The("actor")).Equals(Enclosure(g.The("prop"))),
					False: g.Say("That isn't available."),
					True: Choose{
						If:   g.The("prop").Is("scenery"),
						True: g.Say("You can't take scenery."),
						False: Choose{
							If:   g.The("prop").Is("fixed in place"),
							True: g.Say("It is fixed in place."),
							False: Choose{
								If: Carrier(g.The("prop")).Exists(),
								True: Choose{
									If:    Carrier(g.The("prop")).Equals(g.The("actor")),
									False: g.Say("That'd be stealing!"),
									True:  g.Say(g.The("actor").Upper(), "already has that!"),
								},
								False: g.Go(
									Choose{ // separate report action?
										If:    g.The("actor").Equals(g.The("player")),
										True:  g.Say("You take", g.The("prop").Lower(), "."),
										False: g.Say(g.The("actor").Upper(), "takes", g.The("prop").Lower(), "."),
									},
									g.Go(Give("prop").To("actor"))),
							},
						},
					},
				},
			)),
		// understandings:
		Understand("take|get {{something}}").
			And("pick up {{something}}").
			And("pick {{something}} up").
			As("take it"),
	)
}

// touchable/reach inside checks --, plus:
// check    an actor taking  can't take yourself rule
//       A   "[We] [are] always self-possessed."
// check    an actor taking  can't take other people rule
//       A   "I don't suppose [the noun] [would care] for that."
// check    an actor taking  can't take component parts rule
//       A   "[regarding the noun][Those] [seem] to be a part of [the whole]."
// check    an actor taking  can't take people's possessions rule
//       A   "[regarding the noun][Those] [seem] to belong to [the owner]."
// check    an actor taking  can't take items out of play rule
//       A   "[regarding the noun][Those] [aren't] available."
// check    an actor taking  can't take what you're inside rule
//       A   "[We] [would have] to get [if noun is a supporter]off[otherwise]out of[end if] [the noun] first."
// check    an actor taking  can't take what's already taken rule
//       A   "[We] already [have] [regarding the noun][those]."
// check    an actor taking  can't take scenery rule
//       A   "[regarding the noun][They're] hardly portable."
// check    an actor taking  can only take things rule
//       A   "[We] [cannot] carry [the noun]."
// check    an actor taking  can't take what's fixed in place rule
//       A   "[regarding the noun][They're] fixed in place."
// check    an actor taking  use player's holdall to avoid exceeding carrying capacity rule
//       A   "(putting [the transferred item] into [the current working sack] to make room)[command clarification break]"
// check    an actor taking  can't exceed carrying capacity rule
//       A   "[We]['re] carrying too many things already."
//
// carry out    an actor taking  standard taking rule
//
// report    an actor taking  standard report taking rule
//       A   "Taken."
//       B   "[The actor] [pick] up [the noun]."
