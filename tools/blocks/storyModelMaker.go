package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

// return values? its basically for build cmd i think
// does add process really need command info? why?

func NewStoryModel(db ScriptDB, types inspect.Types) *ModelMaker {
	m := NewModelMaker(db, types)

	m.Blocks.CreateOn("Directive", "dl-group")
	m.Blocks.CreateOn("Execute", "dl-line")
	m.Params.AddFilter("InLocation.Location", QuoteFilter)
	m.Params.AddFilter("ScriptSubject.Subject", QuoteFilter)
	m.Params.AddFilter("Text.Value", QuoteFilter)
	m.Params.AddFilter("Chapter.Name", BlockEndFilter)

	m.Params.AddProcess("Execute?array=true", func(_ *inspect.ParamInfo, stack *Stack) error {
		return stack.NewBlock("dl-scope", m.BuildArray)
	})

	m.Params.AddProcess("AllTrue.Test", func(p *inspect.ParamInfo, stack *Stack) error {
		return stack.NewInlineBlock(p.Name, func(stack *Stack) error {
			return m.BuildElements(stack, func(cmd *inspect.CommandInfo, array *ArrayData, i int) (err error) {
				if e := m.BuildCmd(cmd, stack); e != nil {
					err = e
				} else {
					stack.LastChild.Sep = CommaSep{false}
				}
				return
			})
		})

	})

	return m
}

// m.Params.AddProcess("NounDirective.Fragments", func(_ *inspect.ParamInfo, stack *Stack) error {
// 	return m.BuildElements(stack, func(cmd *inspect.CommandInfo, array *ArrayData, i int) (err error) {
// 		called := cmd.Name == "ScriptSubject"
// 		if e := m.BuildCmd(cmd, stack); e != nil {
// 			err = e
// 		} else if called {
// 			stack.LastChild.Sep = SpaceSep{}
// 		} else {
// 			stack.LastChild.Sep = CommaSep{true}
// 		}
// 		return
// 	})
// })

// m.Params.AddProcess("Directive?array=true", func(stack *Stack) error {
// 	return m.BuildElements(stack, func(data *ArrayData, i int) (err error) {
// 		isLast := (i + 1) == len(data.Array)
// 		if e := m.BuildCmd(stack); e != nil {
// 			err = e
// 		} else {
// 			stack.LastChild.Sep = FullStop{!isLast}
// 		}
// 		return
// 	})
// })
