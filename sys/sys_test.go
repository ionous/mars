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
		var x ExecuteList
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

// 		Try("initially zero", IsNumber{
// 			PropertyNum{Id("i"), "num"},
// 			EqualTo,
// 			I(0),
// 		}),

// 		SetNum{
// 			PropertyNum{Id("i"), "num"},
// 			I(5),
// 		},
// 		Try("now greater than 1",
// 			IsNumber{
// 				PropertyNum{Id("i"), "num"},
// 				GreaterThan,
// 				I(1),
// 			}),
// 		Try("now greater than or equal to 5",
// 			IsNumber{
// 				PropertyNum{Id("i"), "num"},
// 				GreaterThan | EqualTo,
// 				I(5),
// 			}),
// 		g.Say("hello"),
// 		PrintLine{PrintNum{
// 			PropertyNum{Id("i"), "num"},
// 		}},
// 		Try("not greater than or lesser to 5",
// 			IsNot{IsNumber{
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
// 			IsObject{
// 				PropertyRef{Id("i"), "object"},
// 				Id("i"),
// 			}),

// 		Try("i should not equal x",
// 			IsNot{IsObject{
// 				PropertyRef{Id("i"), "object"},
// 				Id("x"),
// 			}}),

// 		SetObj{
// 			PropertyRef{Id("i"), "object"},
// 			Nothing(),
// 		},

// 		Try("i should now be nothing",
// 			IsObject{
// 				PropertyRef{Id("i"), "object"},
// 				Nothing(),
// 			}),

// 		Try("zero test should choose i",
// 			IsObject{
// 				ChooseObj{
// 					If:   IsNumber{I(0), EqualTo, I(0)},
// 					True: Id("i"),
// 				},
// 				Id("i"),
// 			}),

// 		Context{Id("i"), g.Say("In that game you scored", GetNum{"num"}, "out of a possible", I(1000), ".")},
// 	// not implemented:
// 	// g.The("i").Go("give", g.The("x"), 5),

// 	// std.Speaker("player").Says("I don't want to think where that came from."),
// 	)

// }

//2683a5e6912634486b42b4fdf77a0af7
// SetNumber( that uses the buffer capturing)
// TheArticleName, Capitalize, FullStop,DefiniteName
//
//31cdc07aa023e1603e33ed0134d6516f
// i need to be able to say [The] {{thing}}
// for it to detect teh capitial
// for it to know i want definite article if it exists

//1e2a4b3156ede9027566370cfa46706e
// combine a series of text

//1507726b96b8e974e5adcc8fbb1dc695  -> switch
// 17128406c83013cdda9fe25d7bcfddb2-> global
// Empty? if text := g.The("room").Get("scent").GetText(); len(text) > 0 {
//1898a5340e0f100d789aec582201d169-> object list
//
// quips.Converse
//(quips.ChangeTopic("Space animals"))

//3de0a3e45ca7af7e4db3dba5e8df7d9f
// FromClass

//1be0f48c532918678e0eb17bd9d75c67
// addition/subtraction

//4113753b19e079e9f9825f60a9af6c04
//ListContents, RefValueList, range

// Random
// but, really we want the things everyone else has: shuffle etc.

// 465e35a28cc716ffabacdfcfe6e828f4
// Enclosure

// The("placket").Is() ->
// would could add phrases to help with translation --
// basically dusplicate the object interface
// you could put it in a "compat" "g" pkg
// and then the conversion becomes more and more simple, commas instead of newlines
