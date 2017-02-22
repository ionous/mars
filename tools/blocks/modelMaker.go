package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	"strings"
)

type ModelMaker struct {
	db    ScriptDB
	types inspect.Types

	Blocks   BlockTags
	Params   ParamProcs
	Commands CommandProcs
}

type BuildFn func(*Stack) error
type FilterFn func(string) string
type CommandFn func(*inspect.CommandInfo, *Stack) error
type ArrayFn func(src *ArrayData, i int) (err error)

func NewModelMaker(db ScriptDB, types inspect.Types) *ModelMaker {
	return &ModelMaker{db: db, types: types, Blocks: make(BlockTags), Params: make(ParamProcs), Commands: make(CommandProcs)}
}

func (m *ModelMaker) BuildPrimitive(stack *Stack) error {
	return stack.Data(func(data interface{}) (err error) {
		if src, ok := data.(*PrimData); !ok {
			err = errutil.New("no primitive data")
		} else if text, e := Format(src.Value); e != nil {
			err = e
		} else {
			stack.NewSpan("st-prim", func(span *Span) {
				span.Text = text
			})
		}
		return
	})
}

func (m *ModelMaker) BuildArray(stack *Stack) error {
	return m.BuildElements(stack, nil)
}

func (m *ModelMaker) BuildElements(stack *Stack, buildEl ArrayFn) (err error) {
	return stack.Data(func(data interface{}) (err error) {
		if src, ok := data.(*ArrayData); !ok {
			err = errutil.New("array mismatch", stack.Path(), sbuf.Type{data})
		} else {
			for i, kid := range src.Array {
				path := stack.ChildPath(kid)
				if e := stack.NewPath(path, func(data interface{}) (err error) {
					if cmd, e := m.commandFromData(data); e != nil {
						err = e
					} else {
						err = stack.NewCommand(cmd, func() (err error) {
							if buildEl != nil {
								err = buildEl(src, i)
							} else {
								err = m.BuildCmd(stack)
							}
							return
						})
					}
					return
				}); e != nil {
					err = e
					break
				}
			}
		}
		return
	})
}

// ProcessCmd, without looking up procs
func (m *ModelMaker) ProcessCmd(cmd *inspect.CommandInfo, stack *Stack) (err error) {
	tags := []string{}
	for _, name := range cmd.Types() {
		if tag, ok := m.Blocks[name]; ok {
			tags = append(tags, tag)
		}
	}
	if len(tags) == 0 {
		err = m.innerBuild(stack)
	} else {
		tag := strings.Join(tags, " ")
		err = stack.NewBlock(tag, func(stack *Stack) error {
			return m.innerBuild(stack)
		})
	}
	return
}

func (m *ModelMaker) BuildCmd(stack *Stack) error {
	return stack.Command(func(cmd *inspect.CommandInfo) (err error) {
		proc, ok := m.Commands[cmd.Name]
		if !ok {
			proc = m.ProcessCmd
		}
		if e := proc(cmd, stack); e != nil {
			err = errutil.New("couldnt build cmd", cmd.Name, e)
		}
		return
	})
}

//
func (m *ModelMaker) BuildRootCmd(path string) (block *Block, blocks *Blocks, err error) {
	blocks = NewBlocks(m.db)
	stack := NewStack(m.db, blocks)
	err = stack.NewPath(path, func(data interface{}) (err error) {
		if cmd, e := m.commandFromData(data); e != nil {
			err = errutil.New("root command error", e)
		} else {
			err = stack.NewCommand(cmd, func() (err error) {
				if b, e := stack.NewRoot("st-root", func(s *Stack) error {
					return m.BuildCmd(s)
				}); e != nil {
					err = e
				} else {
					block = b
				}
				return
			})
		}
		return
	})
	return
}
