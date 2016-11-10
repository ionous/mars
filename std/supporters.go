package std

import (
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/script/g"
)

func init() {
	addScript("Supporters",
		The("props",
			Called("supporters"),
			HaveMany("contents", "objects").
				Implying("objects", HaveOne("support", "supporter"))),
	)
}
