package std

import (
	. "github.com/ionous/mars/core"
	. "github.com/ionous/mars/lang"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/g"
)

var Attacking = Script(
	The("actors",
		Can("attack it").And("attacking it").RequiresOne("object"),
		To("attack it", g.ReflectToTarget("report attack"))),

	The("objects",
		Can("report attack").And("reporting attack").RequiresOne("actor"),
		To("report attack", Choose{
			If:   g.Our("player").Equals(g.Our("actor")),
			True: g.Say("Violence isn't the answer."),
		})),

	Understand("attack|break|smash|hit|fight|torture {{something}}").
		And("wreck|crack|destroy|murder|kill|punch|thump {{something}}").
		As("attack it"),
)
