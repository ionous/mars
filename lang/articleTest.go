package lang

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/test"
)

func init() {
	pkg.AddTest("Articles",
		test.Setup(NewScript(
			The("kind", Called("lamp-post"), Exists()),
			The("kind", Called("soldiers"), Exists(),
				HasText("indefinite article", T("some"))),
			The("kind", Called("Trevor"), Exists(),
				Is("proper named"))),
		).Try("test articles",
			// examples from inform7
			// PHRASE: say "[a (object)]" & say "[an (object)]"
			test.Execute(
				Say("You can only just make out", ALower{Named{"lamp post"}}, ".")).
				Match("You can only just make out a lamp-post."),
			test.Execute(
				Say("You can only just make out", ALower{Named{"Trevor"}}, ".")).
				Match("You can only just make out Trevor."),
			test.Execute(
				Say("You can only just make out", ALower{Named{"soldiers"}}, ".")).
				Match("You can only just make out some soldiers."),
			// PHRASE: say "[A (object)]" & say "[An (object)]"
			test.Execute(
				Say(AnUpper{Named{"lamp post"}}, "can be made out in the mist.")).
				Match("A lamp-post can be made out in the mist."),
			test.Execute(
				Say(AnUpper{Named{"Trevor"}}, "can be made out in the mist.")).
				Match("Trevor can be made out in the mist."),
			test.Execute(
				Say(AnUpper{Named{"soldiers"}}, "can be made out in the mist.")).
				Match("Some soldiers can be made out in the mist."),
			// PHRASE: say "[the (object)]"
			test.Execute(
				Say("You can only just make out", TheLower{Named{"Lamp post"}}, ".")).
				Match("You can only just make out the lamp-post."),
			test.Execute(
				Say("You can only just make out", TheLower{Named{"trevor"}}, ".")).
				Match("You can only just make out Trevor."),
			test.Execute(
				Say("You can only just make out", TheLower{Named{"soldiers"}}, ".")).
				Match("You can only just make out the soldiers."),
			// PHRASE: say "[The (object)]"
			test.Execute(
				Say(TheUpper{Named{"lamp-post"}}, "may be a trick of the mist.")).
				Match("The lamp-post may be a trick of the mist."),
			test.Execute(
				Say(TheUpper{Named{"trevor"}}, "may be a trick of the mist.")).
				Match("Trevor may be a trick of the mist."),
			test.Execute(
				Say(TheUpper{Named{"soldiers"}}, "may be a trick of the mist.")).
				Match("The soldiers may be a trick of the mist."),
		))

}
