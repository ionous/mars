package blocks

import (
	"github.com/ionous/mars/tools/inspect"
)

func BlockEndFilter(prim string) string {
	return prim + `.`
}
func QuoteFilter(prim string) string {
	return `"` + prim + `"`
}

func FullStopSep(*Block, int) string {
	return "."
}
func StopGapSep(*Block, int) string {
	return ". "
}
func CommaSep(block *Block, idx int) (ret string) {
	cnt := len(block.Spans)
	if cnt > 0 {
		fini := (idx + 1) == cnt
		if fini {
			ret = "."
		} else {
			ret = ", and "
		}
	}
	return ret
}

// return values? its basically for build cmd i think
// does add process really need command info? why?

func NewStoryModel(db ScriptDB, types inspect.Types) *ModelMaker {
	m := NewModelMaker(db, types)

	m.Blocks.CreateOn("Declaration", "dl-group")
	m.Blocks.CreateOn("Execute", "dl-line")
	m.Params.AddFilter("InLocation.Location", QuoteFilter)
	m.Params.AddFilter("ScriptSubject.Subject", QuoteFilter)
	m.Params.AddFilter("Text.Value", QuoteFilter)
	m.Params.AddFilter("Chapter.Name", BlockEndFilter)

	m.Params.AddProcess("Execute?array=true", func(stack *Stack) error {
		return stack.NewBlock("dl-scope", m.BuildArray)
	})

	// m.Params.AddProcess("Declaration?array=true", func(stack *Stack) error {
	// 	return m.BuildElements(stack, func(data *ArrayData, i int) (err error) {
	// 		isLast := (i + 1) == len(data.Array)
	// 		if e := m.BuildCmd(stack); e != nil {
	// 			err = e
	// 		} else if isLast {
	// 			stack.LastChild.Sep = FullStopSep
	// 		} else {
	// 			stack.LastChild.Sep = StopGapSep
	// 		}
	// 		return
	// 	})
	// })

	m.Commands.AddProcess("NounPhrase", func(cmd *inspect.CommandInfo, stack *Stack) (err error) {
		if e := m.ProcessCmd(cmd, stack); e != nil {
			err = e
		} else {
			stack.LastChild.Sep = FullStopSep
		}
		return
	})

	m.Params.AddProcess("NounPhrase.Fragments", func(stack *Stack) error {
		return m.BuildElements(stack, func(array *ArrayData, i int) error {
			return stack.Command(func(cmd *inspect.CommandInfo) (err error) {
				called := cmd.Name == "ScriptSubject"
				if e := m.BuildCmd(stack); e != nil {
					err = e
				} else if called {
					stack.LastChild.Sep = SpaceSep
				} else {
					stack.LastChild.Sep = CommaSep
				}
				return
			})
		})
	})

	return m
}
