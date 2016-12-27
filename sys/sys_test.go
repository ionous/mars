package sys

import (
	"bytes"
	"encoding/gob"
	. "github.com/ionous/mars/core"
	"github.com/ionous/mars/rt"
	"github.com/ionous/mars/rtm"
	"github.com/ionous/mars/script/g"
	"github.com/ionous/mars/std"
	"github.com/ionous/sashimi/compiler/model/modeltest"
	"github.com/ionous/sashimi/metal"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Run(t *testing.T, a ...rt.Execute) {
	src := make(metal.ObjectValueMap)
	m := metal.NewMetal(modeltest.NewModel(), src)
	run := rtm.NewRtm(m)
	run.PushOutput(os.Stdout)

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	// Create an encoder and send some values.
	if err := enc.Encode(a); assert.NoError(t, err, "encode") {
		var x Statements
		if err := dec.Decode(&x); assert.NoError(t, err, "decode") {
			// Create a decoder and receive some values.
			if err := x.Execute(run); assert.NoError(t, err, "execute") {
				//
			}
		}
	}
}

// func Test_TooBig(t *testing.T) {
// 	rtm.RegisterTypes(gob.RegisterName, rt.BuiltIn{}, Core{}, std.Std{})

// 	Run(t,
// 		Try("i exists", IsValid{Id("i")}),
// 		Try("nope does not exist", IsNot{IsValid{Id("nope")}}),
// 		Try("i defaults no", IsState{Id("i"), "no"}),
// 		//"failed okay with ChangeState I does not have choice Borrigard"
// 		Fails{
// 			Change(Id("i")).To("borrigard"),
// 			"no such state should exist"},
// 		//
// 		Change(Id("i")).To("yes"),
// 		//
// 		Try("i now yes", IsState{Id("i"), "yes"}),

// 		Try("initially zero", IsNum{
// 			PropertyNum{Id("i"), "num"},
// 			EqualTo,
// 			I(0),
// 		}),

// 		SetNum{
// 			PropertyNum{Id("i"), "num"},
// 			I(5),
// 		},
// 		Try("now greater than 1",
// 			IsNum{
// 				PropertyNum{Id("i"), "num"},
// 				GreaterThan,
// 				I(1),
// 			}),
// 		Try("now greater than or equal to 5",
// 			IsNum{
// 				PropertyNum{Id("i"), "num"},
// 				GreaterThan | EqualTo,
// 				I(5),
// 			}),
// 		g.Say("hello"),
// 		PrintLine{PrintNum{
// 			PropertyNum{Id("i"), "num"},
// 		}},
// 		Try("not greater than or lesser to 5",
// 			IsNot{IsNum{
// 				PropertyNum{Id("i"), "num"},
// 				GreaterThan | LesserThan,
// 				I(5),
// 			}}),

// 		SetObj{
// 			PropertyRef{Id("i"), "object"},
// 			Id("i"),
// 		},
// 		//
// 		Try("i should point to i",
// 			IsObj{
// 				PropertyRef{Id("i"), "object"},
// 				Id("i"),
// 			}),

// 		Try("i should not equal x",
// 			IsNot{IsObj{
// 				PropertyRef{Id("i"), "object"},
// 				Id("x"),
// 			}}),

// 		SetObj{
// 			PropertyRef{Id("i"), "object"},
// 			Nothing(),
// 		},

// 		Try("i should now be nothing",
// 			IsObj{
// 				PropertyRef{Id("i"), "object"},
// 				Nothing(),
// 			}),

// 		Try("zero test should choose i",
// 			IsObj{
// 				ChooseObj{
// 					If:   IsNum{I(0), EqualTo, I(0)},
// 					True: Id("i"),
// 				},
// 				Id("i"),
// 			}),

// 		Using{Id("i"), g.Say("In that game you scored", GetNum{"num"}, "out of a possible", I(1000), ".")},
// 	// not implemented:
// 	// g.The("i").Go("give", g.The("x"), 5),

// 	// std.Speaker("player").Says("I don't want to think where that came from."),
// 	)

// }

//
// quips.Converse
//(quips.ChangeTopic("Space animals"))

//4113753b19e079e9f9825f60a9af6c04
//ListContents, RefValueList, range

// Random
// but, really we want the things everyone else has: shuffle etc.
