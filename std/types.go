package std

import (
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/script/g"
)

var Types = Script(
	// FIX: hierarchy is a work in progress.
	// kinds> stories, rooms, objects > actors (> animals),  props(> openers(> doors,containers), supporters, devices)

	// vs. descriptions in "kind"
	// it seems to make sense for now to have two separate description fields.
	// rooms like to say their description, while objects like to say their brief initial appearance ( or name, if there's none. )
	The("rooms",
		Have("description", "text")),

	The("objects",
		Have("description", "text"),
		Have("brief", "text")),

	// inform's rooms: lighted, dark; unvisited, visited; description, region
	The("kinds",
		Called("rooms"),
		AreEither("visited").Or("unvisited").Usually("unvisited"),
	),

	// the class hierarchy means rooms cant contain other rooms.
	The("rooms",
		HaveMany("contents", "objects").
			Implying("objects", HaveOne("whereabouts", "room"))),

	// things		unlit, lit
	// 	inedible, edible
	//
	// 	unmarked for listing, marked for listing
	// 	described, undescribed : i think, whether to appear in any room descriptions
	// 	mentioned, unmentioned : i think, whether it has appeared in a room description
	// bool	scenery
	// 	wearable
	// 	pushable between rooms
	// 	.handled
	// 	.description (in objects and rooms)
	// 	.initial appearance (brief)
	// 	matching key
	The("kinds",
		Called("objects"),
		Exist()),

	// hrmmm.... are actors really scenery? handled?
	The("objects",
		AreEither("unhandled").Or("handled"),
		AreEither("scenery").Or("not scenery").Usually("not scenery")),

	// nothing special: just a handy name to mirror inform's.
	The("actors",
		Called("animals"),
		Exist()),

	// hrmm.... implies actors are takeable.
	The("objects",
		Called("actors"),
		HaveMany("clothing", "objects").
			Implying("objects", HaveOne("wearer", "actor")),
		HaveMany("inventory", "objects").
			Implying("objects", HaveOne("owner", "actor"))),

	// changed: used to have "equipment" for held objects
	// will either implement some sort of "relation with value"
	// or will add a "held","holdable", state to objects.

	The("objects",
		Called("props"),
		AreEither("portable").Or("fixed in place"),
	),

	The("props",
		Called("supporters"),
		HaveMany("contents", "objects").
			Implying("objects", HaveOne("support", "supporter"))),
)
