package std

import (
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/script/g"
)

func init() {
	addScript("Actors",
		// hrmm.... implies actors are takeable.
		The("objects",
			Called("actors"),
			HaveMany("clothing", "objects").
				Implying("objects", HaveOne("wearer", "actor")),
			HaveMany("inventory", "objects").
				Implying("objects", HaveOne("owner", "actor"))),
		// nothing special: just a handy name to mirror inform's.
		The("actors",
			Called("animals"),
			Exist()),
	)
}
