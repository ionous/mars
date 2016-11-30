package core

import (
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/script/test"
)

func init() {
	t := NewScript()
	t.The("kinds",
		Have("amBlank", "text"),
		Have("amSet", "text"))

	t.The("kind",
		Called("test"),
		HasText("amSet", T("original")))

	pkg.AddTest("Core Tests",
		test.Setup(t).Try("changing values",
			test.Expect(IsValid{Named{"test"}}),
			test.Expect(IsText{PropertyText{"amSet", Named{"test"}}, EqualTo{}, T("original")}),
			test.Execute(SetTxt{"amSet", Named{"test"}, T("new")}).Expect(IsText{PropertyText{"amSet", Named{"test"}}, EqualTo{}, T("new")}),
			test.Expect(IsEmpty{PropertyText{"amBlank", Named{"test"}}}),
			test.Execute(SetTxt{"amBlank", Named{"test"}, T("not blank any more")}).Expect(IsNot{IsEmpty{PropertyText{"amBlank", Named{"test"}}}}),
		))
}
