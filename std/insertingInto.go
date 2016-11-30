package std

import (
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

// from inform:
// 	"insert applies into two things", doesnt use preferably held. [ possibly a mistake? ]
//  "convert insert into drop where possible" ( if the second noun is down, or if the actor is in the second noun )
// 		?FIX? i dont understand.
//  "can't insert what isn't held"
// 		? again, isnt that point of "something preferably held"?
//  "can't insert something into itself"
//  "can't insert into closed containers"
// 	X "can't insert into what's not a container": implicit in this definition
//  "can't insert clothes being worn"
//  FIX: can't insert if this exceeds carrying capacity
// * "carry out inserting": now in second noun
// * "concise report inserting rule": "done"
// * "standard reporting rule": "actor put thing into thing"

/// insert it into, receive insertion, being inserted.
func init() {
	pkg.AddScript("InsertingInto", // 1. source
		The("actors",
			// FIX? word-wise this is wrong ( see tickle-it-with, though it is "correct" )
			Can("insert it into").And("inserting it into").RequiresOne("container").AndOne("prop"),
			To("insert it into", g.ReflectWithContext("receive insertion")),
			//  can't insert clothes being worn
			Before("inserting it into").Always(
				Choose{
					If: g.The("action.Context").Object("wearer").Exists(),
					True: g.Go(
						g.Say("You can't insert worn clothing."),
						// FIX: try taking off the noun
						g.StopHere(),
					),
				},
			),
			//  can't insert what isn't held
			Before("inserting it into").Always(
				Choose{
					If: Carrier(g.The("action.Context")).Equals(g.The("action.Source")),
					False: g.Go(
						g.Say("You aren't holding", g.The("action.Context").Lower(), "."),
						g.StopHere(),
					),
				}),
			//  can't insert something into itself
			Before("inserting it into").Always(
				Choose{
					If: g.The("action.Target").Equals(g.The("action.Context")),
					True: g.Go(
						g.Say("You can't insert something into itself."),
						g.StopHere(),
					),
				}),
		),

		// 2. containers
		// FIX FIX FIX could this be "receive"? the more shared events the better,
		// and that would certainly work for any acquisition: actor, supporter,..?
		// the only problem of course is that "be inserted" using specific reporting
		// maybe, rather than chaining events 1->2->3 we could do: 1->{2,3}
		// and check the return result?[the ran default action status of the evt.]
		The("containers",
			Can("receive insertion").And("receiving insertion").RequiresOne("prop").AndOne("actor"),
			//  can't insert into closed containers
			Before("receiving insertion").Always(
				Choose{
					If: g.The("container").Is("closed"),
					True: g.Go(
						g.Say(g.The("container").Upper(), "is closed."),
						g.StopHere(),
					),
				}),
			To("receive insertion", g.ReflectWithContext("be inserted")),
		),
		// 3. context
		The("props",
			Can("be inserted").And("being inserted").RequiresOne("actor").AndOne("container"),
			To("be inserted", g.Go(
				g.Say("You insert", g.The("action.Source").Lower(), "into", g.The("action.Context").Lower(), "."),
				Insert(g.The("action.Source")).Into(g.The("action.Context")),
			),
			)),
		// input: actor, container, prop
		Understand("put|insert {{something else}} in|inside|into {{something}}").
			And("drop {{something else}} in|into|down {{something}}").
			As("insert it into"),
	)
}

func Insert(what rt.ObjEval) InsertIntoPhrase {
	return InsertIntoPhrase{what}
}

func (ins InsertIntoPhrase) Into(where rt.ObjEval) rt.Execute {
	return AssignParent(ins.what, "enclosure", where)
}

type InsertIntoPhrase struct {
	what rt.ObjEval
}
