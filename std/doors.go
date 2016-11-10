package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

func init() {
	addScript("Doors",
		The("openers", //--> openingClosing.go
			Called("doors"),
			Exist()),

		The("doors",
			When("printing name text").
				Always(
					g.Say(g.The("door").Upper(),
						ChooseText{
							If: g.The("door").Is("hinged"),
							True: ChooseText{
								If:    g.The("door").Is("open"),
								True:  T(" (open)"),
								False: T(" (closed)"),
							},
						}),
					g.StopHere(),
				)),

		// CAN WE DEFAULT (USUALLY(X)) DOORS TO fixed-in-place???
		The("doors", Before("reporting take").Always(
			g.Say("It is fixed in place."),
			g.StopHere())),
	)
}
