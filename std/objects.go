package std

import (
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/script/g"
)

var Objects = Script(
	// FIX: hierarchy is a work in progress.
	// kinds> stories, rooms, objects > actors (> animals),  props(> openers(> doors,containers), supporters, devices)

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

	// changed: used to have "equipment" for held objects
	// will either implement some sort of "relation with value"
	// or will add a "held","holdable", state to objects.

	The("objects",
		Called("props"),
		AreEither("portable").Or("fixed in place"),
	),
)
