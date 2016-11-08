package std

import (
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/script/g"
)

var Supporters = Script(
	The("props",
		Called("supporters"),
		HaveMany("contents", "objects").
			Implying("objects", HaveOne("support", "supporter"))),
)
