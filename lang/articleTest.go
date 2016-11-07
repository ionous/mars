package lang

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/test"
)

var ArticleTest = test.Suite{"Articles",
	test.Setup(
		The("kind", Called("lamp-post"), Exists()),
		The("kind", Called("soldiers"), Exists(),
			Has("indefinite article", "some")),
		The("kind", Called("Trevor"), Exists(),
			Is("proper named")),
	),
	test.Trials(
		// examples from inform7
		// PHRASE: say "[a (object)]" & say "[an (object)]"
		test.Execute(
			"You can only just make out a lamp-post.",
			Say("You can only just make out", ALower{Named{"lamp post"}}, ".")),
		test.Execute("You can only just make out Trevor.",
			Say("You can only just make out", ALower{Named{"Trevor"}}, ".")),
		test.Execute("You can only just make out some soldiers.",
			Say("You can only just make out", ALower{Named{"soldiers"}}, ".")),
		// PHRASE: say "[A (object)]" & say "[An (object)]"
		test.Execute("A lamp-post can be made out in the mist.",
			Say(AnUpper{Named{"lamp post"}}, "can be made out in the mist.")),
		test.Execute("Trevor can be made out in the mist.",
			Say(AnUpper{Named{"Trevor"}}, "can be made out in the mist.")),
		test.Execute("Some soldiers can be made out in the mist.",
			Say(AnUpper{Named{"soldiers"}}, "can be made out in the mist.")),
		// PHRASE: say "[the (object)]"
		test.Execute("You can only just make out the lamp-post.",
			Say("You can only just make out", TheLower{Named{"Lamp post"}}, ".")),
		test.Execute("You can only just make out Trevor.",
			Say("You can only just make out", TheLower{Named{"trevor"}}, ".")),
		test.Execute("You can only just make out the soldiers.",
			Say("You can only just make out", TheLower{Named{"soldiers"}}, ".")),
		// PHRASE: say "[The (object)]"
		test.Execute("The lamp-post may be a trick of the mist.",
			Say(TheUpper{Named{"lamp-post"}}, "may be a trick of the mist.")),
		test.Execute("Trevor may be a trick of the mist.",
			Say(TheUpper{Named{"trevor"}}, "may be a trick of the mist.")),
		test.Execute("The soldiers may be a trick of the mist.",
			Say(TheUpper{Named{"soldiers"}}, "may be a trick of the mist.")),
	),
}
