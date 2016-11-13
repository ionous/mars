package core

import (
	. "github.com/ionous/mars/script"
	// "github.com/ionous/mars/rt"
	"github.com/ionous/mars/core/stream"
	"github.com/ionous/mars/script/test"
)

func init() {
	t := NewScript()
	t.The("kinds",
		Called("tests"),
		Have("rank", "num"))

	t.The("test",
		Called("C"), Has("rank", 1))
	t.The("test",
		Called("A"), Has("rank", 4))
	t.The("test",
		Called("B"), Has("rank", 2))
	t.The("test",
		Called("D"), Has("rank", 3))

	addTest("Sort",
		test.Setup(t).Try("sorted",
			test.Execute(Say(
				stream.KeySort{"name", stream.ClassStream{Name: "tests"}})).
				Match("A B C D"),
			test.Execute(Say(
				stream.KeySort{"rank", stream.ClassStream{Name: "tests"}})).
				Match("C B D A"),
		),
		test.Setup(t).Try("first",
			test.Expect(IsObj{Named{"A"}, stream.First{In: stream.KeySort{"name", stream.ClassStream{Name: "tests"}}}}),
		),
		// test.Setup(t).Try("sorted",
		// 	test.Execute(Say( PropertyTextList{"text list", Named{"sorting"}})).
		// 		Match("D B A C"),
		// ))
	)
}
