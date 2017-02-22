package blocks

import (
	"github.com/ionous/mars/tools/inspect"
	"github.com/ionous/sashimi/util/errutil"
	"github.com/ionous/sashimi/util/sbuf"
	"log"
	"strings"
)

type ModelMaker struct {
	db    ScriptDB
	types inspect.Types

	Blocks   BlockTags
	Params   ParamProcs
	Commands CommandProcs
}

var _ = log.Println

type BuildFn func(*Stack) error
type FilterFn func(string) string
type ParamFn func(*inspect.ParamInfo, *Stack) error
type CommandFn func(*inspect.CommandInfo, *Stack) error
type ArrayFn func(_ *inspect.CommandInfo, _ *ArrayData, idx int) (err error)

func NewModelMaker(db ScriptDB, types inspect.Types) *ModelMaker {
	return &ModelMaker{db: db, types: types, Blocks: make(BlockTags), Params: make(ParamProcs), Commands: make(CommandProcs)}
}

func (m *ModelMaker) BuildPrimitive(stack *Stack) error {
	return stack.Data(func(data interface{}) (err error) {
		if src, ok := data.(*PrimData); !ok {
			err = errutil.New("no primitive data")
		} else if text, e := Format(src.Value); e != nil {
			err = errutil.New("error building primitive, because", e)
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
			err = errutil.New("array type mismatch at", sbuf.Q(stack.Path()), sbuf.Type{data})
		} else {
			for i, kid := range src.Array {
				if e := stack.NewPath(kid, func(data interface{}) (err error) {
					if cmd, e := m.commandFromData(data); e != nil {
						err = e
					} else if buildEl != nil {
						err = buildEl(cmd, src, i)
					} else {
						err = m.BuildCmd(cmd, stack)
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
		err = stack.NewBlock(tag, m.innerBuild)
	}
	return
}

// BuildCmd,
func (m *ModelMaker) BuildCmd(cmd *inspect.CommandInfo, stack *Stack) error {
	return stack.NewCommand(cmd, func() (err error) {
		proc := m.ProcessCmd
		for _, c := range cmd.Types() {
			if p, ok := m.Commands[c]; ok {
				proc = p
				break
			}
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
			if b, e := stack.NewRoot("st-root", func(s *Stack) error {
				return m.BuildCmd(cmd, s)
			}); e != nil {
				err = e
			} else {
				block = b
			}
		}
		return
	})
	return
}
